package tfgen

import (
	"bufio"
	"fmt"
	"github.com/gigawattio/awsarn"
	"github.com/hashicorp/hcl2/hclwrite"
	"github.com/iambala/terraform-go/gohcl"
	"os"
)

type tf struct {
	Main     *hclwrite.File
	MainBody *hclwrite.Body
}

func NewTf() tf {
	var t tf
	t.Main = hclwrite.NewEmptyFile()
	t.MainBody = t.Main.Body()
	return t
}

func (t tf) AddResourceBlock(block *hclwrite.Block) {
	if block != nil {
		t.MainBody.AppendBlock(block)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func extractResourceName(arn string) string {
	components, err := awsarn.Parse(arn)
	if err != nil {
		panic(err)
	}
	return components.Resource
}

func (t tf) SaveFile() {
	f, err := os.Create("main.tf")
	check(err)
	defer f.Close()
	fw := bufio.NewWriter(f)
	t.Main.WriteTo(fw)
	fw.Flush()
}

func (t tf) AddResource(resource interface{}) {
	var block *hclwrite.Block
	var body *hclwrite.Body
	switch resource.(type) {
	case SQS:
		block = t.MainBody.AppendNewBlock("resource", []string{"aws_sqs_queue", resource.(SQS).Name})
		body = block.Body()
		gohcl.EncodeIntoBody(resource.(SQS), body)
	case DynamoDB:
		block = t.MainBody.AppendNewBlock("resource", []string{"aws_dynamodb_table", resource.(DynamoDB).Name})
		body = block.Body()
		gohcl.EncodeIntoBody(resource.(DynamoDB), body)
	case SNS:
		block = t.MainBody.AppendNewBlock("resource", []string{"aws_sns_topic", resource.(SNS).Name})
		body = block.Body()
		gohcl.EncodeIntoBody(resource.(SNS), body)
	case SnsSubscription:
		name := fmt.Sprintf("%s-%s", extractResourceName(resource.(SnsSubscription).TopicArn), extractResourceName(resource.(SnsSubscription).Endpoint))
		block = t.MainBody.AppendNewBlock("resource", []string{"aws_sns_topic_subscription", name})
		body = block.Body()
		gohcl.EncodeIntoBody(resource.(SnsSubscription), body)
	default:
		fmt.Println("error")
	}
}
