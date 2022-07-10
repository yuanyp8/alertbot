package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/yuanyp8/alertbot/apps"
	"github.com/yuanyp8/alertbot/conf"
)

var configFile string

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "start",
	Long:  "start",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := conf.C().LoadConf(configFile)
		if err != nil {
			return err
		}

		g := gin.Default()
		apps.H().Registry(g)

		return g.Run(conf.C().App.HttpApi)
	},
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&configFile, "config", "f", "etc/config.yaml", "alertbot 配置文件路径")
}
