package debug

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"network-cli-go/commands"
)

var (
	packageName string
	method      string
	params      []string
)

func init() {
	debugCmd := cobra.Command{
		Use: "debug",
		Run: func(cmd *cobra.Command, args []string) {
			logrus.Infof("debug package %s method %s params %s", packageName, method, params)
			if packageName == "net" {
				debugNet(method, params)
			}
		},
	}
	debugCmd.Flags().StringVar(&packageName, "packageName", "net", "debug package name")
	debugCmd.Flags().StringVar(&method, "method", "", "debug method name")
	debugCmd.Flags().StringSliceVar(&params, "params", nil, "params separated by comma")
	commands.RootCmd.AddCommand(&debugCmd)
}
