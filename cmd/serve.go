package cmd

import (
	"fmt"
	"log"
	"net"
	"net/url"

	"github.com/shoebillk/sbs/blob"
	"github.com/shoebillk/sbs/server"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
)

func makeProvider(uri string) (*server.Provider, error) {
	var sp server.Provider
	if uri != "" {
		u, err := url.Parse(uri)
		if err != nil {
			return nil, err
		}

		switch u.Scheme {
		case "":
			log.Print("serving with filesystem base storage")
			sp = server.NewFileBase(uri)
		case "mongodb":
			log.Print("serving with mongodb base storage")
			sp = server.NewMgoFS(uri, "db", "sbs")
		default:
			return nil, fmt.Errorf("not support db - %s", u.Scheme)
		}

	} else {
		log.Print("serving with memory base storage")
		sp = server.NewFileBase(".").WithFs(afero.NewMemMapFs())
	}

	return &sp, nil
}

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

		tls, err := cmd.Flags().GetBool("tls")
		certFile, err := cmd.Flags().GetString("cert_file")
		keyFile, err := cmd.Flags().GetString("key_file")

		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		log.Printf("Listening - %s:%d", host, port)

		var opts []grpc.ServerOption
		if tls {
			if certFile == "" {
				log.Print("tls with testdata server1.pem")
				certFile = testdata.Path("server1.pem")
			}
			if keyFile == "" {
				log.Print("tls with testdata server1.key")
				keyFile = testdata.Path("server1.key")
			}
			creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
			if err != nil {
				log.Fatalf("Failed to generate credentials %v", err)
			}
			opts = []grpc.ServerOption{grpc.Creds(creds)}
		}

		grpcServer := grpc.NewServer(opts...)

		provider, err := makeProvider(path)

		if err != nil {
			log.Fatalf("failed  to make provider : %v", err)
		}

		s := server.NewServer(*provider)

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

	serveCmd.Flags().BoolP("tls", "", false, "use tls connection")
	serveCmd.Flags().StringP("ca_file", "", "", "path to ca file")
	serveCmd.Flags().StringP("key_file", "", "", "path to key file")

}
