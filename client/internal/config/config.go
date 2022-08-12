package config

var configSingleton *Config

type Config struct {
	serverAddress      string
	userToken          string
	httpTimeout        int // seconds
	maxIdleConnections int // seconds
}

func New() Config {
	if configSingleton == nil {
		config := Config{}
		config = initConfigByDefault(config)

		configSingleton = &config
	}

	return *configSingleton
}

func (c Config) GetServerAddress() string {
	return c.serverAddress
}

func (c Config) GetToken() string {
	return c.userToken
}

func (c *Config) SetToken(token string) {
	c.userToken = token
}

func (c Config) GetHttpTimeout() int {
	return c.httpTimeout
}

func (c Config) GetMaxIdleConnections() int {
	return c.maxIdleConnections
}

func initConfigByDefault(config Config) Config {
	if len(config.serverAddress) < 1 {
		config.serverAddress = "http://0.0.0.0:8081"
	}

	if len(config.userToken) < 1 {
		config.userToken = ""
	}

	if config.httpTimeout == 0 {
		config.httpTimeout = 3
	}

	if config.maxIdleConnections == 0 {
		config.maxIdleConnections = 20
	}

	return config
}
