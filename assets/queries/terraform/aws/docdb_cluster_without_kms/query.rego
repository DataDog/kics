package Cx

import data.generic.common as common_lib
import data.generic.terraform as tf_lib

CxPolicy[result] {
	resource := input.document[i].resource.aws_docdb_cluster[name]
	not common_lib.valid_key(resource, "kms_key_id")

	result := {
		"documentId": input.document[i].id,
		"resourceType": "aws_docdb_cluster",
		"resourceName": tf_lib.get_resource_name(resource, name),
		"searchKey": sprintf("aws_docdb_cluster[{{%s}}]", [name]),
		"issueType": "MissingAttribute",
		"keyExpectedValue": "aws_docdb_cluster.kms_key_id should be defined and not null",
		"keyActualValue": "aws_docdb_cluster.kms_key_id is undefined or null",
	}
}

CxPolicy[result] {
	module := input.document[i].module[name]
	keyToCheck := common_lib.get_module_equivalent_key("aws", module.source, "aws_docdb_cluster", "kms_key_id")

	not common_lib.valid_key(module, keyToCheck)

	result := {
		"documentId": input.document[i].id,
		"resourceType": "module",
		"resourceName": sprintf("%s", [name]),
		"searchKey": sprintf("module[%s]", [name]),
		"issueType": "MissingAttribute",
		"keyExpectedValue": sprintf("%s should be defined and not null", [keyToCheck]),
		"keyActualValue": sprintf("%s is undefined or null", [keyToCheck]),
	}
}
