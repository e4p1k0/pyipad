package main

import (
	"github.com/e4p1k0/pyipad/infrastructure/logger"
	"github.com/e4p1k0/pyipad/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("JSTT")
	spawn      = panics.GoroutineWrapperFunc(log)
)
