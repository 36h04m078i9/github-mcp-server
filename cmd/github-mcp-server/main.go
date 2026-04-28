// Package main is the entry point for the GitHub MCP Server.
// It initializes the server configuration and starts the MCP server
// that exposes GitHub API functionality as tools.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/github/github-mcp-server/internal/config"
	"github.com/github/github-mcp-server/internal/server"
	"github.com/spf13/cobra"
)

var (
	// version is set at build time via ldflags.
	version = "dev"
	// commit is the git commit hash, set at build time.
	commit = "none"
)

func main() {
	if err := rootCmd().Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func rootCmd() *cobra.Command {
	var (
		token   string
		logFile string
		readOnly bool
	)

	cmd := &cobra.Command{
		Use:     "github-mcp-server",
		Short:   "GitHub MCP Server — exposes GitHub API as MCP tools",
		Version: fmt.Sprintf("%s (%s)", version, commit),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runServer(cmd.Context(), token, logFile, readOnly)
		},
	}

	cmd.PersistentFlags().StringVar(
		&token, "token", "",
		"GitHub personal access token (overrides GITHUB_PERSONAL_ACCESS_TOKEN env var)",
	)
	cmd.PersistentFlags().StringVar(
		&logFile, "log-file", "",
		"Path to log file (defaults to stderr)",
	)
	cmd.PersistentFlags().BoolVar(
		&readOnly, "read-only", false,
		"Restrict the server to read-only operations",
	)

	cmd.AddCommand(versionCmd())

	return cmd
}

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("github-mcp-server %s (%s)\n", version, commit)
		},
	}
}

func runServer(ctx context.Context, token, logFile string, readOnly bool) error {
	// Resolve token from flag or environment variable.
	if token == "" {
		token = os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	}
	if token == "" {
		return fmt.Errorf("GitHub token is required: set GITHUB_PERSONAL_ACCESS_TOKEN or use --token flag")
	}

	cfg := &config.Config{
		Token:    token,
		LogFile:  logFile,
		ReadOnly: readOnly,
	}

	srv, err := server.New(cfg)
	if err != nil {
		return fmt.Errorf("failed to create server: %w", err)
	}

	// Handle graceful shutdown on interrupt/termination signals.
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	fmt.Fprintln(os.Stderr, "GitHub MCP Server started, listening on stdio...")

	if err := srv.Serve(ctx); err != nil {
		return fmt.Errorf("server error: %w", err)
	}

	return nil
}
