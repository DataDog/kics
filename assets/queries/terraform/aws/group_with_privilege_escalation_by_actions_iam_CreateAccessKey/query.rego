package Cx

import data.generic.common as common_lib
import data.generic.terraform as tf_lib

CxPolicy[result] {

	# get a AWS IAM group
	group := input.document[i].resource.aws_iam_group[targetGroup]

    common_lib.group_unrecommended_permission_policy_scenarios(targetGroup, "iam:CreateAccessKey")


	result := {
		"documentId": input.document[i].id,
		"resourceType": "aws_iam_group",
        "resourceName": tf_lib.get_resource_name(group, targetGroup),
		"searchKey": sprintf("aws_iam_group[%s]", [targetGroup]),
		"issueType": "IncorrectValue",
        "keyExpectedValue": sprintf("group %s should not be associated with a policy that has Action set to 'iam:CreateAccessKey' and Resource set to '*'", [targetGroup]),
		"keyActualValue": sprintf("group %s is associated with a policy that has Action set to 'iam:CreateAccessKey' and Resource set to '*'", [targetGroup]),
        "searchLine": common_lib.build_search_line(["resource", "aws_iam_group", targetGroup], []),
	}
}

CxPolicy[result] {
	module := input.document[i].module[name]
	keyToCheck := common_lib.get_module_equivalent_key("aws", module.source, "aws_iam_group", "group_policy")
	targetGroup := input.document[i].module[name][keyToCheck][idx]

	groupUnrecommendedScenarios := common_lib.unrecommended_permission_policy(targetGroup, "iam:CreateAccessKey")

	result := {
		"documentId": input.document[i].id,
		"resourceType": "module",
        "resourceName": sprintf("%s", [name]),
		"searchKey": sprintf("module[%s].%s", [name, keyToCheck]),
		"issueType": "IncorrectValue",
        "keyExpectedValue": sprintf("group %s should not be associated with a policy that has Action set to 'iam:CreateAccessKey' and Resource set to '*'", [targetGroup]),
		"keyActualValue": sprintf("group %s is associated with a policy that has Action set to 'iam:CreateAccessKey' and Resource set to '*'", [targetGroup]),
        "searchLine": common_lib.build_search_line(["module", name, keyToCheck, idx], []),
	}
}

