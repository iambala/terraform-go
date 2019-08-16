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
	ddb.GlobalSecondaryIndex = &tfgen.GlobalSecondaryIndex{
		Name: "test",
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

	tf.AddResource(ddb)
	tf.AddResource(sqs)
	jsont, _ := json.Marshal(ddb)
	fmt.Println(string(jsont))

	tf.SaveFile()
}
