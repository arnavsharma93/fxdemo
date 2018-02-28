package fxdemofx

import (
	"github.com/arnavsharma93/fxdemo/logfx"
	"github.com/arnavsharma93/fxdemo/serverfx"
	"go.uber.org/fx"
)

// Module provides common modules which can be picked up by other projects
var Module = fx.Options(
	logfx.Module,
	serverfx.Module,
)
