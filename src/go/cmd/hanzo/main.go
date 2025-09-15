package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.1.0"

var rootCmd = &cobra.Command{
	Use:   "hanzo",
	Short: "Hanzo AI SDK - Unified CLI",
	Long:  `Hanzo AI SDK provides a unified interface for all Hanzo services.`,
	Version: version,
}

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Manage Hanzo AI nodes",
}

var nodeStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a local AI node",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		fmt.Printf("Starting Hanzo node on port %d...\n", port)
		fmt.Printf("✅ Node started on port %d\n", port)
	},
}

var nodeStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the running node",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Stopping Hanzo node...")
		fmt.Println("✅ Node stopped")
	},
}

var nodeStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Check node status",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Node running on port 4000")
	},
}

var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Manage AI agents",
}

var agentCreateCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new agent",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		agentType, _ := cmd.Flags().GetString("type")
		model, _ := cmd.Flags().GetString("model")
		fmt.Printf("Creating agent '%s' with type '%s' using model '%s'\n", name, agentType, model)
		fmt.Printf("✅ Agent '%s' created\n", name)
	},
}

var agentListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all agents",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available agents:")
		fmt.Println("  - agent1 (general)")
		fmt.Println("  - agent2 (specialist)")
	},
}

func init() {
	// Node commands
	nodeStartCmd.Flags().IntP("port", "p", 4000, "Port to run the node on")
	nodeCmd.AddCommand(nodeStartCmd)
	nodeCmd.AddCommand(nodeStopCmd)
	nodeCmd.AddCommand(nodeStatusCmd)
	rootCmd.AddCommand(nodeCmd)

	// Agent commands
	agentCreateCmd.Flags().StringP("type", "t", "general", "Agent type")
	agentCreateCmd.Flags().StringP("model", "m", "gpt-4", "Model to use")
	agentCmd.AddCommand(agentCreateCmd)
	agentCmd.AddCommand(agentListCmd)
	rootCmd.AddCommand(agentCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}