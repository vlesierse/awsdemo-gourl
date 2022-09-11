package cmd

import (
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/vlesierse/awsdemo-gourl/api"
)

var port string
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "This starts the API server",
	Long:  `This command starts the gourl api server`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting server on " + port)
		router := api.NewRouter()
		log.Fatal(http.ListenAndServe(":"+port, router))
	},
}

func init() {
	serverCmd.PersistentFlags().StringVarP(&port, "port", "p", "8080", "This flag sets the port of our API server")
	rootCmd.AddCommand(serverCmd)
}
