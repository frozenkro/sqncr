package main

import (
	"fmt"
	"os"

	"github.com/frozenkro/sqncr/internal/mcp"
	"github.com/frozenkro/sqncr/internal/tui"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	subcommand := os.Args[1]
	args := os.Args[2:]

	switch subcommand {
	case "mcp":
		mcp.Run(args)
	case "tui":
		tui.Run(args)
	case "help", "--help", "-h":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "Unknown subcommand: %s\n\n", subcommand)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`sqncr - LLM Project Sequencer

Usage:
  sqncr <subcommand> [flags]

Subcommands:
  mcp    Start the MCP server (stdio by default, --http [port] for HTTP mode)
  tui    Start the terminal UI

Examples:
  sqncr mcp              Start MCP server in stdio mode
  sqncr mcp --http 8080  Start MCP server in HTTP mode on port 8080
  sqncr mcp --dev        Start MCP server in dev mode
  sqncr tui              Start the terminal UI`)
}