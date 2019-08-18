package tfgen

type Var struct {
	Name        string `yaml:"name" hcl:"name,label"`
	Description string `yaml:"description" hcl:"description,omitempty"`
	Default     string `yaml:"default" hcl:"default,omitempty"`
}
