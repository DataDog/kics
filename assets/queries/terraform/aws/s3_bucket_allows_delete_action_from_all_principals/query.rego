package Cx

import data.generic.common as common_lib
import data.generic.terraform as tf_lib

pl := {"aws_s3_bucket_policy", "aws_s3_bucket"}

CxPolicy[result] {



	tf_lib.allows_action_from_all_principals(resource.policy, "delete")

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",



	}
}

CxPolicy[result] {
	module := input.document[i].module[name]



	tf_lib.allows_action_from_all_principals(module[keyToCheck], "delete")

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("module[%s].policy", [name]),
		"issueType": "IncorrectValue",
		"keyExpectedValue": "'policy.Statement.Action' should not be a 'Delete' action",
		"keyActualValue": "'policy.Statement.Action' is a 'Delete' action",
		"searchLine": common_lib.build_search_line(["module", name, "policy"], []),
	}
}
