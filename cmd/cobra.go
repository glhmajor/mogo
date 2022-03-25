package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"github.com/glhmajor/mogo/cmd/api/duplicate"
	"fmt"
)


var rootCmd = &cobra.Command{
	Use: "mogo",
	Short: "mogo is a Tool aggregate。",
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `欢迎使用 mogo 可以使用 -h 查看命令`
	fmt.Printf("%s\n", usageStr)
}

func init() {
	rootCmd.AddCommand(duplicate.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
