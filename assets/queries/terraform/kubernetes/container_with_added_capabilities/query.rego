package Cx

import data.generic.terraform as tf_lib
import data.generic.common as common_lib

types := {"init_container", "container"}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_array(containers) == true
	common_lib.valid_key(containers[y].security_context.capabilities, "add")

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_object(containers) == true
	common_lib.valid_key(containers.security_context.capabilities, "add")

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}
