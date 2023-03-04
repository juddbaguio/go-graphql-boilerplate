package config

type Mongo struct {
	URI string `env:"URI" envDefault:"mongodb://localhost:27017"`
}
