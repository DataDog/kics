package Cx

import data.generic.terraform as tf_lib
import data.generic.common as common_lib

types := {"init_container", "container"}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])

	containers := specInfo.spec[types[x]]

	is_array(containers) == true

	has_secret_key_ref(containers[y])

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])

	containers := specInfo.spec[types[x]]

	is_object(containers) == true

	has_secret_key_ref(containers)

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])

	containers := specInfo.spec[types[x]]

	is_array(containers) == true

	has_secret_key_ref(containers[y])

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])

	containers := specInfo.spec[types[x]]

	is_object(containers) == true

	has_secret_key_ref(containers)

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}

has_secret_key_ref(container) {
	is_array(container.env) == true

	common_lib.valid_key(container.env[x].value_from, "secret_key_ref")
}

has_secret_key_ref(container) {
	is_object(container.env) == true

	common_lib.valid_key(container.env.value_from, "secret_key_ref")
}

has_secret_key_ref(container) {
	is_array(container.env) == true

	common_lib.valid_key(container.env_from, "secret_ref")
}

has_secret_key_ref(container) {
	is_object(container.env) == true

	common_lib.valid_key(container.env_from, "secret_ref")
}
