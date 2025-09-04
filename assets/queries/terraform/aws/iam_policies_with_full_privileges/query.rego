package Cx

import data.generic.common as common_lib
import data.generic.terraform as tf_lib

CxPolicy[result] {



	policy := common_lib.json_unmarshal(resource.policy)

	st := common_lib.get_statement(policy)
	statement := st[_]

	common_lib.is_allow_effect(statement)
	common_lib.equalsOrInArray(statement.Resource, "*")
	common_lib.equalsOrInArray(statement.Action, "*")

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",
		"keyExpectedValue": "'policy.Statement.Action' shouldn't contain '*'",
		"keyActualValue": "'policy.Statement.Action' contains '*'",

	}
}

CxPolicy[result] {
	resource := input.document[i].data.aws_iam_policy_document[name]

    policy := {"Statement": resource.statement}

	st := common_lib.get_statement(policy)
	statement := st[_]

	common_lib.is_allow_effect(statement)
	common_lib.equalsOrInArray(statement.resources, "*")
	common_lib.equalsOrInArray(statement.actions, "*")

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("aws_iam_policy_document[%s].policy", [name]),
		"issueType": "IncorrectValue",
		"keyExpectedValue": "'policy.Statement.Action' shouldn't contain '*'",
		"keyActualValue": "'policy.Statement.Action' contains '*'",
		"searchLine": common_lib.build_search_line(["resource", "aws_iam_policy_document", name, "policy"], []),
	}
}
