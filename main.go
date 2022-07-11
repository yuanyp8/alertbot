package main

import (
	"fmt"
	"github.com/yuanyp8/alertbot/cmd"
	"os"
)

func main() {
	_ = os.Setenv("CGO_ENABLED", "0")
	_ = os.Setenv("GOARCH", "amd64")
	_ = os.Setenv("GOOS", "linux")

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}

}
