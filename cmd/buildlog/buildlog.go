package main

import (
	"fmt"

	"github.com/fajran/buildlog/server"
)

func main() {
	s := server.NewServer(":8080")
	fmt.Printf("Address: %s\n", s.Addr)
	s.Start()
}
