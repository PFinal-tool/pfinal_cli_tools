package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/pfinal/pfinal_cli_tools/internal/config"
	"github.com/pfinal/pfinal_cli_tools/internal/ui"
	"github.com/pfinal/pfinal_cli_tools/internal/utils"
	"github.com/spf13/cobra"
	"github.com/fatih/color"
)

// nmapCmd represents the nmap command
var nmapCmd = &cobra.Command{
	Use:   "nmap",
	Short: "交互式配置nmap扫描参数",
	Long: `交互式配置nmap扫描参数，无需记忆复杂的选项。
例如：
  pfinal_cli_tools nmap  # 启动交互式配置界面
  pfinal_cli_tools nmap -t simple -targs target_ip  # 使用简单模板快速扫描`,
	Run: func(cmd *cobra.Command, args []string) {
		// 创建颜色打印器
		red := color.New(color.FgRed).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		
		// 检查nmap工具是否已安装
		if !utils.IsToolInstalled("nmap") {
			fmt.Println(red("错误: nmap工具未安装，请先安装nmap。"))
			fmt.Println("您可以使用以下命令安装nmap：")
			fmt.Println("- Ubuntu/Debian: sudo apt install nmap")
			fmt.Println("- CentOS/RHEL: sudo yum install nmap")
			fmt.Println("- macOS: brew install nmap")
			return
		}
		
		fmt.Println(green("✓ nmap工具已安装，正在启动交互式配置..."))

		// 获取nmap模板
		template := config.GetNmapTemplate()

		// 启动交互式配置
		result := ui.StartInteractiveConfig(template)

		// 生成nmap命令
		cmdStr := buildNmapCommand(result)

		// 显示生成的命令
		fmt.Println("生成的nmap命令：")
		fmt.Println(cmdStr)

		// 询问是否执行命令
		fmt.Print("是否执行该命令？(y/n): ")
		var confirm string
		fmt.Scanln(&confirm)

		if strings.ToLower(confirm) == "y" {
			// 执行命令
			execCmd := exec.Command("/bin/bash", "-c", cmdStr)
			output, err := execCmd.CombinedOutput()
			if err != nil {
				fmt.Printf("命令执行失败：%s\n", err)
			}
			fmt.Printf("命令执行结果：\n%s\n", string(output))
		}
	},
}

// 构建nmap命令
func buildNmapCommand(params map[string]string) string {
	var cmdBuilder strings.Builder
	cmdBuilder.WriteString("nmap")

	// 添加参数
	if target, ok := params["target"]; ok && target != "" {
		cmdBuilder.WriteString(" ")
		cmdBuilder.WriteString(target)
	}

	if scanType, ok := params["scanType"]; ok && scanType != "" {
		cmdBuilder.WriteString(" ")
		cmdBuilder.WriteString(scanType)
	}

	if ports, ok := params["ports"]; ok && ports != "" {
		cmdBuilder.WriteString(" -p ")
		cmdBuilder.WriteString(ports)
	}

	if osDetect, ok := params["osDetect"]; ok && osDetect == "true" {
		cmdBuilder.WriteString(" -O")
	}

	if serviceDetect, ok := params["serviceDetect"]; ok && serviceDetect == "true" {
		cmdBuilder.WriteString(" -sV")
	}

	if verbose, ok := params["verbose"]; ok && verbose == "true" {
		cmdBuilder.WriteString(" -v")
	}

	return cmdBuilder.String()
}

func init() {
	rootCmd.AddCommand(nmapCmd)

	// 这里可以添加nmap命令的标志
	nmapCmd.Flags().StringP("template", "t", "", "使用预设模板")
	nmapCmd.Flags().StringP("template-args", "a", "", "模板参数")
}