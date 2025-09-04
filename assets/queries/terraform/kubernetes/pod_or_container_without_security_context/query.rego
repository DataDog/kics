package Cx

import data.generic.terraform as tf_lib
import data.generic.common as common_lib

CxPolicy[result] {
	resource := input.document[i].resource.kubernetes_pod[name]

	spec := resource.spec

	not common_lib.valid_key(spec, "security_context")

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("kubernetes_pod[%s].spec", [name]),
		"issueType": "MissingAttribute",
		"keyExpectedValue": sprintf("kubernetes_pod[%s].spec.security_context should be set", [name]),
		"keyActualValue": sprintf("kubernetes_pod[%s].spec.security_context is undefined", [name]),
	}
}

types := {"init_container", "container"}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_object(containers) == true
	not common_lib.valid_key(containers, "security_context")

	result := {
		"documentId": input.document[i].id,



		"issueType": "MissingAttribute",


	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_array(containers) == true
	containersType := containers[_]

	not common_lib.valid_key(containersType, "security_context")

	result := {
		"documentId": input.document[i].id,



		"issueType": "MissingAttribute",


	}
}
