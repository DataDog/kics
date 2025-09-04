package Cx

import data.generic.terraform as tf_lib

CxPolicy[result] {
	resource := input.document[i].resource.google_container_cluster[primary]
	resource.monitoring_service == "none"

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("google_container_cluster[%s].monitoring_service", [primary]),
		"issueType": "IncorrectValue",
		"keyExpectedValue": "Attribute 'monitoring_service' should be undefined or 'monitoring.googleapis.com/kubernetes'",
		"keyActualValue": "Attribute 'monitoring_service' is 'none'",
	}
}

# Legacy Stackdriver was decomissioned Mar 31, 2021
CxPolicy[result] {
	resource := input.document[i].resource.google_container_cluster[primary]
	resource.monitoring_service == "monitoring.googleapis.com"

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("google_container_cluster[%s].monitoring_service", [primary]),
		"issueType": "IncorrectValue",
		"keyExpectedValue": "Attribute 'monitoring_service' should be undefined or 'monitoring.googleapis.com/kubernetes'",
		"keyActualValue": "Attribute 'monitoring_service' is 'monitoring.googleapis.com'",
	}
}
