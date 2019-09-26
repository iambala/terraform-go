package tfgen

type AwsCloudwatchMetricAlarm struct {
	AlarmName               string        `yaml:"alarmName" hcl:"alarm_name,omitempty"`
	ComparisonOperator      string        `yaml:"comparisonOperator" hcl:"comparison_operator,omitempty"`
	EvaluationPeriods       string        `yaml:"evaluationPeriods" hcl:"evaluation_periods,omitempty"`
	MetricName              string        `yaml:"metricName" hcl:"metric_name,omitempty"`
	Namespace               string        `yaml:"namespace" hcl:"namespace,omitempty"`
	Period                  string        `yaml:"period" hcl:"period,omitempty"`
	Statistic               string        `yaml:"statistic" hcl:"statistic,omitempty"`
	Threshold               string        `yaml:"threshold" hcl:"threshold,omitempty"`
	AlarmDescription        string        `yaml:"alarmDescription" hcl:"alarm_description,omitempty"`
	InsufficientDataActions []string      `yaml:"insufficientDataActions" hcl:"insufficient_data_actions,omitempty"`
	AlarmActions            []string      `yaml:"alarmActions" hcl:"alarm_actions,omitempty"`
	MetricQuery             []MetricQuery `yaml:"metricQuery" hcl:"metric_query,omitempty"`
	Dimensions              []Dimensions  `yaml:"dimensions"`
}

type MetricQuery struct {
	ID         string  `yaml:"id" hcl:"id,omitempty"`
	Expression string  `yaml:"expression" hcl:"expression,omitempty"`
	Label      string  `yaml:"label" hcl:"label,omitempty"`
	ReturnData string  `yaml:"returnData" hcl:"return_data,omitempty"`
	Metric     *Metric `yaml:"metric" hcl:"metric,omitempty"`
}

type Metric struct {
	MetricName string      `yaml:"metricName" hcl:"metric_name,omitempty"`
	Namespace  string      `yaml:"namespace" hcl:"namespace,omitempty"`
	Stat       string      `yaml:"stat" hcl:"stat,omitempty"`
	Unit       string      `yaml:"unit" hcl:"unit,omitempty"`
	Dimensions *Dimensions `yaml:"dimensions" hcl:"dimensions,omitempty"`
}

type Dimensions struct {
	DimensionName string
	Value         string
}
