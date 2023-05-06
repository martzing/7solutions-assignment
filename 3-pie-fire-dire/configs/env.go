package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Environment struct {
}

func NewEnvironment() *Environment {
	env := Environment{}
	godotenv.Load()

	return &env
}

func (env *Environment) GetString(key string) *string {
	s := os.Getenv(key)
	return &s
}

func (env *Environment) GetInt(key string) (*int, error) {
	s := os.Getenv(key)

	i, err := strconv.Atoi(s)

	if err != nil {
		return nil, err
	}

	return &i, nil
}

func (env *Environment) GetBool(key string) (*bool, error) {
	s := os.Getenv(key)

	i, err := strconv.ParseBool(s)

	if err != nil {
		return nil, err
	}

	return &i, nil
}
