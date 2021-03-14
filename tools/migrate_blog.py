import os
import argparse
import yaml
import re
from multiprocessing import Pool

'''
A small python script to migrate the yilia theme hexo markdown to 
the Renj.io blog theme.
The code is so simple that I don't want to give any annotation:)

created by xk-wang@2021/3/14 afternoon.
'''

parser = argparse.ArgumentParser()
parser.add_argument('-s', '--srcpath', type=str, help='The path to the '
	'markdown files that needs to be migrated')
parser.add_argument('-d', '--destpath', type=str, help='The path to the '
	'markdown files that have been migrated')

args = parser.parse_args()


def process_markdown(paths):

	srcpath, destpath = paths

	with open(srcpath) as file:
		filecontent = file.read()
	rangex = re.match('---[\s\S]*?---', filecontent).span()
	meta = yaml.load(filecontent[rangex[0]+4: rangex[1]-4], Loader=yaml.SafeLoader)
	migrated_meta = {'title': meta['title'], 'name': meta['abbrlink'], 'date': meta['date'], 
		'id': meta['abbrlink'], 'tags': meta['tags'], 'categories': '', 'abstract': ''}
	body = filecontent[rangex[1]:]
	meta_str = '---\n'
	for key, value in migrated_meta.items():
		meta_str += str(key) + ': ' + str(value) + '\n'
	meta_str += '---\n'
	migrated_content = meta_str + body

	with open(destpath, 'w') as f:
		f.write(migrated_content)


def migrate_markdown(args):
	pool = Pool(8)
	srcpath = args.srcpath
	destpath = args.destpath
	srcpaths = [os.path.join(srcpath, filename) for filename in os.listdir(srcpath) if filename.endswith('.md')]
	destpaths = [os.path.join(destpath, filename) for filename in os.listdir(srcpath)]

	pool.map_async(process_markdown, list(zip(srcpaths, destpaths)))


if __name__ == '__main__':
	migrate_markdown(args)