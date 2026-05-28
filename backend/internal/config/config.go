package config

import "os"

type Config struct {
	Port     string
	MongoURI string
	MongoDB  string
}

func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbName := os.Getenv("MONGODB_DB")
	if dbName == "" {
		dbName = "store-admin"
	}

	return Config{
		Port:     port,
		MongoURI: os.Getenv("MONGODB_URI"),
		MongoDB:  dbName,
	}
}
