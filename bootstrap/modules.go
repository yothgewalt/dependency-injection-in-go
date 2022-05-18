package bootstrap

import (
	"github.com/yongyuth-chuankhuntod/libraries"
	"go.uber.org/fx"
)

var CommomModules = fx.Option(
	libraries.Module,
)
