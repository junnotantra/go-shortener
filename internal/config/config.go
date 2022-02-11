package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

var (
	config *Config
)

// option defines configuration option
type option struct {
	configFile string
}

// Init initializes `config` from the default config file.
// use `WithConfigFile` to specify the location of the config file
func Init(opts ...Option) error {
	opt := &option{
		configFile: getDefaultConfigFile(),
	}
	for _, optFunc := range opts {
		optFunc(opt)
	}

	out, err := ioutil.ReadFile(opt.configFile)
	if err != nil {
		log.Println(err)
		return getDefaultConfig()
	}

	return yaml.Unmarshal(out, &config)
}

// Option define an option for config package
type Option func(*option)

// WithConfigFile set `config` to use the given config file
func WithConfigFile(file string) Option {
	return func(opt *option) {
		opt.configFile = file
	}
}

// getDefaultConfigFile get default config file.
// - files/etc/shortener/shortener.development.yaml in dev
// - otherwise /etc/shortener/shortener.{JTENV}.yaml
func getDefaultConfigFile() string {
	var (
		repoPath   = filepath.Join(os.Getenv("GOPATH"), "src/github.com/junnotantra/go-shortener")
		configPath = filepath.Join(repoPath, "files/etc/shortener/shortener.development.yaml")
	)

	env := os.Getenv("JTENV")
	if env == "" {
		env = "development"
	}

	if env != "development" {
		configPath = fmt.Sprintf("/etc/shortener/shortener.%s.yaml", env)
	}
	return configPath
}

// Get config
func Get() *Config {
	return config
}

// getDefaultConfig will be fallback if cannot read config from file
func getDefaultConfig() error {
	config = &Config{
		Server: ServerConfig{
			HTTPPort: ":" + os.Getenv("PORT"),
		},
		Database: DatabaseConfig{
			Main: DatabaseSetting{
				FileName: "main.db",
				Timeout:  1000,
			},
			Statistic: DatabaseSetting{
				FileName: "statistic.db",
				Timeout:  1000,
			},
		},
		Shortener: ShortenerConfig{
			Charset:            "23456789bcdfghjkmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ",
			UniqueStringLength: 7,
		},
		Redirect: RedirectConfig{
			BaseRedirect: BaseRedirectSetting{
				Active: false,
			},
		},
	}

	return nil
}
