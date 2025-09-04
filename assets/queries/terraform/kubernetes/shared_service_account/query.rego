package Cx

import data.generic.terraform as tf_lib

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])

	service_account_name := specInfo.spec.service_account_name

	service := input.document[j].resource.kubernetes_service_account[name_service]

	service_account_name == name_service

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}
