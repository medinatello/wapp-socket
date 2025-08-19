package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newGroupsCmd())
}

// newGroupsCmd sets up the parent 'groups' command
func newGroupsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "groups",
		Short: "Manage groups (using dummy data for Sprint 1).",
		Long:  `The groups command provides subcommands to list and create fake groups.`,
	}

	cmd.AddCommand(newGroupsListCmd())
	cmd.AddCommand(newGroupsCreateCmd())

	return cmd
}

// newGroupsListCmd sets up the 'groups list' subcommand
func newGroupsListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists dummy groups.",
		Long:  `Retrieves and displays a hardcoded list of fake groups.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			container, err := initContainer()
			if err != nil {
				return err
			}
			logger := container.Logger
			logger.Info("Executing 'groups list' command...")

			groups, err := container.GroupsUseCase.List(context.Background())
			if err != nil {
				logger.Error("Groups list command failed", err)
				return err
			}

			logger.Info("Groups list command successful.", "count", len(groups))
			fmt.Println("Available Groups:")
			for _, group := range groups {
				fmt.Printf("- ID: %s, Name: %s\n", group.ID, group.Name)
			}
			return nil
		},
	}
	return cmd
}

// newGroupsCreateCmd sets up the 'groups create' subcommand
func newGroupsCreateCmd() *cobra.Command {
	var name string
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a dummy group.",
		Long:  `Simulates the creation of a group with the given name.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			container, err := initContainer()
			if err != nil {
				return err
			}
			logger := container.Logger
			logger.Info("Executing 'groups create' command...", "name", name)

			group, err := container.GroupsUseCase.Create(context.Background(), name)
			if err != nil {
				logger.Error("Groups create command failed", err)
				return err
			}

			logger.Info("Groups create command successful.", "id", group.ID, "name", group.Name)
			fmt.Printf("Successfully created group '%s' with fake ID '%s'.\n", group.Name, group.ID)
			return nil
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "Name of the group to create")
	_ = cmd.MarkFlagRequired("name")
	return cmd
}
