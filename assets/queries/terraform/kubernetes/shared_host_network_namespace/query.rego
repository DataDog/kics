package Cx

import data.generic.terraform as tf_lib
import data.generic.common as common_lib

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])

	specInfo.spec.host_network == true

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",



		"remediation": json.marshal({
			"before": "true",
			"after": "false"
		}),
		"remediationType": "replacement",
	}
}
