package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wukongcloud/xxdns/agent"
	"os"
)

var (
	ServerIP string
)

func init() {
	RunAgent.Flags().StringVarP(&ServerIP, "server.ip", "S", "127.0.0.1:8080", "xxdns server ip and port.")
}

var RunAgent = &cobra.Command{
	Use:     "run.agent",
	Short:   "run xxdns agent.",
	Example: `	xxdns run.agent -S 127.0.0.1:8080`,
	Run: func(cmd *cobra.Command, args []string) {
		err := agent.RunAgent(ServerIP)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}
