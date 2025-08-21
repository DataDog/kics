package Cx

import data.generic.terraform as tf_lib
import data.generic.common as common_lib

CxPolicy[result] {
	resource := input.document[i].resource.aws_sns_topic[name]

	not common_lib.valid_key(resource, "kms_master_key_id")

	result := {
		"documentId": input.document[i].id,
		"resourceType": "aws_sns_topic",
		"resourceName": tf_lib.get_resource_name(resource, name),
		"searchKey": sprintf("aws_sns_topic[%s]", [name]),
		"issueType": "MissingAttribute",
		"keyExpectedValue": "kms_master_key_id should be defined and not null",
        "keyActualValue": "kms_master_key_id is undefined or null",
        "searchLine": common_lib.build_search_line(["resource", "aws_sns_topic", name], []),
        "remediation": "kms_master_key_id = \"alias/aws/sns\"",
        "remediationType": "addition",
	}
}

CxPolicy[result] {
	resource := input.document[i].resource.aws_sns_topic[name]

	resource.kms_master_key_id == ""

	result := {
		"documentId": input.document[i].id,
		"resourceType": "aws_sns_topic",
		"resourceName": tf_lib.get_resource_name(resource, name),
		"searchKey": sprintf("aws_sns_topic[%s].kms_master_key_id", [name]),
		"issueType": "IncorrectValue",
		"keyExpectedValue": "kms_master_key_id should be defined and not null",
        "keyActualValue": "kms_master_key_id is empty string",
        "searchLine": common_lib.build_search_line(["resource", "aws_sns_topic", name, "kms_master_key_id"], []),
        "remediation": json.marshal({
            "before": "\"\"",
            "after":  "\"alias/aws/sns\"",
        }),
        "remediationType": "replacement",
	}
}

#######################################################################################################

CxPolicy[result] {
	module := input.document[i].module[name]
	keyToCheck := common_lib.get_module_equivalent_key("aws", module.source, "aws_sns_topic", "kms_master_key_id")

	not common_lib.valid_key(module, keyToCheck)

	result := {
		"documentId": input.document[i].id,
		"resourceType": "module",
		"resourceName": sprintf("%s", [name]),
		"searchKey": sprintf("module[%s]", [name]),
		"issueType": "MissingAttribute",
		"keyExpectedValue": sprintf("module[%s].%s should be defined and not null", [name, keyToCheck]),
		"keyActualValue": sprintf("module[%s].%s is undefined or null", [name, keyToCheck]),
        "searchLine": common_lib.build_search_line(["module", name], []),
        "remediation": sprintf("%s = \"alias/aws/sns\"", [keyToCheck]),
        "remediationType": "addition",
	}
}

CxPolicy[result] {
	module := input.document[i].module[name]
	keyToCheck := common_lib.get_module_equivalent_key("aws", module.source, "aws_sns_topic", "kms_master_key_id")

	module[keyToCheck] == ""

	result := {
		"documentId": input.document[i].id,
		"resourceType": "module",
		"resourceName": sprintf("%s", [name]),
		"searchKey": sprintf("module[%s].%s", [name, keyToCheck]),
		"issueType": "IncorrectValue",
		"keyExpectedValue": sprintf("module[%s].%s should be defined and not null", [name, keyToCheck]),
        "keyActualValue": sprintf("module[%s].%s is empty string", [name, keyToCheck]),
        "searchLine": common_lib.build_search_line(["module", name, keyToCheck], []),
        "remediation": json.marshal({
            "before": "\"\"",
            "after":  "\"alias/aws/sns\"",
        }),
        "remediationType": "replacement",
	}
}
