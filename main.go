package main

import (
	"log"
	"os"

	"github.com/quangghaa/grpc-demo/cmd"
)

func main() {
	os.Setenv("TZ", "Asia/Ho_Chi_Minh")
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(-1)
	}
}
