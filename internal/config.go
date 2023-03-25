package apiserver

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	PG       string
	Mongo    *Mongo
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		PG:       "",
		Mongo: &Mongo{
			ConnectionString: "",
			DatabaseName:     "",
			CollectionName:   "",
		},
	}
}

type Mongo struct {
	ConnectionString, DatabaseName, CollectionName string
}
