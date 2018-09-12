package server

import (
	"net"
	"testing"
)

func availableLocalMongo() bool {
	conn, err := net.Dial("tcp", ":27017")
	if err != nil {
		return false
	}
	defer conn.Close()

	return true

}

func TestMgoCreate(t *testing.T) {
	if !availableLocalMongo() {
		t.Skip("local mongo not available")
	}

	m := NewMgoFS("mongodb://localhost:27017", "sbs", "fs")

	f, err := m.Create("test")

	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

}

func TestMgoOpen(t *testing.T) {
	if !availableLocalMongo() {
		t.Skip("local mongo not available")
	}

	m := NewMgoFS("mongodb://localhost:27017", "sbs", "fs")

	f, err := m.Open("test")

	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

}
