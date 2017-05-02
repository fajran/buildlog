package main

import (
	"fmt"

	"github.com/fajran/buildlog/pkg/buildlog"
	"github.com/fajran/buildlog/pkg/server"
)

func main() {
	bl := buildlog.NewBuildLog()
	s := server.NewServer(":8080", bl)
	fmt.Printf("Address: %s\n", s.Addr)
	s.Start()
}
