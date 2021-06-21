package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	 "github.com/wukongcloud/xxdns/utils"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of workLog",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version:   %s\n", utils.Version)
		fmt.Printf("CommitID:  %s\n", utils.CommitID)
		fmt.Printf("BuildTime: %s\n", utils.BuildTime)
		fmt.Printf("GoVersion: %s\n", utils.GoVersion)
		fmt.Printf("BuildUser: %s\n", utils.BuildUser)
	},
}
