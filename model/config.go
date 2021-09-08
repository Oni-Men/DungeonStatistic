package model

type Config struct {
	Year    int
	Month   int
	Host    string        `json:"host"`
	Elastic ElasticConfig `json:"elastic"`
}

type ElasticConfig struct {
	Host string
}
