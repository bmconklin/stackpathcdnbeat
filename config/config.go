// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Credentials struct {
	Alias  string `config:"alias"`
	Key    string `config:"key"`
	Secret string `config:"secret"`
	Sites  string `config:"site_id"`
}

type Config struct {
	Period      time.Duration `config:"period"`
	Endpoint    string        `config:"endpoint"`
	Path        string        `config:"path"`
	Credentials Credentials   `config:"api_credentials"`
	Start       string        `config:"start_time"`
	End         string        `config:"end_time"`
	Delay       time.Duration `config:"query_delay"`
}

var DefaultConfig = Config{
	Period:   1 * time.Second,
	Endpoint: "https://rws.maxcdn.com",
	Start:    time.Now().UTC().Add(-1 * time.Hour).Format(time.RFC3339),
	End:      time.Now().UTC().Format(time.RFC3339),
	Delay:    20 * time.Second,
}
