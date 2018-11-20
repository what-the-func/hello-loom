package main

import (
	"log"

	"github.com/loomnetwork/go-loom/builtin/commands"
	"github.com/loomnetwork/go-loom/cli"
	"github.com/spf13/cobra"
	"github.com/what-the-func/hello-loom/types"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "cli",
		Short: "Hello World",
	}
	callCmd := cli.ContractCallCommand()
	rootCmd.AddCommand(callCmd)
	commands.Add(callCmd)

	defaultContract := "helloworld"
	var name string

	helloCmd := &cobra.Command{
		Use:   "hello",
		Short: "Calls the Hello method of the helloworld contract",
		RunE: func(cmd *cobra.Command, args []string) error {
			params := &types.HelloRequest{
				In: name,
			}
			var resp types.HelloResponse
			err := cli.StaticCallContract(defaultContract, "Hello", params, &resp)
			if err != nil {
				return err
			}
			log.Printf("%s", resp.Out)
			return nil
		},
	}
	helloCmd.Flags().StringVarP(&name, "name", "n", "Bob", "Your name")
	callCmd.AddCommand(helloCmd)

	rootCmd.Execute()
}
