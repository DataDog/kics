package Cx

import data.generic.terraform as tf_lib

CxPolicy[result] {


	checkMetadata(resource[name].metadata)

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}

types := {"init_container", "container"}

CxPolicy[result] {


	spec := resource[name].spec
	containers := spec[types[x]]

	is_array(containers) == true
	some y
	contains(object.get(containers[y], "image", "undefined"), "tiller")

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}

CxPolicy[result] {


	spec := resource[name].spec
	containers := spec[types[x]]

	is_object(containers) == true
	contains(object.get(containers, "image", "undefined"), "tiller")

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}

CxPolicy[result] {


	spec := resource[name].spec

	checkMetadata(spec.template.metadata)

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}

CxPolicy[result] {


	spec := resource[name].spec.template.spec
	containers := spec[types[x]]

	is_object(containers) == true

	contains(object.get(containers, "image", "undefined"), "tiller")

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}

CxPolicy[result] {


	spec := resource[name].spec.template.spec
	containers := spec[types[x]]

	is_array(containers) == true

	contains(object.get(containers[y], "image", "undefined"), "tiller")

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}

checkMetadata(metadata) {
	contains(metadata.name, "tiller")
}

checkMetadata(metadata) {
	object.get(metadata.labels, "app", "undefined") == "helm"
}

checkMetadata(metadata) {
	contains(object.get(metadata.labels, "name", "undefined"), "tiller")
}
