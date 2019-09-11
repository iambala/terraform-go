package tfgen

type AwsCloudwatchMetricAlarm struct {
	AlarmName               string   `yaml:"alarmName" hcl:"alarm_name,omitempty"`
	ComparisonOperator      string   `yaml:"comparisonOperator" hcl:"comparison_operator,omitempty"`
	EvaluationPeriods       string   `yaml:"evaluationPeriods" hcl:"evaluation_periods,omitempty"`
	MetricName              string   `yaml:"metricName" hcl:"metric_name,omitempty"`
	Namespace               string   `yaml:"namespace" hcl:"namespace,omitempty"`
	Period                  string   `yaml:"period" hcl:"period,omitempty"`
	Statistic               string   `yaml:"statistic" hcl:"statistic,omitempty"`
	Threshold               string   `yaml:"threshold" hcl:"threshold,omitempty"`
	AlarmDescription        string   `yaml:"alarmDescription" hcl:"alarm_description,omitempty"`
	InsufficientDataActions []string `yaml:"insufficientDataActions" hcl:"insufficient_data_actions,omitempty`
}
