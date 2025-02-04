package ewhs

type Config struct {
	Username     string
	Password     string
	WmsCode      string
	CustomerCode string
	Testing      bool
	Domain 		 string
	App			 string
	AppVersion	 string
}

func NewConfig(username string, password string, wmsCode string, customerCode string, testing bool, domain string, app string, appVersion string) *Config {
	return &Config{
		Username:     username,
		Password:     password,
		WmsCode:      wmsCode,
		CustomerCode: customerCode,
		Testing:      testing,
		Domain:	      domain,
		App:		  app,
		AppVersion:   appVersion,
	}
}
