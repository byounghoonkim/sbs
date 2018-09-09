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

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start sbs server",
	Run: func(cmd *cobra.Command, args []string) {

		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			log.Fatalf("failed to get port number: %v", err)
		}

		host, err := cmd.Flags().GetString("host")
		if err != nil {
			log.Fatalf("failed to get host: %v", err)
		}

		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatalf("failed to get path: %v", err)
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

		var fb server.Provider
		if path != "" {
			fb = server.NewFileBase(path)
		} else {
			log.Print("serving with memory storage")
			fb = server.NewFileBase(".").WithFs(afero.NewMemMapFs())
		}

		s := server.NewServer(fb)

		blob.RegisterBlobServiceServer(grpcServer, s)
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatalf("failed to serve grpcServer : %v", err)
		}

	},
}

// Default Values.
const (
	DefaultPort = 2018
	DefaultHost = "localhost"
)

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntP("port", "p", DefaultPort, "Port number of server")
	serveCmd.Flags().StringP("host", "s", DefaultHost, "host address to listen")

	serveCmd.Flags().StringP("path", "", "", "path to save blobs")

}
