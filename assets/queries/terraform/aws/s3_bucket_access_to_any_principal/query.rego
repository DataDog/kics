package Cx

import data.generic.common as common_lib
import data.generic.terraform as tf_lib

pl := {"aws_s3_bucket_policy", "aws_s3_bucket"}

CxPolicy[result] {



	access_to_any_principal(resource.policy)

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",



	}
}

CxPolicy[result] {
	module := input.document[i].module[name]



	access_to_any_principal(module[keyToCheck])

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("module[%s].policy", [name]),
		"issueType": "IncorrectValue",
		"keyExpectedValue": "'policy.Principal' should not equal to, nor contain '*'",
		"keyActualValue": "'policy.Principal' is equal to or contains '*'",
		"searchLine": common_lib.build_search_line(["module", name, "policy"], []),
	}
}

access_to_any_principal(policyValue) {
  policy := common_lib.json_unmarshal(policyValue)

  st := common_lib.get_statement(policy)
  statement := st[_]

  common_lib.is_allow_effect(statement)
  tf_lib.anyPrincipal(statement)

  keys := [k | k := key; _ := policy[k]]
}
