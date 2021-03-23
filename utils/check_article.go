/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package utils

import (
	"bytes"
	"github.com/adrg/frontmatter"
	"regexp"
	"strconv"
	"strings"
)

// 校验文章的合法性

type YamlData struct {
	Name       string `yaml:"name"`
	ID         int	`yaml:"id"`
	Title      string `yaml:"title"`
	DatePlus   string `yaml:"date"`
	Tags       []string `yaml:"tags"`
	Categories []string `yaml:"categories"`
}

type Meta struct {
	Name       string
	ID         int
	Title      string
	Date       string
	DatePlus   string
	Tags       []string
	Categories []string
}

type MdData struct {
	Meta Meta
	Abstract string
	Body string
}

// 解析md源文件
// 一级解析器 生成不带meta的md字节
func ParseMd(fileRaw []byte) []byte {
	reg := regexp.MustCompile(`(?s)(^---[\s\S]*?---)(.*)`)
	res := reg.Find(fileRaw)
	res = reg.ReplaceAll(fileRaw, []byte("$2"))

	return res
}

// 解析出yaml front
// 使用yaml front解析yaml头部信息
func ParseYamlFront(fileRaw []byte) (Meta, error) {
	var ym YamlData
	_, parseErr := frontmatter.Parse(bytes.NewReader(fileRaw), &ym)
	if parseErr != nil {
		return Meta{}, parseErr
	}
	name := strings.TrimSpace(ym.Name)
	id, e := strconv.Atoi(string(ym.ID))
	if e != nil {
		id = 0
	}
	title := strings.TrimSpace(ym.Title)
	date := strings.Split(ym.DatePlus, " ")[0]
	dateplus := ym.DatePlus
	tags := ym.Tags
	cates := ym.Categories

	return Meta{
		Name:       name,
		ID:         id,
		Title:      title,
		Date:       date,
		DatePlus:   dateplus,
		Tags:       tags,
		Categories: cates,
	}, nil
}

// 解析摘要 传入为去除meta的文本
// 三级解析器 生成摘要信息
func ParseAbs(fileRaw []byte) []byte {
	// 摘要不存在时 使用默认摘要
	reg := regexp.MustCompile(`(?s)(.*)(<!--more-->)`)
	res := reg.Find(fileRaw)
	if len(res) <= 0 {
		return []byte("<code>Sorry</code>该文章暂无概述")
	}
	res = reg.ReplaceAll(res, []byte("$1"))

	return res
}

// 生成md文件的格式化结构体
func GenMdData(fileData []byte) (MdData, error) {
	// 得到meta
	meta, e := ParseYamlFront(fileData)
	if e != nil {
		return MdData{}, e
	}
	// 得到摘要
	abs := ParseAbs(ParseMd(fileData))
	// 得到正文
	body := ParseMd(fileData)

	return MdData{
		Meta:     meta,
		Abstract: string(abs),
		Body:     string(body),
	}, nil
}

// 校验上传的文章字节流是否合法
// 合法即包含需要的yaml front信息 且信息内的title和name不为空
func CheckArticleOK(fileData []byte) bool {
	meta, e := ParseYamlFront(fileData)
	if e != nil {
		return false
	}
	if meta.Name == "" || meta.Title == "" {
		return false
	}
	return true
}

// 重新格式化meta信息
// 可能出现时间不合法等情况 id不存在等情况
func ReformatMeta(meta Meta) Meta {
	var m Meta
	m.Title = meta.Title
	m.Name = meta.Name
	m.ID = meta.ID
	m.DatePlus = meta.DatePlus
	m.Date = meta.Date
	m.Tags = meta.Tags
	m.Categories = meta.Categories

	if !CheckDate(meta.Date) {
		m.Date = GetDate()
	}
	if !CheckDatePlus(meta.DatePlus) {
		m.DatePlus = GetDatePlus()
	}

	return m
}