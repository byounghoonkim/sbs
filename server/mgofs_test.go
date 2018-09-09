package server

import "testing"

func TestMgoCreate(t *testing.T) {

	m := NewMgoFS("mongodb://localhost:27017", "sbs", "fs")

	f, err := m.Create("test")

	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

}

func TestMgoOpen(t *testing.T) {

	m := NewMgoFS("mongodb://localhost:27017", "sbs", "fs")

	f, err := m.Open("test")

	if err != nil {
		t.Fatal(err)
	}

	defer f.Close()

}
