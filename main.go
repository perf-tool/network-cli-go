package main

import (
	"github.com/sirupsen/logrus"
	"network-cli-go/commands"
	_ "network-cli-go/commands/debug"
)

func main() {
	err := commands.RootCmd.Execute()
	if err != nil {
		logrus.Error("error is ", err)
		return
	}
}
