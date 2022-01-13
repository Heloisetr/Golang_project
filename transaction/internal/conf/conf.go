package conf

type Configuration struct {
	Host        string     `yaml:"host"`
	Port        string     `yaml:"port"`
	Credentials Credential `yaml:"credential"`
}

type Credential struct {
	APIKey string `yaml:"api_key"`
}
