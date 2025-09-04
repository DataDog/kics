package Cx

import data.generic.terraform as tf_lib
import data.generic.common as common_lib

types := {"init_container", "container"}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_array(containers) == true
	containersRequest := containers[_].resources.requests

	not common_lib.valid_key(containersRequest, "memory")

	result := {
		"documentId": input.document[i].id,



		"issueType": "MissingAttribute",


	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_object(containers) == true

	not common_lib.valid_key(containers.resources.requests, "memory")

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

	not common_lib.valid_key(containersType, "resources")

	result := {
		"documentId": input.document[i].id,



		"issueType": "MissingAttribute",


	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_object(containers) == true
	not common_lib.valid_key(containers, "resources")

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

	not common_lib.valid_key(containersType.resources, "requests")

	result := {
		"documentId": input.document[i].id,



		"issueType": "MissingAttribute",


	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_object(containers) == true

	not common_lib.valid_key(containers.resources, "requests")

	result := {
		"documentId": input.document[i].id,



		"issueType": "MissingAttribute",


	}
}
