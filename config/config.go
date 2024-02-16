package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

const (
	envDevelopment = "dev"
	envProduction  = "prod"
)

type Env struct {
	// Some of the allowed values are INFO and DEBUG. Default value is INFO
	LogLevel string `envconfig:"LOG_LEVEL" default:"INFO"`

	// Env is used to determine whether current setting is dev or prd.
	Env string `envconfig:"ENV" required:"true"`

	// Port for http server
	HTTPPort int `envconfig:"PORT" default:"8080"`

	// Database config
	DBName     string `envconfig:"DB_NAME" required:"true"`
	DBUser     string `envconfig:"DB_USER" required:"true"`
	DBPassword string `envconfig:"DB_PASSWORD" required:"true"`
	DBPort     string `envconfig:"DB_PORT" default:"5432"`
	DBHost     string `envconfig:"DB_HOST" default:"localhost"`
}

func (e *Env) IsProduction() bool {
	return e.Env == envProduction
}

func (e *Env) validate() error {
	checks := []struct {
		bad    bool
		errMsg string
	}{
		{
			e.Env != envDevelopment && e.Env != envProduction,
			fmt.Sprintf("invalid env is specified: %q", e.Env),
		},
	}

	for _, check := range checks {
		if check.bad {
			return errors.Errorf(check.errMsg)
		}
	}

	return nil
}

func ReadFromEnv() (*Env, error) {
	var env Env
	if err := envconfig.Process("", &env); err != nil {
		return nil, errors.Wrap(err, "failed to process envconfig")
	}

	if err := env.validate(); err != nil {
		return nil, errors.Wrap(err, "failed to validate")
	}

	return &env, nil
}
