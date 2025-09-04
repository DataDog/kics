package Cx

import data.generic.terraform as tf_lib
import data.generic.common as common_lib

CxPolicy[result] {

	router := input.document[i].resource.nifcloud_router[name]
	not common_lib.valid_key(router, "security_group")

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("nifcloud_router[%s]", [name]),
		"issueType": "MissingAttribute",
		"keyExpectedValue": sprintf("'nifcloud_router[%s]' should include a security_group for security purposes", [name]),
		"keyActualValue": sprintf("'nifcloud_router[%s]' does not have a security_group", [name]),
	}
}
