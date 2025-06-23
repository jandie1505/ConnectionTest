package main

import (
	"ConnectionTest/ping"
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	switch flag.Arg(0) {
	case "ping":
		ping.Application()
		break
	default:
		fmt.Println("Available modes: ping, analyze (planned)")
		return
	}
}
