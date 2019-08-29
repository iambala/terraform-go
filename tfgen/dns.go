package tfgen

type Route53Record struct {
	ZoneID  string `yaml:"zoneId" hcl:"zone_id",omitempty`
	Name    string `yaml:"name" hcl:"name,omitempty"`
	Type    int64  `yaml:"type" hcl:"type,omitempty"`
	TTL     int64  `yaml:"ttl" hcl:"ttl,omitempty"`
	Records string `yaml:"records" hcl:"records,omitempty"`
	Alias   Alias  `yaml:"alias" hcl:"alias,block"`
}

type Alias struct {
	Name                 string `yaml:"name" hcl:"name,omitempty"`
	ZoneID               string `yaml:"zoneId" hcl:"zone_id",omitempty`
	EvaluateTargetHealth bool   `yaml:"evaluateTargetHealth" hcl:"evaluate_target_health",omitempty`
}
