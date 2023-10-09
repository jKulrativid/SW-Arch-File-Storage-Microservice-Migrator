package config

import envParser "github.com/Netflix/go-env"

type Environment struct {
	PORT                          uint   `env:"PORT"`
	MINIO_ENDPOINT                string `env:"MINIO_ENDPOINT"`
	MINIO_BUCKET                  string `env:"MINIO_BUCKET"`
	MINIO_ACCESS_KEY              string `env:"MINIO_ACCESS_KEY"`
	MINIO_SECRET_KEY              string `env:"MINIO_SECRET_KEY"`
	SUBJECT_MICROSERVICE_ENDPOINT string `env:"SUBJECT_MICROSERVICE_ENDPOINT"`
}

var Env Environment

func SetupEnvironment() {
	envParser.UnmarshalFromEnviron(&Env)
}
