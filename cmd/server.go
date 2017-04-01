package cmd

import (
	"github.com/danlin/mutago/httpapi"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(serverCmd)
	serverCmd.PersistentFlags().StringVarP(&path, "data", "", "./data", "path where store data")
	serverCmd.PersistentFlags().StringVarP(&bind, "bind", "", "127.0.0.1", "interface to which the server will bind")
	serverCmd.PersistentFlags().IntVarP(&port, "port", "p", 9000, "port on which the server will listen")
}

var path string
var bind string
var port int

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the API Server",
	Long:  "Start the API Server",
	Run: func(cmd *cobra.Command, args []string) {
		httpapi.Start(path, bind, port)
	},
}
