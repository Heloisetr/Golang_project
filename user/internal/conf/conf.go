package conf

type Configuration struct {
	Host         string       `yaml:"host"`
	Port         string       `yaml:"port"`
	Credential   Credential   `yaml:"credential"`
	OrderService OrderService `yaml:"order_service"`
}

type Credential struct {
	ApiKey string `yaml:"api_key"`
}

type OrderService struct {
	URL    string `yaml:"url"`
	APIKey string `yaml:"api_key"`
}
