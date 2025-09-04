package Cx

import data.generic.terraform as tf_lib

CxPolicy[result] {
	resource := input.document[i].resource.databricks_token[name]
	not resource.lifetime_seconds

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("databricks_token[%s]", [name]),
		"issueType": "MissingAttribute",
		"keyExpectedValue": sprintf("'databricks_token[%s]' should not have indefinitely lifetime", [name]),
		"keyActualValue": sprintf("'databricks_token[%s]' have an indefinitely lifetime", [name]),
	}
}
