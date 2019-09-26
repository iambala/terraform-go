package tfgen

import (
	"bufio"
	"fmt"
	"github.com/hashicorp/hcl2/hclwrite"
	"github.com/iambala/terraform-go/gohcl"
	"github.com/zclconf/go-cty/cty"
	"os"
)

type Tf struct {
	Main       *hclwrite.File
	MainBody   *hclwrite.Body
	Var        *hclwrite.File
	VarBody    *hclwrite.Body
	Output     *hclwrite.File
	OutputBody *hclwrite.Body
}

func NewTf() Tf {
	var t Tf
	t.Main = hclwrite.NewEmptyFile()
	t.MainBody = t.Main.Body()
	t.Var = hclwrite.NewEmptyFile()
	t.VarBody = t.Var.Body()
	t.Output = hclwrite.NewEmptyFile()
	t.OutputBody = t.Output.Body()
	return t
}

func (t Tf) AddResourceBlock(block *hclwrite.Block) {
	if block != nil {
		t.MainBody.AppendBlock(block)
	}
}

func (t Tf) SaveFile() {
	//create and save Main.tf
	mf, err := os.Create("main.tf")
	check(err)
	defer mf.Close()
	mfw := bufio.NewWriter(mf)
	t.Main.WriteTo(mfw)
	mfw.Flush()

	//create and save Variables.tf
	vf, err := os.Create("variables.tf")
	check(err)
	defer vf.Close()
	vfw := bufio.NewWriter(vf)
	t.Var.WriteTo(vfw)
	vfw.Flush()

	//create and save Outputs.tf
	of, err := os.Create("outputs.tf")
	check(err)
	defer of.Close()
	ofw := bufio.NewWriter(of)
	t.Output.WriteTo(ofw)
	ofw.Flush()
}

func (t Tf) AddResource(resource interface{}) {
	var block *hclwrite.Block
	var body *hclwrite.Body
	switch resource.(type) {
	case SQS:
		block = t.MainBody.AppendNewBlock("resource", []string{"aws_sqs_queue", resourceName(resource.(SQS).Name)})
		body = block.Body()
		gohcl.EncodeIntoBody(resource.(SQS), body)
	case DynamoDB:
		block = t.MainBody.AppendNewBlock("resource", []string{"aws_dynamodb_table", resourceName(resource.(DynamoDB).Name)})
		body = block.Body()
		gohcl.EncodeIntoBody(resource.(DynamoDB), body)
	case SNS:
		block = t.MainBody.AppendNewBlock("resource", []string{"aws_sns_topic", resourceName(resource.(SNS).Name)})
		body = block.Body()
		gohcl.EncodeIntoBody(resource.(SNS), body)
	case SnsSubscription:
		name := fmt.Sprintf("%s-sub-%s", extractResourceName(resource.(SnsSubscription).TopicArn), extractResourceName(resource.(SnsSubscription).Endpoint))
		block = t.MainBody.AppendNewBlock("resource", []string{"aws_sns_topic_subscription", name})
		body = block.Body()
		gohcl.EncodeIntoBody(resource.(SnsSubscription), body)
	case Var:
		block = t.VarBody.AppendNewBlock("variable", []string{resource.(Var).Name})
		body = block.Body()
		gohcl.EncodeIntoBody(resource.(Var), body)
	case Provider:
		block = t.MainBody.AppendNewBlock("provider", []string{resource.(Provider).Name})
		body = block.Body()
		gohcl.EncodeIntoBody(resource.(Provider), body)
	case Route53Record:
		block = t.MainBody.AppendNewBlock("resource", []string{"aws_route53_record", resourceName(resource.(Route53Record).Name)})
		body = block.Body()
		gohcl.EncodeIntoBody(resource.(Route53Record), body)
	case AwsCloudwatchMetricAlarm:
		block = t.MainBody.AppendNewBlock("resource", []string{"aws_cloudwatch_metric_alarm", resourceName(resource.(AwsCloudwatchMetricAlarm).AlarmName)})
		gohcl.EncodeIntoBody(resource.(AwsCloudwatchMetricAlarm), body)
		body = block.Body()
		for _, dimension := range resource.(AwsCloudwatchMetricAlarm).Dimensions {
			dimensions := body.AppendNewBlock("dimensions", nil)
			dimensionBody := dimensions.Body()
			dimensionBody.SetAttributeValue(dimension.DimensionName, cty.StringVal(dimension.Value))

		}
	default:
		fmt.Println("error")
	}
}
