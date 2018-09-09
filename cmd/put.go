package cmd

import (
	"bufio"
	"log"
	"os"

	"github.com/shoebillk/sbs/client"
	"github.com/spf13/cobra"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "put a file to sbs server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("put called")

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

		b, err := client.NewBlobServiceClient(host, port, tls, caFile, serverHostOverride)
		if err != nil {
			log.Fatal(err)
		}

		c, err := client.NewClient(b)
		if err != nil {
			log.Fatal(err)
		}

		file := args[0]
		log.Printf("file : %s\n", file)
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		reader := bufio.NewReader(f)

		ID := file
		pushStatus, err := c.Push(ID, reader)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%#v", pushStatus)
	},
}

func init() {
	rootCmd.AddCommand(putCmd)

	putCmd.Flags().StringP("host", "s", DefaultHost, "Host string of server")
	putCmd.Flags().IntP("port", "p", DefaultPort, "Port number of server")

	putCmd.Flags().BoolP("tls", "", false, "use tls connection")
	putCmd.Flags().StringP("ca_file", "", "", "path to ca file")
	putCmd.Flags().StringP("server_host_override", "", "", "host name for override")
}
