package cli

import (
	"context"
	"fmt"

	"github.com/medinatello/wapp-socket/internal/domain"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newSendCmd())
}

func newSendCmd() *cobra.Command {
	var to, text string

	cmd := &cobra.Command{
		Use:   "send",
		Short: "Sends a fake text message.",
		Long:  `Simulates sending a text message to a specified JID. Requires a connection to be active.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if to == "" || text == "" {
				return fmt.Errorf("--to and --text flags are required")
			}

			container, err := initContainer()
			if err != nil {
				return err
			}

			logger := container.Logger
			logger.Info("Executing 'send' command...", "to", to)

			// This is a problem. The SendMessageUseCase needs the active connection,
			// but a new container doesn't have it. This is a limitation of the current
			// simple design. For Sprint 1, we can assume the CLI runs as a single process
			// and we need to connect first, but the state is not shared between commands.
			//
			// To solve this for the CLI, we need a persistent container or a way to
			// serialize the session state.
			//
			// A HACK for sprint 1: we will call ConnectUseCase implicitly. This is not ideal.
			// Let's document this in 00-bloqueos.md

			logger.Info("Implicitly connecting before sending...")
			_, err = container.ConnectUseCase.Execute(context.Background())
			if err != nil {
				logger.Error("Implicit connection failed", err)
				return fmt.Errorf("could not connect before sending: %w", err)
			}

			err = container.SendMessageUseCase.Execute(context.Background(), domain.JID(to), text)
			if err != nil {
				logger.Error("Send command failed", err)
				return err
			}

			logger.Info("Send command successful.")
			fmt.Println("Message sent successfully.")
			return nil
		},
	}

	cmd.Flags().StringVar(&to, "to", "", "JID of the recipient (e.g., 1234567890@s.whatsapp.net)")
	cmd.Flags().StringVar(&text, "text", "", "Text content of the message")
	_ = cmd.MarkFlagRequired("to")
	_ = cmd.MarkFlagRequired("text")

	return cmd
}
