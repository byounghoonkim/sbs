package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/shoebillk/sbs/blob"
	"github.com/shoebillk/sbs/server"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var port = 2018

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start sbs server",
	Run: func(cmd *cobra.Command, args []string) {

		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			log.Fatalf("failed to get port number", err)
		}

		host, err := cmd.Flags().GetString("host")
		if err != nil {
			log.Fatalf("failed to get host", err)
		}

		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		log.Printf("Listening - %s:%d", host, port)

		var opts []grpc.ServerOption
		/*
			if *tls {
				if *certFile == "" {
					*certFile = testdata.Path("server1.pem")
				}
				if *keyFile == "" {
					*keyFile = testdata.Path("server1.key")
				}
				creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
				if err != nil {
					log.Fatalf("Failed to generate credentials %v", err)
				}
				opts = []grpc.ServerOption{grpc.Creds(creds)}
			}
		*/

		grpcServer := grpc.NewServer(opts...)

		fb := server.NewFileBase(".").WithFs(afero.NewMemMapFs())
		s := server.NewServer(fb)

		blob.RegisterBlobServiceServer(grpcServer, s)
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatalf("failed to serve grpcServer : %v", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	serveCmd.Flags().IntP("port", "p", 2018, "Port number of server")
	serveCmd.Flags().StringP("host", "s", "localhost", "host address to listen")

}
