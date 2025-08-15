package config

type ExpenseTrackerDBConfig struct {
	DbName     string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
}

// Define a config struct to hold all the configuration settings for our application.
type Config struct {
	Port int
	Env  string
}
