package config

type Config struct {
	Port        int    `env:"PORT" envDefault:"8080"`
	DBURI       string `env:"DB_URI" envDefault:"mongodb://localhost:27017"`
	RAEXTAPIURL string `env:"RAEXTAPI_URL" envDefault:"https://iosite.reclameaqui.com.br/raichu-io-site-v1"`
}
