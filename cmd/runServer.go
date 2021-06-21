package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wukongcloud/xxdns/server"
	common "github.com/wukongcloud/xxdns/utils"
)

func init() {
	RunServer.Flags().StringVarP(&Port, "port", "P", "8080", "web server listen port.")
}


var RunServer = &cobra.Command{
	Use:   "run.server",
	Short: "run xxdns server.",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("\r  \033[36%s\033[m  ", Logo)
		common.OpenBrowser(fmt.Sprintf("http://%s:%s/#/", "127.0.0.1", Port))
		server.RunServer(Port)


	},
}
