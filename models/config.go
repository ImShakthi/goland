package models

type Config struct {
	RBMS RBMS `toml:"rdbms"`
}

type RBMS struct {
	Database Database `toml:"database"`
	User     User     `toml:"user"`
}

type Database struct {
	Name string `toml:"name"`
	Port int    `toml:"port"`
}

type User struct {
	Name     string `toml:"name"`
	Password string `toml:"password"`
}
