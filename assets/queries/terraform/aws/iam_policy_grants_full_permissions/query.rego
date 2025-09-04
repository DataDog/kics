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
		"keyExpectedValue": "'policy.Statement.Resource' and 'policy.Statement.Action' should not equal '*'",
		"keyActualValue": "'policy.Statement.Resource' and 'policy.Statement.Action' are equal to '*'",

	}
}
