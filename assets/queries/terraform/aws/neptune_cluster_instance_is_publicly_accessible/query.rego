package Cx

import data.generic.terraform as tf_lib
import data.generic.common as common_lib

CxPolicy[result] {
	neptuneClusterInstance := input.document[i].resource.aws_neptune_cluster_instance[name]

	neptuneClusterInstance.publicly_accessible == true

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("aws_neptune_cluster_instance[%s].publicly_accessible", [name]),
		"searchLine": common_lib.build_search_line(["resource", "aws_neptune_cluster_instance", name, "publicly_accessible"], []),
		"issueType": "IncorrectValue",
		"keyExpectedValue": sprintf("aws_neptune_cluster_instance[%s].publicly_accessible should be set to false", [name]),
		"keyActualValue": sprintf("aws_neptune_cluster_instance[%s].publicly_accessible is set to true", [name]),
		"remediation": json.marshal({
			"before": "true",
			"after": "false"
		}),
		"remediationType": "replacement",
	}
}

#######################################################################################################

CxPolicy[result] {
	module := input.document[i].module[name]
	keyToCheck := common_lib.get_module_equivalent_key("aws", module.source, "aws_neptune_cluster_instance", "publicly_accessible")

	module[keyToCheck] == true

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("module[%s].%s", [name, keyToCheck]),
		"searchLine": common_lib.build_search_line(["module", name, keyToCheck], []),
		"issueType": "IncorrectValue",
		"keyExpectedValue": sprintf("module[%s].%s should be set to false", [name, keyToCheck]),
		"keyActualValue": sprintf("module[%s].%s is set to true", [name, keyToCheck]),
		"remediation": json.marshal({
			"before": "true",
			"after": "false"
		}),
		"remediationType": "replacement"
	}
}
