package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newConnectCmd())
}

func newConnectCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "connect",
		Short: "Simulates connecting to the WhatsApp server.",
		Long: `Establishes a fake WebSocket connection, performs a handshake, and stores a session.
The outcome (success, timeout, failure) is determined by the configured probabilities and seed.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			container, err := initContainer()
			if err != nil {
				return err
			}

			logger := container.Logger
			logger.Info("Executing 'connect' command...")

			_, err = container.ConnectUseCase.Execute(context.Background())
			if err != nil {
				logger.Error("Connection command failed", err)
				// The error from the use case is already descriptive.
				return err
			}

			logger.Info("Connect command successful. Session is ready.")
			fmt.Println("Connection successful. Session ready.")
			return nil
		},
	}
}
