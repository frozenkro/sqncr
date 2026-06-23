package tui

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/frozenkro/sqncr/internal/db"
	"github.com/frozenkro/sqncr/internal/globals"
	"github.com/frozenkro/sqncr/internal/services"
	"github.com/frozenkro/sqncr/internal/tui/components/projects"
	"github.com/frozenkro/sqncr/internal/tui/components/taskdetail"
	"github.com/frozenkro/sqncr/internal/tui/components/tasks"
	tuilog "github.com/frozenkro/sqncr/internal/tui/logger"
)

func InitialModel(logger *log.Logger) (Model, error) {
	tuilog.SetLogger(logger)
	ctx := context.Background()
	svc := services.NewServices()

	width := 100 // Default width
	height := 30 // Default height

	// Initialize components
	projectsModel, err := projects.NewModel(svc, ctx, width, height)
	if err != nil {
		return Model{}, err
	}

	tasksModel := tasks.NewModel(svc, width, height)
	taskDetailModel := taskdetail.NewModel(svc, width, height)

	// Create main model
	model := Model{
		Projects:   projectsModel,
		Tasks:      tasksModel,
		TaskDetail: taskDetailModel,
		ActivePane: ProjectPane,
		Services:   svc,
		Context:    ctx,
		Width:      width,
		Height:     height,
	}

	return model, nil
}

// Run starts the TUI application. Pass args after the subcommand (usually empty).
func Run(args []string) {
	if err := globals.Init(); err != nil {
		log.Fatalf("Application Initialization failed. \nError: %v\n", err.Error())
	}

	logger, closeLogFile, err := setupLogger()
	if err != nil {
		fmt.Printf("Error creating logger: %v\n", err.Error())
	}
	defer closeLogFile()

	db.Init()

	m, err := InitialModel(logger)
	if err != nil {
		log.Fatalf("Error initializing state: %v", err.Error())
	}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func setupLogger() (*log.Logger, func() error, error) {
	logPath := filepath.Join(globals.AppStateDir, "sqncrtui.log")
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, nil, err
	}

	logger := log.New(file, "", log.LstdFlags)

	return logger, file.Close, nil
}
