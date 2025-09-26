package utils

import (
	"fmt"
	"github.com/fatih/color"
)

// 显示pfinalclub的彩色logo
func DisplayLogo() {
	// 创建不同颜色的打印器
	red := color.New(color.FgRed).SprintFunc()	
	green := color.New(color.FgGreen).SprintFunc()	
	cyan := color.New(color.FgCyan).SprintFunc()	
	magenta := color.New(color.FgMagenta).SprintFunc()	
	yellow := color.New(color.FgYellow).SprintFunc()

	// 打印彩色logo
	fmt.Println()
	fmt.Println(magenta("        ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄ "))
	fmt.Println(magenta("       ▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌"))
	fmt.Println(magenta("       ▐░█▀▀▀▀▀▀▀█░▌ ▀▀▀▀█░█▀▀▀▀ ▐░█▀▀▀▀▀▀▀▀▀ "))
	fmt.Println(magenta("       ▐░▌       ▐░▌      ▐░▌     ▐░▌          "))
	fmt.Println(red("       ▐░█▄▄▄▄▄▄▄█░▌      ▐░▌     ▐░█▄▄▄▄▄▄▄▄▄ "))
	fmt.Println(red("       ▐░░░░░░░░░░░▌      ▐░▌     ▐░░░░░░░░░░░▌"))
	fmt.Println(red("       ▐░█▀▀▀▀▀▀▀█░▌      ▐░▌     ▐░█▀▀▀▀▀▀▀▀▀ "))
	fmt.Println(green("       ▐░▌       ▐░▌      ▐░▌     ▐░▌          "))
	fmt.Println(green("       ▐░█▄▄▄▄▄▄▄█░▌      ▐░▌     ▐░█▄▄▄▄▄▄▄▄▄ "))
	fmt.Println(green("       ▐░░░░░░░░░░░▌      ▐░▌     ▐░░░░░░░░░░░▌"))
	fmt.Println(cyan("        ▀▀▀▀▀▀▀▀▀▀▀        ▀        ▀▀▀▀▀▀▀▀▀▀▀ "))
	fmt.Println()
	fmt.Println(yellow("               工具参数可视化配置工具              "))
	fmt.Println()
}