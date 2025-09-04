package Cx

import data.generic.common as common_lib
import data.generic.terraform as tf_lib

CxPolicy[result] {
	document := input.document[i]
	resource = document.resource.aws_mq_broker[name]

	not common_lib.valid_key(resource, "encryption_options")

	result := {
		"documentId": document.id,


		"searchKey": sprintf("resource.aws_mq_broker[%s]", [name]),
		"issueType": "MissingAttribute",
		"keyExpectedValue": sprintf("resource.aws_mq_broker[%s].encryption_options should be defined", [name]),
		"keyActualValue": sprintf("resource.aws_mq_broker[%s].encryption_options is not defined", [name]),
	}
}

#######################################################################################################

CxPolicy[result] {
	module := input.document[i].module[name]
	keyToCheck := common_lib.get_module_equivalent_key("aws", module.source, "aws_mq_broker", "encryption_options")
	not common_lib.valid_key(module, keyToCheck)

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("module[%s]", [name]),
		"issueType": "MissingAttribute",
		"keyExpectedValue": sprintf("module[%s].%s should be defined", [name, keyToCheck]),
		"keyActualValue": sprintf("module[%s].%s is not defined", [name, keyToCheck]),
	}
}
