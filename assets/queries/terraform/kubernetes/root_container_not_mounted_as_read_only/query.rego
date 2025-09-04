package Cx

import data.generic.terraform as tf_lib
import data.generic.common as common_lib

types := {"init_container", "container"}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_array(containers) == true
	containersType := containers[_]

	not common_lib.valid_key(containersType, "security_context")

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",



	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_array(containers) == true
    common_lib.valid_key(containers[j], "security_context")
    not common_lib.valid_key(containers[j].security_context, "read_only_root_filesystem")

	result := {
		"documentId": input.document[i].id,


		"issueType": "IncorrectValue",



		"remediation": "read_only_root_filesystem = true",
		"remediationType": "addition",
	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_array(containers) == true

	common_lib.valid_key(containers[y], "security_context")
	containers[y].security_context.read_only_root_filesystem == false

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",



		"remediation": json.marshal({
			"before": "false",
			"after": "true"
		}),
		"remediationType": "replacement",
	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_object(containers) == true
	not common_lib.valid_key(containers, "security_context")

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",



	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_object(containers) == true
	not common_lib.valid_key(containers.security_context, "read_only_root_filesystem")

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",



		"remediation": "read_only_root_filesystem = true",
		"remediationType": "addition",
	}
}

CxPolicy[result] {


	specInfo := tf_lib.getSpecInfo(resource[name])
	containers := specInfo.spec[types[x]]

	is_object(containers) == true
	containers.security_context.read_only_root_filesystem == false

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",



		"remediation": json.marshal({
			"before": "false",
			"after": "true"
		}),
		"remediationType": "replacement",
	}
}
