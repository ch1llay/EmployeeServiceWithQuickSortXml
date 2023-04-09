package config

type Config struct {
	Port                int    `json:"port"`
	PostgresConnection  string `json:"postgresConnection"`
	MongoConnection     string `json:"mongoConnection"`
	MongoDbName         string `json:"mongoDbName"`
	MongoCollectionName string `json:"mongoCollectionName"`
}
