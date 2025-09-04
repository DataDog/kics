package Cx

import data.generic.terraform as tf_lib
import data.generic.common as common_lib

types := {"init_container", "container"}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_array(containers) == true
	requestedContainers := containers[y].resources.requests

	not common_lib.valid_key(requestedContainers, "cpu")

	result := {
		"documentId": input.document[i].id,



		"issueType": "MissingAttribute",


	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_object(containers) == true

	not common_lib.valid_key(containers.resources.requests, "cpu")

	result := {
		"documentId": input.document[i].id,



		"issueType": "MissingAttribute",


	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_array(containers) == true
	containerTypes := containers[_]

	not common_lib.valid_key(containerTypes, "resources")

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
	requestedContainers := containers[y].resources

	not common_lib.valid_key(requestedContainers, "requests")

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
