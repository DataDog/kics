package Cx

import data.generic.common as common_lib
import data.generic.terraform as tf_lib

CxPolicy[result] {
	resource := input.document[i].resource
	aws_dms_replication_instance := resource.aws_dms_replication_instance[name]
	aws_dms_replication_instance.publicly_accessible == true

	result := {
		"documentId": input.document[i].id,
		"resourceType": "aws_dms_replication_instance",
		"resourceName": tf_lib.get_resource_name(resource, name),
		"searchKey": sprintf("aws_dms_replication_instance[%s].publicly_accessible", [name]),
		"issueType": "IncorrectValue",
		"keyExpectedValue": sprintf("aws_dms_replication_instance[%s].publicly_accessible should be set to false", [name]),
		"keyActualValue": sprintf("aws_dms_replication_instance[%s].publicly_accessible is set to true", [name]),
	}
}

CxPolicy[result] {
	module := input.document[i].module[name]
	keyToCheck := common_lib.get_module_equivalent_key("aws", module.source, "aws_dms_replication_instance", "publicly_accessible")

	module[keyToCheck] == true

	result := {
		"documentId": input.document[i].id,
		"resourceType": "module",
		"resourceName": sprintf("%s", [name]),
		"searchKey": sprintf("module[%s].publicly_accessible", [name]),
		"issueType": "IncorrectValue",
		"keyExpectedValue": "'publicly_accessible' should be set to false",
		"keyActualValue": "'publicly_accessible' is set to true",
	}
}

