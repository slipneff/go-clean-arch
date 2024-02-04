package main

import (
	"fmt"

	"github.com/4kayDev/logger/log"
	"github.com/slipneff/go-clean-arch/internal/di"
	"github.com/slipneff/go-clean-arch/internal/pkg/seed"
	"github.com/slipneff/go-clean-arch/internal/utils/config"
	"github.com/slipneff/go-clean-arch/internal/utils/flags"
)

func main() {
	flags := flags.MustParseFlags()
	cfg := config.MustLoadConfig(flags.EnvMode, flags.ConfigPath)
	container := di.New(cfg)

	for _, seed := range seed.All() {
		if err := seed.Run(container.GetDB()); err != nil {
			log.Error(err, fmt.Sprintf("Running seed '%s'", seed.Name), "Seeder")
		}
	}
	log.Info("Seeding Success", "Seeder")
}
