/*
Copyright © 2025 pfinalclub <contact@pfinalclub.com>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/pfinal/pfinal_cli_tools/internal/utils"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pfinal_cli_tools",
	Short: "一个工具参数可视化配置工具",
	Long: `pfinal_cli_tools 是一个用于简化命令行工具参数配置的辅助工具。
它可以帮助您可视化地配置 nmap、wrk、sqlmap 等工具的参数，
无需记忆复杂的参数选项。`,
	// 当直接运行命令时，显示logo
	Run: func(cmd *cobra.Command, args []string) {
		utils.DisplayLogo()
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pfinal_cli_tools.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("check-tools", "c", false, "检查支持的工具是否已安装")
	
	// 处理检查工具的flag
	oldPreRun := rootCmd.PersistentPreRun
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		// 先执行原来的preRun函数
		if oldPreRun != nil {
			oldPreRun(cmd, args)
		}
		
		// 检查是否设置了检查工具的flag
		checkTools, _ := cmd.Flags().GetBool("check-tools")
		if checkTools {
			utils.DisplayToolCheckResults()
			os.Exit(0)
		}
	}
}


