package cmd

import (

	"fmt"
	"github.com/spf13/cobra"
	"log"

	//"github.com/spf13/viper"
	"os"
)

//var cfgFile string
var (
	Name ="xxdns"

	Port string


)

const Logo = `                                                                      
                   .___             
___  ______  ___ __| _/____   ______
\  \/  /\  \/  // __ |/    \ /  ___/
 >    <  >    </ /_/ |   |  \\___ \ 
/__/\_ \/__/\_ \____ |___|  /____  >
      \/      \/    \/    \/     \/                    
`
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   Name,
	Short: fmt.Sprintf("%s 是基于bind9的DNS管理服务",Name),
	Long:  fmt.Sprintf("%s 是基于bind9的DNS管理服务",Name),

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	//cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(RunServer)
	rootCmd.AddCommand(RunAgent)

}



