package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// RootCmd is root command
var rootCmd = &cobra.Command{
	Use:   "cal",
	Short: "レーザー強度を与えると, 規格化強度a0を計算します",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
var lambda int

func init() {
	rootCmd.AddCommand(a2iCmd)
	rootCmd.AddCommand(i2aCmd)

	pflag.IntVar(&lambda, "lambda", 810, "波長を設定します(単位 nm)")
}
func Execute() error {
	return rootCmd.Execute()
}
