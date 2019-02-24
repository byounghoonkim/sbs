package cmd

import (
	"io/ioutil"
	"log"
	"net"
	"os"
	"testing"
	"time"
)

func TestMakeProvider(t *testing.T) {
	_, err := makeProvider("")
	if err != nil {
		log.Fatal(err)
	}

	_, err = makeProvider("mongodb://localhost:20217")
	if err != nil {
		log.Fatal(err)
	}

	_, err = makeProvider("xxxxxxx://localhost:20217")
	if err == nil {
		log.Fatalf("makeProvider return unsupport db error")
	}
}

func TestServeArg(t *testing.T) {
	err := serveCmd.ParseFlags([]string{"--path", ""})
	if err != nil {
		log.Fatalf("failed to parse commandline arg : %v", err)
	}

	log.Print(serveCmd)
	//serveCmd.Run(serveCmd, nil)
}

func TestServe(t *testing.T) {
	content := []byte("temporary file's content")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	go func() {
		serveCmd.ParseFlags([]string{"--tls"})
		serveCmd.Run(serveCmd, nil)
	}()

	for {
		conn, _ := net.DialTimeout("tcp", net.JoinHostPort("localhost", "2018"), time.Second*5)
		if conn != nil {
			conn.Close()
			break
		}
	}

	putCmd.ParseFlags([]string{
		"--tls",
		"--server_host_override",
		"sbs.test.youtube.com",
	})
	putCmd.Run(putCmd, []string{tmpfile.Name()})

	getCmd.ParseFlags([]string{
		"--tls",
		"--server_host_override",
		"sbs.test.youtube.com",
	})
	getCmd.Run(getCmd, []string{tmpfile.Name()})

}
