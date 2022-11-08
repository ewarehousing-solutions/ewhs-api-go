package ewhs

type Config struct {
	Username     string
	Password     string
	WmsCode      string
	CustomerCode string
}

func NewConfig(username string, password string, wmsCode string, customerCode string) *Config {
	return &Config{
		Username:     username,
		Password:     password,
		WmsCode:      wmsCode,
		CustomerCode: customerCode,
	}
}
