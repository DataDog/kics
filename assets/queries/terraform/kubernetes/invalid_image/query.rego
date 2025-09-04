package Cx

import data.generic.terraform as tf_lib
import data.generic.common as common_lib

types := {"init_container", "container"}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_array(containers) == true
	containerTypes := containers[_]

	not common_lib.valid_key(containerTypes, "image")

	result := {
		"documentId": input.document[i].id,



		"issueType": "MissingAttribute",


	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_array(containers) == true
	image := containers[y].image
	check_content(image)

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_object(containers) == true
	not common_lib.valid_key(containers, "image")

	result := {
		"documentId": input.document[i].id,



		"issueType": "MissingAttribute",


	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_object(containers) == true
	image := containers.image
	check_content(image)

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}

check_content(image) {
	options := {"", "latest"}

	image == options[x]
}
