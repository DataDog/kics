package Cx

import data.generic.common as common_lib
import data.generic.terraform as tf_lib

resources := {"aws_s3_bucket_policy", "aws_s3_bucket"}

CxPolicy[result] {



	policy_unmarshaled := common_lib.json_unmarshal(resource.policy)
	not deny_http_requests(policy_unmarshaled)

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",



	}
}

CxPolicy[result] {
	module := input.document[i].module[name]



	policy := module[keyToCheck]

	policy_unmarshaled := common_lib.json_unmarshal(policy)
	not deny_http_requests(policy_unmarshaled)

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("module[%s].policy", [name]),
		"issueType": "IncorrectValue",
		"keyExpectedValue": "'policy' should not accept HTTP Requests",
		"keyActualValue": "'policy' accepts HTTP Requests",
		"searchLine": common_lib.build_search_line(["module", name, "policy"], []),
	}
}

any_s3_action(action) {
	any([action == "*", startswith(action, "s3:")])
}
check_action(st) {
	is_string(st.Action)
	any_s3_action(st.Action)
} else {
	any_s3_action(st.Action[a])
} else {
	is_string(st.Actions)
	any_s3_action(st.Actions)
} else {
	any_s3_action(st.Actions[a])
}

is_equal(secure, target)
{
    secure == target
}else {
    secure[_]==target
}

deny_http_requests(policyValue) {
    st := common_lib.get_statement(policyValue)
    statement := st[_]
    check_action(statement)
    statement.Effect == "Deny"
    is_equal(statement.Condition.Bool["aws:SecureTransport"], "false")
} else {
    st := common_lib.get_statement(policyValue)
    statement := st[_]
    check_action(statement)
    statement.Effect == "Allow"
    is_equal(statement.Condition.Bool["aws:SecureTransport"], "true")
}
