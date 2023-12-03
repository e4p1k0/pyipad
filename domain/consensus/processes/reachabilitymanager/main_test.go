package reachabilitymanager_test

import (
	"os"
	"testing"

	"github.com/e4p1k0/pyipad/infrastructure/logger"
)

const logLevel = logger.LevelWarn

func TestMain(m *testing.M) {
	logger.SetLogLevels(logLevel)
	logger.InitLogStdout(logLevel)
	os.Exit(m.Run())
}
