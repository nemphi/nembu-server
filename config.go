package main

// Config for this instance of nembu-server
type Config struct {
	DB DBConfig
}

// DBConfig to use for this instance of nembu-server
type DBConfig struct {
	DBHost     string
	DBUsername string
	DBPassword string
}
