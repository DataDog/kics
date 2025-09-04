package Cx

import data.generic.terraform as tf_lib

CxPolicy[result] {
	resource := input.document[i].resource.google_compute_instance[name]
	scopes := resource.service_account.scopes

	some j
	scopes[j] == "cloud-platform"

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("google_compute_instance[%s].service_account.scopes", [name]),
		"issueType": "IncorrectValue",
		"keyExpectedValue": "'service_account.scopes' should not contain 'cloud-platform'",
		"keyActualValue": "'service_account.scopes' contains 'cloud-platform'",
	}
}
