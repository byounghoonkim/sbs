package cmd

import (
	"io/ioutil"
	"log"
	"net"
	"os"
	"testing"
	"time"
)

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
		serveCmd.Run(serveCmd, nil)
	}()

	for {
		conn, _ := net.DialTimeout("tcp", net.JoinHostPort("localhost", "2018"), time.Second*5)
		if conn != nil {
			conn.Close()
			break
		}
	}

	putCmd.Run(putCmd, []string{tmpfile.Name()})
	getCmd.Run(getCmd, []string{tmpfile.Name()})

}
