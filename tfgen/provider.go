package tfgen

type Provider struct {
	Name                  string `yaml:"name" hcl:"name,label"`
	Region                string `yaml:"region" hcl:"region,omitempty"`
	Version               string `yaml:"version" hcl:"version,omitempty"`
	AccessKey             string `yaml:"accessKey" hcl:"access_key,omitempty"`
	SecretKey             string `yaml:"secretKey" hcl:"secret_key,omitempty"`
	SharedCredentialsFile string `yaml:"sharedCredentialsFile" hcl:"shared_credentials_file,omitempty"`
	Profile               string `yaml:"profile" hcl:"profile,omitempty"`
}
