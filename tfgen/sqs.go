package tfgen

type SQS struct {
	Name                    string            `yaml:"queueName" hcl:"name"`
	DelaySeconds            int64             `yaml:"delaySeconds" hcl:"delay_seconds,omitempty"`
	MaxMessageSize          int64             `yaml:"maxMessageSize" hcl:"max_message_size,omitempty"`
	MessageRetentionSeconds int64             `yaml:"messageRetentionSeconds" hcl:"message_retention_seconds,omitempty"`
	ReceiveWaitTimeSeconds  int64             `yaml:"receiveWaitTimeSeconds" hcl:"receive_wait_time_seconds,omitempty"`
	RedrivePolicy           string            `yaml:"redrivePolicy" hcl:"redrive_policy,omitempty"`
	Tags                    map[string]string `yaml:"tags" hcl:"tags,omitempty"`
}
