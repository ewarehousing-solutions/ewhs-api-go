package ewhs

type Config struct {
	Username     string
	Password     string
	WmsCode      string
	CustomerCode string
	Testing      bool
}

func NewConfig(username string, password string, wmsCode string, customerCode string, testing bool) *Config {
	return &Config{
		Username:     username,
		Password:     password,
		WmsCode:      wmsCode,
		CustomerCode: customerCode,
		Testing:      testing,
	}
}
