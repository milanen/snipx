package config

type Config struct {
	Network  Network  `yaml:"network"`
	Webhook  Webhook  `yaml:"webhook"`
	Target   Target   `yaml:"target"`
	Input    Input    `yaml:"input"`
	Engine   Engine   `yaml:"engine"`
	Detector Detector `yaml:"detector"`
}

type Network struct {
	EnableGateway  bool `yaml:"enableGateway"`
	MaxThreads     int  `yaml:"maxThreads"`
	RequestTimeout int  `yaml:"requestTimeout"`
	RateLimit      int  `yaml:"rateLimit"`
}

type Webhook struct {
	Enabled  bool   `yaml:"enabled"`
	Url      string `yaml:"url"`
	MaxTries int    `yaml:"maxTries"`
}

type Target struct {
	Method   string            `yaml:"method"`
	Template string            `yaml:"template"`
	Headers  map[string]string `yaml:"headers"`
	Body     string            `yaml:"body"`
}

type Input struct {
	Path   string   `yaml:"path"`
	Inline []string `yaml:"inline"`
}

type Engine struct {
	RetryOnError    bool `yaml:"retryOnError"`
	FollowRedirects bool `yaml:"followRedirects"`
	MaxRetries      int  `yaml:"maxRetries"`
}

type Detector struct {
	MatchStatus []int    `yaml:"matchStatus"`
	MatchWords  []string `yaml:"matchWords"`
	ExcludeSize []int    `yaml:"excludeSize"`
}