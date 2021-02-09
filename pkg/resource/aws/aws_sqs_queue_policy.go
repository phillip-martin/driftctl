// GENERATED, DO NOT EDIT THIS FILE
package aws

const AwsSqsQueuePolicyResourceType = "aws_sqs_queue_policy"

type AwsSqsQueuePolicy struct {
	Id       string  `cty:"id" computed:"true"`
	Policy   *string `cty:"policy" jsonstring:"true"`
	QueueUrl *string `cty:"queue_url"`
}

func (r *AwsSqsQueuePolicy) TerraformId() string {
	return r.Id
}

func (r *AwsSqsQueuePolicy) TerraformType() string {
	return AwsSqsQueuePolicyResourceType
}
