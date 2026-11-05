package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/agentgateway/agentgateway/controller/pkg/setup"
	"github.com/agentgateway/agentgateway/controller/pkg/version"
)

func main() {
	var agentgatewayVersion bool
	cmd := &cobra.Command{
		Use:          "agentgateway",
		Short:        "Runs the Moonwitness Gateway controller",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if agentgatewayVersion {
				fmt.Println(version.String())
				return nil
			}
			s, err := setup.New(setup.Options{})
			if err != nil {
				return fmt.Errorf("error setting up Moonwitness Gateway: %w", err)
			}
			if err := s.Start(cmd.Context()); err != nil {
				return fmt.Errorf("err in main: %w", err)
			}

			return nil
		},
	}
	cmd.Flags().BoolVarP(&agentgatewayVersion, "version", "v", false, "Print the Moonwitness Gateway controller version")

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
