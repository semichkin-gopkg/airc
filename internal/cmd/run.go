package cmd

import (
	"context"

	"github.com/cosmtrek/air/runner"
)

func run(ctx context.Context, configPath string) error {
	config, err := runner.InitConfig(configPath)
	if err != nil {
		return err
	}

	r, err := runner.NewEngineWithConfig(config, false)
	if err != nil {
		return err
	}

	go func() {
		<-ctx.Done()
		r.Stop()
	}()

	r.Run()

	return nil
}
