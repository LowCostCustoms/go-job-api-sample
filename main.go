package main

import (
	"go-scheduler-api/cmd"
)

func main() {
	_ = cmd.NewCli().Execute()
}
