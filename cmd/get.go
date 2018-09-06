package cmd

import (
	"io/ioutil"
	"log"

	"github.com/shoebillk/sbs/client"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get blob data from server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("get called")
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			log.Fatal(err)
		}
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			log.Fatal(err)
		}

		c := client.NewClient(host, port)

		ID := args[0]
		n, err := c.Get(ID, ioutil.Discard)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Read : %d", n)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringP("host", "s", DefaultHost, "Host string of server")
	getCmd.Flags().IntP("port", "p", DefaultPort, "Port number of server")
}
