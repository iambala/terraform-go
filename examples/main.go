package main

import (
	"encoding/json"
	"fmt"
	"github.com/iambala/terraform-go/tfgen"
)

func main() {
	tf := tfgen.NewTf()
	var sqs tfgen.SQS

	sqs.Name = "test"
	sqs.DelaySeconds = 12
	sqs.Tags = make(map[string]string)
	sqs.Tags["myTag"] = "myTagValue"
	sqs.Tags["myTag1"] = "myTagValue1"
	var ddb tfgen.DynamoDB
	ddb.Name = "mytable"
	ddb.GlobalSecondaryIndex = []tfgen.GlobalSecondaryIndex{
		tfgen.GlobalSecondaryIndex{
			Name: "test",
		},
		tfgen.GlobalSecondaryIndex{
			Name: "test",
		},
	}
	ddb.Attributes = []tfgen.Attribute{
		tfgen.Attribute{
			Name: "test",
			Type: "sdf",
		},
		tfgen.Attribute{
			Name: "rest",
			Type: "sdf",
		},
	}
	ddb.Tags = make(map[string]string)
	ddb.Tags["mytag1"] = "sdafads"
	ddb.Tags["mytag1"] = "sdfsf"

	tf.AddResource(sqs)
	tf.AddResource(ddb)
	jsont, _ := json.Marshal(ddb)
	fmt.Println(string(jsont))

	myVar := tfgen.Var{
		Name:    "region",
		Default: "eu-west-1",
	}
	account_id := tfgen.Var{
		Name:    "account_id",
		Default: "12345678910",
	}

	provider := tfgen.Provider{
		Name:   "aws",
		Region: "eu-west-1",
	}
	snssub := tfgen.SnsSubscription{
		TopicArn: "arn:aws:sns:${var.region}:${var.account_id}:my_corporate_topic",
		Endpoint: "arn:aws:sqs:${var.region}:${var.account_id}:queuename",
		Protocol: "sqs",
	}

	tf.AddResource(myVar)
	tf.AddResource(provider)
	tf.AddResource(account_id)
	tf.AddResource(snssub)

	tf.SaveFile()
	cleanupFile()
}
