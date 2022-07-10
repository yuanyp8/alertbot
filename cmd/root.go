package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd *cobra.Command = &cobra.Command{
	Use:   "alertbot",
	Short: "alertbot",
	Long:  "alertbot",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	RootCmd.AddCommand(StartCmd)
}
