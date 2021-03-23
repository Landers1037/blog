/*
landers Apps
Author: landers
Github: github.com/landers1037
*/

package logger

import (
	"fmt"
)

const logo = "\n ____  _     ____  _____\n/  __\\/ \\   /  _ \\/  __/\n| | //| |   | / \\|| |  _\n| |_\\\\| |_/\\| \\_/|| |_//\n\\____/\\____/\\____/\\____\\"

func PrintLogo() {
	fmt.Printf("++++++++++++++++++++++++%s\n\n++++++++++++++++++++++++\n\n", logo)
}
