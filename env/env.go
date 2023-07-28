package env

type Environment map[string]interface{}

func NewEnvironment() *Environment {
	var env Environment
	return &env
}
