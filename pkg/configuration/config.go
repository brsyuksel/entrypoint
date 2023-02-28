package configuration

type TLSConf struct {
	Provider string            `yaml:"provider"`
	Options  map[string]string `yaml:"options"`
}

type Server struct {
	Port int     `yaml:"port"`
	TLS  TLSConf `yaml:"tls"`
}

type Metrics struct {
	Expose bool   `yaml:"expose"`
	Path   string `yaml:"path,omitempty"`
}

type Target struct {
	URL    string `yaml:"url"`
	Weight int    `yaml:"weight,omitempty"`
}

type Health struct {
	Path     string `yaml:"path"`
	Interval int    `yaml:"interval"`
	Window   int    `yaml:"window"`
}

type Upstream struct {
	BalancingPolicy string   `yaml:"balancing"`
	Targets         []Target `yaml:"targets"`
	Health          Health   `yaml:"health,omitempty"`
	RetryPolicyName string   `yaml:"retry,omitempty"`
}

type VerifyOptions struct {
	List         []string `yaml:"list,omitempty"`
	Secret       string   `yaml:"secret,omitempty"`
	UpstreamName string   `yaml:"upstream,omitempty"`
	Path         string   `yaml:"path,omitempty"`
}

type AuthenticationPolicy struct {
	Type   string        `yaml:"type"`
	Verify VerifyOptions `yaml:"verify"`
}

type RetryPolicy struct {
	Policy      string `yaml:"policy"`
	MaxAttempts int    `yaml:"max_attempts"`
	Interval    int    `yaml:"interval,omitempty"`
}

type EntrypointAuthOptions struct {
	Matches []string `yaml:"matches,omitempty"`
	Allowed []string `yaml:"allowed,omitempty"`
}

type EntrypointAuth struct {
	With    string                `yaml:"with"`
	Options EntrypointAuthOptions `yaml:"options,omitempty"`
}

type Entrypoint struct {
	UpstreamName string         `yaml:"upstream"`
	Auth         EntrypointAuth `yaml:"auth,omitempty"`
	Compression  bool           `yaml:"compression,omitempty"`
	RateLimiter  string         `yaml:"rate_limit,omitempty"`
}

type Config struct {
	Entrypoints            map[string]Entrypoint           `yaml:"entrypoints"`
	RetryPolicies          map[string]RetryPolicy          `yaml:"retry_policies,omitempty"`
	AuthenticationPolicies map[string]AuthenticationPolicy `yaml:"authentication_policies,omitempty"`
	Upstreams              map[string]Upstream             `yaml:"upstreams"`
	Metrics                Metrics                         `yaml:"metrics,omitempty"`
	Server                 Server                          `yaml:"server"`
}
