package config


type Config struct{
	Server struct {
		Host string `yaml:"host"`
		Port string   `yaml:"port"`
	}	`yaml:"server"`

	Cors struct {
		AllowedOrigins string `yaml:"allowed_origins"`
		AllowedMethods string `yaml:"allowed_methods"`
		AllowedHeaders string `yaml:"allowed_headers"`
	} 	`yaml:"cors"`
}
