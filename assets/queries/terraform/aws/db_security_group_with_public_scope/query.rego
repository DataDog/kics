package Cx

import data.generic.terraform as tf_lib

CxPolicy[result] {
	resource := input.document[i].resource.aws_db_security_group[name].ingress
	resource.cidr == "0.0.0.0/0"

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("aws_db_security_group[%s].ingress.cidr", [name]),
		"issueType": "IncorrectValue",
		"keyExpectedValue": "'aws_db_security_group.ingress.cidr' != 0.0.0.0/0",
		"keyActualValue": "'aws_db_security_group.ingress.cidr'= 0.0.0.0/0",
	}
}

#######################################################################################################

CxPolicy[result] {
	module := input.document[i].module[name]
	keyToCheck := data.lib.get_module_equivalent_key("aws", module.source, "aws_db_security_group", "ingress")
	module[keyToCheck].cidr == "0.0.0.0/0"

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("module[%s].%s.cidr", [name, keyToCheck]),
		"issueType": "IncorrectValue",
		"keyExpectedValue": sprintf("'module[%s].%s.cidr' != 0.0.0.0/0", [name, keyToCheck]),
		"keyActualValue": sprintf("'module[%s].%s.cidr'= 0.0.0.0/0", [name, keyToCheck]),
	}
}
