package tfgen

type SnsSubscription struct {
	TopicArn string `yaml:"topicArn" hcl:"topic_arn"`
	Protocol string `yaml:"protocol" hcl:"protocol"`
	Endpoint string `yaml:"endpoint" hcl:"endpoint"`
}
