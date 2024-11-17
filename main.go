package main

import (
	"fmt"
	"os"

	"github.com/xsamael/github-user-activity/model"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Use")
	}

	name := os.Args[1]
	model.ReqGithub(name)
}
