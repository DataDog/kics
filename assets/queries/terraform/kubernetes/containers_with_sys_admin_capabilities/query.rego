package Cx

import data.generic.terraform as tf_lib

types := {"init_container", "container"}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_array(containers) == true
	containers[y].security_context.capabilities.add[_] = "SYS_ADMIN"

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_object(containers) == true
	containers.security_context.capabilities.add[_] = "SYS_ADMIN"

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}
