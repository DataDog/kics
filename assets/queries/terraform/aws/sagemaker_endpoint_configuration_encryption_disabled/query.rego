package Cx

import data.generic.common as common_lib
import data.generic.terraform as tf_lib

CxPolicy[result] {
	sagemakerEndpoint := input.document[i].resource.aws_sagemaker_endpoint_configuration[name]

	not common_lib.valid_key(sagemakerEndpoint, "kms_key_arn")

	result := {
		"documentId": input.document[i].id,
		"resourceType": "aws_sagemaker_endpoint_configuration",
		"resourceName": tf_lib.get_resource_name(sagemakerEndpoint, name),
		"searchKey": sprintf("aws_sagemaker_endpoint_configuration[%s]", [name]),
		"issueType": "MissingAttribute",
		"keyExpectedValue": sprintf("aws_sagemaker_endpoint_configuration[%s] should be defined and not null", [name]),
		"keyActualValue": sprintf("aws_sagemaker_endpoint_configuration[%s] is undefined or null", [name]),
	}
}

#######################################################################################################

CxPolicy[result] {
	module := input.document[i].module[name]
	keyToCheck := common_lib.get_module_equivalent_key("aws", module.source, "aws_sagemaker_endpoint_configuration", "endpoint_config")

	not common_lib.valid_key(module[keyToCheck], "kms_key_arn")

	result := {
		"documentId": input.document[i].id,
		"resourceType": "module",
		"resourceName": sprintf("%s", [name]),
		"searchKey": sprintf("module[%s]", [name]),
		"issueType": "MissingAttribute",
		"keyExpectedValue": sprintf("module[%s].%s.kms_key_arn should be defined and not null", [name, keyToCheck]),
		"keyActualValue": sprintf("module[%s].%s.kms_key_arn is undefined or null", [name, keyToCheck]),
	}
}
