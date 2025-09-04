package Cx

import data.generic.terraform as tf_lib
import data.generic.common as common_lib

CxPolicy[result] {
	msk_cluster := input.document[i].resource.aws_msk_cluster[name]

	msk_cluster.broker_node_group_info.connectivity_info.public_access.type == "SERVICE_PROVIDED_EIPS"

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("aws_msk_cluster[%s].broker_node_group_info.connectivity_info.public_access.type", [name]),
		"issueType": "IncorrectValue",
		"keyExpectedValue": sprintf("aws_msk_cluster[%s].broker_node_group_info.connectivity_info.public_access.type should be set to 'DISABLED' or undefined", [name]),
		"keyActualValue": sprintf("aws_msk_cluster[%s].broker_node_group_info.connectivity_info.public_access.type is set to 'SERVICE_PROVIDED_EIPS'", [name]),
		"searchLine": common_lib.build_search_line(["resource", "aws_msk_cluster", name, "broker_node_group_info", "connectivity_info", "public_access", "type"], []),
	}
}
