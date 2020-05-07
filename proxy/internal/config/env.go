package config

import (
	"errors"
	"fmt"
	"os"
)

const (
	DevelopmentEnvironment = "dev"
	ProductionEnvironment  = "prod"
)

var (
	ServerConfiguration          = Server{}
	requiredEnvironemntVariables = []string{
		"ENV",
	}
	optionsForEnvironment = []string{
		DevelopmentEnvironment,
		ProductionEnvironment,
	}
)

type Server struct {
	Environment string
}

func validateRequiredVariables(options []string) error {
	for _, val := range options {
		currEnv := os.Getenv(val)
		if currEnv == " " || currEnv == "" {
			return errors.New(fmt.Sprintf("missing/invalid environment variable: %s", val))
		}
	}

	return nil
}

func validateEnvironment(selectedEnv string, options []string) error {
	for _, val := range options {
		if val == selectedEnv {
			return nil
		}
	}

	return errors.New(fmt.Sprintf("the selected environment is not supported: %s", selectedEnv))
}

func SetupServerConfiguration() (*Server, error) {
	if err := validateRequiredVariables(requiredEnvironemntVariables); err != nil {
		return nil, err
	}

	if err := validateEnvironment(os.Getenv("ENV"), optionsForEnvironment); err != nil {
		return nil, err
	}

	return &ServerConfiguration, nil
}
