package Cx

import data.generic.terraform as tf_lib

CxPolicy[result] {


	labels := resource[name].metadata.labels

	regex.match("^(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])?$", labels[key]) == false

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}
