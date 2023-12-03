package grpcclient

import (
	"github.com/e4p1k0/pyipad/infrastructure/logger"
	"github.com/e4p1k0/pyipad/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)
