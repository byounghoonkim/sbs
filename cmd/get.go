package cmd

import (
	"io"
	"io/ioutil"
	"log"
	"os"

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

		tls, err := cmd.Flags().GetBool("tls")
		caFile, err := cmd.Flags().GetString("ca_file")
		serverHostOverride, err := cmd.Flags().GetString("server_host_override")

		output, err := cmd.Flags().GetString("output")

		b, err := client.NewBlobServiceClient(host, port, tls, caFile, serverHostOverride)
		if err != nil {
			log.Fatal(err)
		}

		c, err := client.NewClient(b)
		if err != nil {
			log.Fatal(err)
		}

		ID := args[0]

		var w io.Writer

		if output != "" {
			log.Printf("output : %s", output)
			f, err := os.Create(output)
			if err != nil {
				log.Fatal(err)
			}
			w = f
			defer f.Close()
		} else {
			w = ioutil.Discard
		}

		n, err := c.Get(ID, w)
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

	getCmd.Flags().BoolP("tls", "", false, "use tls connection")
	getCmd.Flags().StringP("ca_file", "", "", "path to ca file")
	getCmd.Flags().StringP("server_host_override", "", "", "host name for override")

	getCmd.Flags().StringP("output", "o", "", "Path to save blob to file")
}
