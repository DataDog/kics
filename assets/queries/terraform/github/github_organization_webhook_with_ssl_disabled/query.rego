package Cx

import data.generic.terraform as tf_lib

CxPolicy[result] {
	webhook := input.document[i].resource.github_organization_webhook[name]
	webhook.configuration.insecure_ssl == true

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("github_organization_webhook[%s].configuration.insecure_ssl", [name]),
		"issueType": "IncorrectValue",
		"keyExpectedValue": sprintf("github_organization_webhook[%s].configuration.insecure_ssl should be set to false", [name]),
		"keyActualValue": sprintf("github_organization_webhook[%s].configuration.insecure_ssl is true", [name]),
	}
}
