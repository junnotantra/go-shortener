package config

type (
	// Config will holds mapped key value for service configuration
	Config struct {
		Server    ServerConfig    `yaml:"server"`
		Database  DatabaseConfig  `yaml:"database"`
		Shortener ShortenerConfig `yaml:"shortener"`
		Redirect  RedirectConfig  `yaml:"redirect"`
	}

	// ServerConfig server config
	ServerConfig struct {
		HTTPPort string `yaml:"http_port"`
	}

	// DatabaseConfig db config
	DatabaseConfig struct {
		Main      DatabaseSetting `yaml:"main"`
		Statistic DatabaseSetting `yaml:"statistic"`
	}

	// DatabaseSetting specific DB setting
	DatabaseSetting struct {
		FileName string `yaml:"file_name"`
		Timeout  int64  `yaml:"timeout"`
	}

	// ShortenerConfig is config for shorting URL
	ShortenerConfig struct {
		Charset            string `yaml:"charset"`
		UniqueStringLength int    `yaml:"unique_string_length"`
	}

	// RedirectConfig is config for redirection
	RedirectConfig struct {
		BaseRedirect BaseRedirectSetting `yaml:"base_redirect"`
	}

	// BaseRedirectSetting specific rediretion setting for base url
	BaseRedirectSetting struct {
		URL    string `yaml:"url"`
		Active bool   `yaml:"active"`
	}
)
