package aws

import (
	"github.com/snyk/driftctl/pkg/resource"
)

const AwsIamRoleResourceType = "aws_iam_role"

func initAwsIAMRoleMetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.UpdateSchema(AwsIamRoleResourceType, map[string]func(attributeSchema *resource.AttributeSchema){
		"assume_role_policy": func(attributeSchema *resource.AttributeSchema) {
			attributeSchema.JsonString = true
		},
	})
	resourceSchemaRepository.SetNormalizeFunc(AwsIamRoleResourceType, func(res *resource.Resource) {
		val := res.Attrs
		val.SafeDelete([]string{"force_detach_policies"})
	})
	resourceSchemaRepository.SetFlags(AwsIamRoleResourceType, resource.FlagDeepMode)
}
