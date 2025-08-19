package cli

import (
	"fmt"
	"os"

	"github.com/medinatello/wapp-socket/internal/app"
	"github.com/spf13/cobra"
)

var (
	// Global flags that can be used by any command.
	seed     int64
	logLevel string
	eventing string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "whats-cli",
	Short: "A CLI to interact with the wapp-socket application.",
	Long: `whats-cli is a tool to test and simulate WhatsApp flows using the wapp-socket backend.
It operates in a fake mode for Sprint 1, allowing for reproducible, simulated interactions.`,
	// PersistentPreRunE can be used to initialize the container once for all commands.
	// This is a good place to wire things up.
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// This function runs before any subcommand.
		// We can use it to override config from flags.
		// The actual container is initialized within each command's RunE function
		// to ensure dependencies are only created when needed.
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is the main entrypoint for the CLI application called by main.go.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your command: '%s'", err)
		os.Exit(1)
	}
}

func init() {
	// Define global flags for the root command.
	rootCmd.PersistentFlags().Int64Var(&seed, "seed", 0, "Random seed for reproducibility (0 for time-based)")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "Log level (debug, info, warn, error)")
	rootCmd.PersistentFlags().StringVar(&eventing, "eventing", "bus", "Eventing model to use (bus or chan)")

	// Here is where we would add child commands.
	// e.g., rootCmd.AddCommand(newConnectCmd())
}

// initContainer is a helper function to initialize the DI container.
// It will be called by the subcommands that need access to application services.
func initContainer() (*app.Container, error) {
	container, err := app.NewContainer()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize container: %w", err)
	}

	// Override config values from global flags.
	// This is a simple way to let flags take precedence over config files.
	if logLevel != "" {
		container.Config.App.LogLevel = logLevel
	}
	if seed != 0 {
		container.Config.Fakes.Seed = seed
	}
	if eventing != "" {
		container.Config.Features.EventingModel = eventing
	}

	// Note: The logger inside the container is already initialized.
	// For Sprint 1, we won't re-initialize it with the new log level from the flag.
	// A more advanced setup would handle this more gracefully.

	return container, nil
}
