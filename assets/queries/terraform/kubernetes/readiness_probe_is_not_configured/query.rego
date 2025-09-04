package Cx

import data.generic.terraform as tf_lib
import data.generic.common as common_lib

types := {"init_container", "container"}

CxPolicy[result] {




	specInfo := tf_lib.getSpecInfo(resource[name])
	container := specInfo.spec[types[x]]

	is_object(container) == true

	not common_lib.valid_key(container, "readiness_probe")

	result := {
		"documentId": input.document[i].id,



		"issueType": "MissingAttribute",


	}
}

CxPolicy[result] {




	specInfo := tf_lib.getSpecInfo(resource[name])
	container := specInfo.spec[types[x]]

	is_array(container) == true
	containersType := container[_]

	not common_lib.valid_key(containersType, "readiness_probe")

	result := {
		"documentId": input.document[i].id,



		"issueType": "MissingAttribute",


	}
}

resource_equal(type) {
	resources := {"kubernetes_cron_job", "kubernetes_job"}

	type == resources[_]
}
