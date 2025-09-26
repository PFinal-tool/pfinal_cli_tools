package utils

import (
	"fmt"
	"os/exec"
	"strings"
	"github.com/fatih/color"
)

// 工具检测结果结构体
type ToolCheckResult struct {
	Name     string
	Version  string
	Installed bool
}

// 检查工具是否已安装
func CheckToolInstalled(toolName string) ToolCheckResult {
	result := ToolCheckResult{
		Name:     toolName,
		Installed: false,
		Version:  "未安装",
	}

	// 尝试执行工具命令来检查是否安装
	cmd := exec.Command("which", toolName)
	_, err := cmd.CombinedOutput()

	if err == nil {
		result.Installed = true
		// 尝试获取版本信息
		var versionCmd *exec.Cmd

		// 根据不同工具使用不同的版本参数
		switch toolName {
		case "nmap":
			versionCmd = exec.Command(toolName, "--version")
		case "wrk":
			versionCmd = exec.Command(toolName, "-v")
		case "sqlmap":
			versionCmd = exec.Command(toolName, "--version")
		default:
			versionCmd = exec.Command(toolName, "--version")
		}

		// 执行版本命令
		versionOutput, err := versionCmd.CombinedOutput()
		if err == nil {
			versionStr := string(versionOutput)
			// 提取前两行作为版本信息摘要
			lines := strings.Split(versionStr, "\n")
			var versionSummary string
			for i, line := range lines {
				if i < 2 && line != "" {
					versionSummary += line + "\n"
				} else if i >= 2 {
					break
				}
			}
			result.Version = strings.TrimSpace(versionSummary)
		} else {
			result.Version = "已安装但无法获取版本信息"
		}
	}

	return result
}

// 显示所有支持的工具检测结果
func DisplayToolCheckResults() {
	// 创建颜色打印器
	green := color.New(color.FgGreen).SprintFunc()	
	red := color.New(color.FgRed).SprintFunc()	
	cyan := color.New(color.FgCyan).SprintFunc()

	tools := []string{"nmap", "wrk", "sqlmap"}

	fmt.Println(cyan("\n工具安装状态检测结果："))
	fmt.Println(cyan("----------------------"))

	for _, tool := range tools {
		result := CheckToolInstalled(tool)
		status := red("✗ 未安装")
		if result.Installed {
			status = green("✓ 已安装")
		}
		fmt.Printf("%s: %s\n", tool, status)
		if result.Installed && result.Version != "" {
			fmt.Printf("  版本信息: %s\n", result.Version)
		}
	}

	fmt.Println(cyan("----------------------\n"))
}

// 检查指定工具是否已安装，返回布尔值
func IsToolInstalled(toolName string) bool {
	result := CheckToolInstalled(toolName)
	return result.Installed
}