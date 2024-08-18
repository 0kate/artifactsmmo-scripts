package artifactsapi

const (
	DEFAULT_ARTIFACTS_API_HOST = "api.artifactsmmo.com"
)

type Config struct {
	host  string
	token string
}

func NewDefaultConfig(token string) *Config {
	return &Config{
		host:  DEFAULT_ARTIFACTS_API_HOST,
		token: token,
	}
}

func (c *Config) Host() string {
	return c.host
}

func (c *Config) Token() string {
	return c.token
}
