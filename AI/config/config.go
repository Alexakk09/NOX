package config

type Config struct {
	APIURL     string
	SearXNGURL string
}

func Load() Config {
	return Config{
		APIURL:     "http://localhost:8080",
		SearXNGURL: "http://localhost:8888",
	}
}