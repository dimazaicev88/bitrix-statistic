package config

import (
	"os"
	"unicode/utf8"
)

const (
	StorageHost     = "STORAGE_HOST"
	StorageUser     = "STORAGE_USER"
	StoragePassword = "STORAGE_PASSWORD"
	StorageDbName   = "STORAGE_DB_NAME"
	StoragePort     = "STORAGE_PORT"
)

type ServerEnvConfig struct{}

func (s ServerEnvConfig) StorageHost() string {
	return os.Getenv(StorageHost)
}

func (s ServerEnvConfig) StorageUser() string {
	return os.Getenv(StorageUser)
}

func (s ServerEnvConfig) StoragePassword() string {
	return os.Getenv(StoragePassword)
}

func (s ServerEnvConfig) StorageDbName() string {
	return os.Getenv(StorageDbName)
}

func (s ServerEnvConfig) StoragePort() string {
	if utf8.RuneCountInString(os.Getenv(StoragePort)) == 0 {
		return "3306"
	}
	return os.Getenv(StoragePort)
}

func (s ServerEnvConfig) ValidateStorageParams() {
	switch {
	case len(s.StorageHost()) == 0:
		panic(StorageHost + " is not set")

	case len(s.StorageUser()) == 0:
		panic(StorageUser + " is not set")

	//case len(s.StoragePassword()) == 0:
	//	panic(StoragePassword + " is not set")

	case len(s.StorageDbName()) == 0:
		panic(StorageDbName + " is not set")
	}
}
