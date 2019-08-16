package tfgen

import (
	"bufio"
	"fmt"
	"github.com/hashicorp/hcl2/hclwrite"
	"github.com/iambala/terraform-go/gohcl"
	"os"
)

type tf struct {
	F        *hclwrite.File
	RootBody *hclwrite.Body
}

func NewTf() tf {
	var tf tf
	tf.F = hclwrite.NewEmptyFile()
	tf.RootBody = tf.F.Body()
	return tf
}

func (t tf) AddResourceBlock(block *hclwrite.Block) {
	if block != nil {
		t.RootBody.AppendBlock(block)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (t tf) SaveFile() {
	f, err := os.Create("main.tf")
	check(err)
	defer f.Close()
	fw := bufio.NewWriter(f)
	t.F.WriteTo(fw)
	fw.Flush()
}

func (t tf) AddResource(resource interface{}) {
	var block *hclwrite.Block
	var body *hclwrite.Body
	switch resource.(type) {
	case SQS:
		block = t.RootBody.AppendNewBlock("resource", []string{"aws_sqs_queue", resource.(SQS).Name})
		body = block.Body()
		gohcl.EncodeIntoBody(resource.(SQS), body)
	case DynamoDB:
		block = t.RootBody.AppendNewBlock("resource", []string{"aws_dynamodb_table", resource.(DynamoDB).Name})
		body = block.Body()
		gohcl.EncodeIntoBody(resource.(DynamoDB), body)
	case SNS:
		block = t.RootBody.AppendNewBlock("resource", []string{"aws_sns_topic", resource.(SNS).Name})
		body = block.Body()
		gohcl.EncodeIntoBody(resource.(SNS), body)
	case SnsSubscription:
		// TODO Trim topic name from ARN
		block = t.RootBody.AppendNewBlock("resource", []string{"aws_sns_topic_subscription", resource.(SnsSubscription).TopicArn})
		body = block.Body()
		gohcl.EncodeIntoBody(resource.(SnsSubscription), body)
	default:
		fmt.Println("error")
	}
}
