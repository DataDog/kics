package Cx

import data.generic.terraform as tf_lib
import data.generic.common as common_lib

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])

	specInfo.spec.host_pid == true

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",
		"keyExpectedValue": "Attribute 'host_pid' should be undefined or false",
		"keyActualValue": "Attribute 'host_pid' is true",

		"remediation": json.marshal({
			"before": "true",
			"after": "false"
		}),
		"remediationType": "replacement",
	}
}
