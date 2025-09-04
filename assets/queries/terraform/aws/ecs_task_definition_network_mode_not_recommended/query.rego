package Cx

import data.generic.terraform as tf_lib
import data.generic.common as common_lib

CxPolicy[result] {
	resource := input.document[i].resource.aws_ecs_task_definition[name]
	lower(resource.network_mode) != "awsvpc"

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("aws_ecs_task_definition[%s].network_mode", [name]),
		"searchLine": common_lib.build_search_line(["resource", "aws_ecs_task_definition", name, "network_mode"], []),
		"issueType": "IncorrectValue",
		"keyExpectedValue": "'network_mode' should equal to 'awsvpc'",
		"keyActualValue": sprintf("'network_mode' is equal to '%s'", [resource.network_mode]),
		"remediation": json.marshal({
			"before": sprintf("%s",[resource.network_mode]),
			"after": "awsvpc"
		}),
		"remediationType": "replacement",
	}
}

#######################################################################################################

CxPolicy[result] {
	module := input.document[i].module[name]
	keyToCheck := common_lib.get_module_equivalent_key("aws", module.source, "aws_ecs_task_definition", "network_mode")

	lower(module[keyToCheck]) != "awsvpc"

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("module[%s].%s", [name, keyToCheck]),
		"searchLine": common_lib.build_search_line(["module", name, keyToCheck], []),
		"issueType": "IncorrectValue",
		"keyExpectedValue": sprintf("'module[%s].%s' should equal to 'awsvpc'", [name, keyToCheck]),
		"keyActualValue": sprintf("'module[%s].%s' is equal to '%s'", [name, keyToCheck, module[keyToCheck]]),
		"remediation": json.marshal({
			"before": sprintf("%s",[module[keyToCheck]]),
			"after": "awsvpc"
		}),
		"remediationType": "replacement",
	}
}
