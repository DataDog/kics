package Cx

import data.generic.terraform as tf_lib

create := "create"



pods := "pods"

CxPolicy[result] {

	resource.rule[ru].verbs[l] == create
	resource.rule[ru].resources[r] == pods

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}

CxPolicy[result] {

	resource.rule[ru].verbs[l] == create
	isWildCardValue(resource.rule[ru].resources[r])

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}

CxPolicy[result] {

	isWildCardValue(resource.rule[ru].verbs[l])
	resource.rule[ru].resources[r] == pods

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}

CxPolicy[result] {

	isWildCardValue(resource.rule[ru].verbs[l])
	isWildCardValue(resource.rule[ru].resources[r])

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}

CxPolicy[result] {

	resource.rule.verbs[l] == create
	resource.rule.resources[r] == pods

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}

CxPolicy[result] {

	resource.rule.verbs[l] == create
	isWildCardValue(resource.rule.resources[r])

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}

CxPolicy[result] {

	isWildCardValue(resource.rule.verbs[l])
	resource.rule.resources[r] == pods

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}

CxPolicy[result] {

	isWildCardValue(resource.rule.verbs[l])
	isWildCardValue(resource.rule.resources[r])

	result := {
		"documentId": input.document[i].id,



		"issueType": "IncorrectValue",


	}
}

isWildCardValue(val) {
	regex.match(".*\\*.*", val)
}
