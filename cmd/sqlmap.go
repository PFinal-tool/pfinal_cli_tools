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

// sqlmapCmd represents the sqlmap command
var sqlmapCmd = &cobra.Command{
	Use:   "sqlmap",
	Short: "交互式配置sqlmap扫描参数",
	Long: `交互式配置sqlmap SQL注入检测工具参数，无需记忆复杂的选项。
例如：
  pfinal_cli_tools sqlmap  # 启动交互式配置界面
  pfinal_cli_tools sqlmap -u http://example.com/?id=1  # 快速扫描`,
	Run: func(cmd *cobra.Command, args []string) {
		// 创建颜色打印器
		red := color.New(color.FgRed).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		
		// 检查sqlmap工具是否已安装
		if !utils.IsToolInstalled("sqlmap") {
			fmt.Println(red("错误: sqlmap工具未安装，请先安装sqlmap。"))
			fmt.Println("您可以使用以下命令安装sqlmap：")
			fmt.Println("- git clone https://github.com/sqlmapproject/sqlmap.git")
			fmt.Println("- 或访问sqlmap.org下载最新版本")
			return
		}
		
		fmt.Println(green("✓ sqlmap工具已安装，正在启动交互式配置..."))

		// 获取sqlmap模板
		template := config.GetSqlmapTemplate()

		// 启动交互式配置
		result := ui.StartInteractiveConfig(template)

		// 生成sqlmap命令
		cmdStr := buildSqlmapCommand(result)

		// 显示生成的命令
		fmt.Println("生成的sqlmap命令：")
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

// 构建sqlmap命令
func buildSqlmapCommand(params map[string]string) string {
	var cmdBuilder strings.Builder
	cmdBuilder.WriteString("sqlmap")

	// 添加URL
	if url, ok := params["url"]; ok && url != "" {
		cmdBuilder.WriteString(" -u ")
		cmdBuilder.WriteString(url)
	}

	// 添加HTTP方法
	if method, ok := params["method"]; ok && method != "" && method != "GET" {
		cmdBuilder.WriteString(" --method=")
		cmdBuilder.WriteString(method)
	}

	// 添加POST数据
	if data, ok := params["data"]; ok && data != "" {
		cmdBuilder.WriteString(" --data=")
		cmdBuilder.WriteString(data)
	}

	// 添加测试级别
	if level, ok := params["level"]; ok && level != "" {
		cmdBuilder.WriteString(" --level=")
		cmdBuilder.WriteString(level)
	}

	// 添加风险级别
	if risk, ok := params["risk"]; ok && risk != "" {
		cmdBuilder.WriteString(" --risk=")
		cmdBuilder.WriteString(risk)
	}

	return cmdBuilder.String()
}

func init() {
	rootCmd.AddCommand(sqlmapCmd)

	// 这里可以添加sqlmap命令的标志
	sqlmapCmd.Flags().StringP("url", "u", "", "目标URL")
	sqlmapCmd.Flags().StringP("method", "X", "GET", "HTTP方法")
	sqlmapCmd.Flags().StringP("data", "d", "", "POST数据")
	sqlmapCmd.Flags().StringP("level", "l", "1", "测试级别(1-5)")
	sqlmapCmd.Flags().StringP("risk", "r", "1", "风险级别(1-3)")
}