// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/shoebillk/sbs/blob"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "put a file to sbs server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("put called")

		host, err := cmd.Flags().GetString("host")
		if err != nil {
			log.Fatal(err)
		}
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			log.Fatal(err)
		}

		file := args[0]
		log.Printf("file : %s\n", file)

		target := fmt.Sprintf("%s:%d", host, port)
		log.Printf("Server : %s", target)

		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		reader := bufio.NewReader(f)

		chunkContentSize := 0x1000
		chunk := blob.Chunk{
			Id:      args[0],
			Content: make([]byte, chunkContentSize),
		}

		conn, err := grpc.Dial(target, grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		client := blob.NewBlobServiceClient(conn)

		var callopts []grpc.CallOption
		pushClient, err := client.Push(context.Background(), callopts...)

		if err != nil {
			log.Fatal(err)
		}

		for {
			n, err := reader.Read(chunk.Content)

			if err == io.EOF {
				log.Printf("Done to read")
				break
			} else if err != nil {
				log.Fatal(err)
				break
			}

			chunk.Content = chunk.Content[:n]
			log.Printf("Read %d", n)

			err = pushClient.Send(&chunk)
			if err != nil {
				log.Fatal(err)
				break
			}
		}

		pushStatus, err := pushClient.CloseAndRecv()
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
}
