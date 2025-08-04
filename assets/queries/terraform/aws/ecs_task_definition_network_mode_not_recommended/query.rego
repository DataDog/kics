package Cx

import data.generic.terraform as tf_lib
import data.generic.common as common_lib

CxPolicy[result] {
	resource := input.document[i].resource.aws_ecs_task_definition[name]
	lower(resource.network_mode) != "awsvpc"

	result := {
		"documentId": input.document[i].id,
		"resourceType": "aws_ecs_task_definition",
		"resourceName": tf_lib.get_resource_name(resource, name),
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

CxPolicy[result] {
	module := input.document[i].module[name]
	keyToCheck := common_lib.get_module_equivalent_key("aws", module.source, "aws_ecs_task_definition", "network_mode")

	lower(module[keyToCheck]) != "awsvpc"

	result := {
		"documentId": input.document[i].id,
		"resourceType": "module",
		"resourceName": sprintf("%s", [name]),
		"searchKey": sprintf("module[%s].network_mode", [name]),
		"searchLine": common_lib.build_search_line(["module", name, "network_mode"], []),
		"issueType": "IncorrectValue",
		"keyExpectedValue": "'network_mode' should equal to 'awsvpc'",
		"keyActualValue": sprintf("'network_mode' is equal to '%s'", [module.network_mode]),
		"remediation": json.marshal({
			"before": sprintf("%s",[module.network_mode]),
			"after": "awsvpc"
		}),
		"remediationType": "replacement",
	}
}

