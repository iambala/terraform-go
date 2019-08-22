package tfgen

type DynamoDB struct {
	Name                 string                 `yaml:"name" hcl:"name"`
	BillingMode          string                 `yaml:"billingMode" hcl:"billing_mode,omitempty"`
	ReadCapacity         int64                  `yaml:"readCapacity" hcl:"read_capacity,omitempty"`
	WriteCapacity        int64                  `yaml:"writeCapacity" hcl:"write_capacity,omitempty"`
	HashKey              string                 `yaml:"hashKey" hcl:"hash_key,omitempty"`
	RangeKey             string                 `yaml:"rangeKey" hcl:"range_key,omitempty"`
	Attributes           []Attribute            `yaml:"attributes" hcl:"attribute,block"`
	TTL                  *Ttl                   `yaml:"ttl" hcl:"ttl,block"`
	GlobalSecondaryIndex []GlobalSecondaryIndex `yaml:"globalSecondaryIndex" hcl:"global_secondary_index,block"`
	Tags                 map[string]string      `yaml:"tags" hcl:"tags,omitempty"`
}

type Attribute struct {
	Name string `yaml:"name" hcl:"name"`
	Type string `yaml:"type" hcl:"type"`
}

type Ttl struct {
	AttributeName string `yaml:"attributeName" hcl:"attribute_name"`
	Enabled       bool   `yaml:"enabled" hcl:"enabled"`
}

type GlobalSecondaryIndex struct {
	Name             string   `yaml:"name" hcl:"name"`
	HashKey          string   `yaml:"hashKey" hcl:"hash_key"`
	RangeKey         string   `yaml:"rangeKey" hcl:"range_key"`
	WriteCapacity    int64    `yaml:"writeCapacity" hcl:"write_capacity"`
	ReadCapacity     int64    `yaml:"readCapacity" hcl:"read_capacity"`
	ProjectionType   string   `yaml:"projectionType" hcl:"projection_type"`
	NonKeyAttributes []string `yaml:"nonKeyAttributes" hcl:"non_key_attributes"`
}
