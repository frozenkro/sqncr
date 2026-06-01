package globals

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type Environments int

const (
	Test Environments = iota
	Dev
	Prod
)

const AppName = "sqncr"

var (
	Environment Environments
	DbName      string
	AppDataDir  string
	AppStateDir string
)

func Init() error {
	if err := initXDGDirs(); err != nil {
		return err
	}
	if isDev() {
		return initEnv(Dev)
	}
	return initEnv(Prod)
}

func InitTest() error {
	tmpDir, err := os.MkdirTemp("", "sqncr-test-*")
	if err != nil {
		return err
	}
	AppDataDir = tmpDir
	AppStateDir = tmpDir
	return initEnv(Test)
}

func initEnv(env Environments) error {
	Environment = env

	switch env {
	case Test:
		DbName = filepath.Join(AppDataDir, "test.db")
	case Dev:
		DbName = filepath.Join(AppDataDir, "dev.db")
	default:
		DbName = filepath.Join(AppDataDir, fmt.Sprintf("%s.db", AppName))
	}

	return nil
}

func isDev() bool {
	for _, v := range os.Args {
		if v == "--dev" {
			return true
		}
	}
	return false
}

func initXDGDirs() error {
	var dataDir, stateDir string

	if runtime.GOOS == "windows" {
		dataDir = filepath.Join(os.Getenv("LOCALAPPDATA"), AppName)
		stateDir = dataDir
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		dataHome := os.Getenv("XDG_DATA_HOME")
		if dataHome == "" {
			dataHome = filepath.Join(homeDir, ".local", "share")
		}
		stateHome := os.Getenv("XDG_STATE_HOME")
		if stateHome == "" {
			stateHome = filepath.Join(homeDir, ".local", "state")
		}

		dataDir = filepath.Join(dataHome, AppName)
		stateDir = filepath.Join(stateHome, AppName)
	}

	for _, dir := range []string{dataDir, stateDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	AppDataDir = dataDir
	AppStateDir = stateDir
	return nil
}
