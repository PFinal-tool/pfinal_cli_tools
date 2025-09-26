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

// wrkCmd represents the wrk command
var wrkCmd = &cobra.Command{
	Use:   "wrk",
	Short: "交互式配置wrk性能测试参数",
	Long: `交互式配置wrk性能测试工具参数，无需记忆复杂的选项。
例如：
  pfinal_cli_tools wrk  # 启动交互式配置界面
  pfinal_cli_tools wrk -u http://example.com -t 2 -c 100 -d 10s  # 快速测试`,
	Run: func(cmd *cobra.Command, args []string) {
		// 创建颜色打印器
		red := color.New(color.FgRed).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		
		// 检查wrk工具是否已安装
		if !utils.IsToolInstalled("wrk") {
			fmt.Println(red("错误: wrk工具未安装，请先安装wrk。"))
			fmt.Println("您可以使用以下命令安装wrk：")
			fmt.Println("- Ubuntu/Debian: sudo apt install wrk")
			fmt.Println("- CentOS/RHEL: 需要从源码编译安装")
			fmt.Println("- macOS: brew install wrk")
			return
		}
		
		fmt.Println(green("✓ wrk工具已安装，正在启动交互式配置..."))

		// 获取wrk模板
		template := config.GetWrkTemplate()

		// 启动交互式配置
		result := ui.StartInteractiveConfig(template)

		// 生成wrk命令
		cmdStr := buildWrkCommand(result)

		// 显示生成的命令
		fmt.Println("生成的wrk命令：")
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

// 构建wrk命令
func buildWrkCommand(params map[string]string) string {
	var cmdBuilder strings.Builder
	cmdBuilder.WriteString("wrk")

	// 添加线程数
	if threads, ok := params["threads"]; ok && threads != "" {
		cmdBuilder.WriteString(" -t")
		cmdBuilder.WriteString(threads)
	}

	// 添加连接数
	if connections, ok := params["connections"]; ok && connections != "" {
		cmdBuilder.WriteString(" -c")
		cmdBuilder.WriteString(connections)
	}

	// 添加测试持续时间
	if duration, ok := params["duration"]; ok && duration != "" {
		cmdBuilder.WriteString(" -d")
		cmdBuilder.WriteString(duration)
	}

	// 添加URL
	if url, ok := params["url"]; ok && url != "" {
		cmdBuilder.WriteString(" ")
		cmdBuilder.WriteString(url)
	}

	return cmdBuilder.String()
}

func init() {
	rootCmd.AddCommand(wrkCmd)

	// 这里可以添加wrk命令的标志
	wrkCmd.Flags().StringP("url", "u", "", "测试目标URL")
	wrkCmd.Flags().StringP("threads", "t", "2", "线程数量")
	wrkCmd.Flags().StringP("connections", "c", "100", "连接数量")
	wrkCmd.Flags().StringP("duration", "d", "10s", "测试持续时间")
}