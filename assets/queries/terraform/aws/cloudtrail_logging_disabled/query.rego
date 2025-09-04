package Cx

import data.generic.terraform as tf_lib
import data.generic.common as common_lib

CxPolicy[result] {
	resource := input.document[i].resource.aws_cloudtrail[name]
	resource.enable_logging == false

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("aws_cloudtrail.%s.enable_logging", [name]),
		"searchLine": common_lib.build_search_line(["resource", "aws_cloudtrail", name, "enable_logging"], []),
		"issueType": "IncorrectValue",
		"keyExpectedValue": sprintf("aws_cloudtrail.%s.enable_logging should be true", [name]),
		"keyActualValue": sprintf("aws_cloudtrail.%s.enable_logging is false", [name]),
		"remediation": json.marshal({
			"before": "false",
			"after": "true"
		}),
		"remediationType": "replacement",
	}
}

#######################################################################################################

CxPolicy[result] {
	module := input.document[i].module[name]
	keyToCheck := common_lib.get_module_equivalent_key("aws", module.source, "aws_cloudtrail", "enable_logging")
	module[keyToCheck] == false

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("module[%s].%s", [name, keyToCheck]),
		"searchLine": common_lib.build_search_line(["module", name, keyToCheck], []),
		"issueType": "IncorrectValue",
		"keyExpectedValue": sprintf("module[%s].%s should be true", [name, keyToCheck]),
		"keyActualValue": sprintf("module[%s].%s is false", [name, keyToCheck]),
		"remediation": json.marshal({
			"before": "false",
			"after": "true"
		}),
		"remediationType": "replacement",
	}
}
