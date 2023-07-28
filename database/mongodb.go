package database

type MongoDBConfig struct {
	DBURL string
}

func BuildMongoConfigs() *MongoDBConfig {
	return &MongoDBConfig{}
}

type MongoClient struct {
}

func Connect(conf *MongoDBConfig) {

}
