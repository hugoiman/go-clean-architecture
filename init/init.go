package init

import (
	"go-clean-architecture/config"
	"os"
)

func RunInit() {
	if config.ReadEnv() != nil {
		os.Exit(0)
	}
}
