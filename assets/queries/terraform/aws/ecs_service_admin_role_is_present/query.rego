package Cx

import data.generic.terraform as tf_lib

CxPolicy[result] {
	resource := input.document[i].resource.aws_ecs_service[name]
	contains(lower(resource.iam_role), "admin")

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("aws_ecs_service[%s].iam_role", [name]),
		"issueType": "IncorrectValue",
		"keyExpectedValue": sprintf("'aws_ecs_service[%s].iam_role' should not equal to 'admin'", [name]),
		"keyActualValue": sprintf("'aws_ecs_service[%s].iam_role' is equal to 'admin'", [name]),
	}
}
