package cmd

import (
	"github.com/spf13/cobra"
	"goconsul/core"
)

var (
	consulAddr  string
	consulToken string
	prefix string
	output string


	rootCmd = &cobra.Command{
		Use:   "goconsul",
		Short: "cli to list all the keys",
		Long:  `goconsul  is a library for Go to list/export all the keys and their values in consul-kv.`,
		Run: func(cmd *cobra.Command, args []string) {
			core.App(consulAddr, consulToken, prefix, output)

		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&consulAddr, "consulAddr", "a", "127.0.0.1:8500", "Consul http address.")
	rootCmd.PersistentFlags().StringVarP(&consulToken, "consulToken", "t", "", "Consul http token.")
	rootCmd.PersistentFlags().StringVarP(&prefix, "prefix", "p", "all", "Consul prefix. ex:- config/")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Export all the content to `json|yaml|xml|csv` format.")
}
