package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newStreamCmd())
}

func newStreamCmd() *cobra.Command {
	var duration time.Duration

	cmd := &cobra.Command{
		Use:   "stream",
		Short: "Streams simulated events from the server.",
		Long:  `Connects to the server and listens for incoming events, printing them to the console.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			container, err := initContainer()
			if err != nil {
				return err
			}
			logger := container.Logger
			logger.Info("Executing 'stream' command...", "duration", duration)

			// Implicitly connect before streaming
			logger.Info("Implicitly connecting before streaming...")
			_, err = container.ConnectUseCase.Execute(context.Background())
			if err != nil {
				logger.Error("Implicit connection failed", err)
				return fmt.Errorf("could not connect before streaming: %w", err)
			}

			// Create a context that will be cancelled after the duration
			ctx, cancel := context.WithTimeout(context.Background(), duration)
			defer cancel()

			logger.Info("Starting event stream...")
			fmt.Printf("Streaming events for %s. Press Ctrl+C to stop early.\n", duration)

			err = container.ReceiveUseCase.Execute(ctx)
			if err != nil && err != context.DeadlineExceeded {
				logger.Error("Stream command failed", err)
				return err
			}

			logger.Info("Stream command finished.")
			fmt.Println("\nStream finished.")
			return nil
		},
	}

	cmd.Flags().DurationVar(&duration, "duration", 10*time.Second, "Duration to stream events for (e.g., 5s, 1m)")

	return cmd
}
