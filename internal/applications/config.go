package applications

import (
	"fmt"
	"log"
	"sync"
	"trading-bot/internal/common"
	"trading-bot/internal/consts"
)

var (
	once sync.Once
	_cfg *Config
)

func NewConfig() *Config {
	fpath := []string{consts.ConfigPath}
	once.Do(func() {
		c, err := readCfg("app.yaml", fpath...)
		if err != nil {
			log.Fatal(err)
		}

		_cfg = c
	})

	return _cfg
}

type Config struct {
	App       *Common   `yaml:"app" json:"app"`
	Exchanges Exchanges `yaml:"exchanges" json:"exchanges"`
}

type Common struct {
	AppName  string `yaml:"name" json:"name"`
	Debug    bool   `yaml:"debug" json:"debug"`
	Timezone string `yaml:"timezone" json:"timezone"`
	Env      string `yaml:"env" json:"env"`
	Port     int    `yaml:"port" json:"port"`
}

func readCfg(fname string, ps ...string) (*Config, error) {
	var cfg *Config
	var errs []error

	for _, p := range ps {
		f := fmt.Sprint(p, fname)

		err := common.ReadFromYAML(f, &cfg)
		if err != nil {
			errs = append(errs, fmt.Errorf("file %s error %s", f, err.Error()))
			continue
		}
		break
	}

	if cfg == nil {
		return nil, fmt.Errorf("file config parse error %v", errs)
	}

	return cfg, nil
}

type Exchanges struct {
	Binance Binance `yaml:"binance" json:"binance"`
}

type (
	Binance struct {
		Futures Futures `yaml:"futures" json:"futures"`
	}

	Futures struct {
		BaseUrl          string `yaml:"base_url" json:"base_url"`
		PathConnectivity string `yaml:"path_connectivity" json:"path_connectivity"`
	}
)
