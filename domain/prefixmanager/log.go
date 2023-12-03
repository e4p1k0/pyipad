package prefixmanager

import (
	"github.com/e4p1k0/pyipad/infrastructure/logger"
	"github.com/e4p1k0/pyipad/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
