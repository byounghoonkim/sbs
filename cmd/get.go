package cmd

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/shoebillk/sbs/blob"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
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

		Id := args[0]
		log.Printf("ID : %s\n", Id)

		target := fmt.Sprintf("%s:%d", host, port)
		log.Printf("Server : %s", target)
		conn, err := grpc.Dial(target, grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		client := blob.NewBlobServiceClient(conn)

		req := blob.GetRequest{Id: Id}

		var callopts []grpc.CallOption
		getClient, err := client.Get(context.Background(), &req, callopts...)
		if err != nil {
			log.Fatal(err)
		}

		for {
			chunk, err := getClient.Recv()

			if err == io.EOF {
				log.Printf("Done to read")
				break
			} else if err != nil {
				log.Fatal(err)
				break
			}

			log.Printf("Read %d", len(chunk.Content))
			log.Printf("Read %s", chunk.Content)

		}

	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringP("host", "s", DefaultHost, "Host string of server")
	getCmd.Flags().IntP("port", "p", DefaultPort, "Port number of server")
}
