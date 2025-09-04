package Cx

import data.generic.terraform as tf_lib
import data.generic.common as common_lib

CxPolicy[result] {
	awsGuardDuty := input.document[i].resource.aws_guardduty_detector[name]
	detector := awsGuardDuty.enable

	detector == false

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("aws_guardduty_detector[%s].enable", [name]),
		"searchLine": common_lib.build_search_line(["resource", "aws_guardduty_detector", name, "enable"], []),
		"issueType": "IncorrectValue",
		"keyExpectedValue": "GuardDuty Detector should be Enabled",
		"keyActualValue": "GuardDuty Detector is not Enabled",
		"remediation": json.marshal({
			"before": "false",
			"after": "true"
		}),
		"remediationType": "replacement",
	}
}

#######################################################################################################

CxPolicy[result] {
	module := input.document[i].module[name]
	keyToCheck := common_lib.get_module_equivalent_key("aws", module.source, "aws_guardduty_detector", "enable")
	detector := module[keyToCheck]

	detector == false

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("module[%s].%s", [name, keyToCheck]),
		"searchLine": common_lib.build_search_line(["module", name, keyToCheck], []),
		"issueType": "IncorrectValue",
		"keyExpectedValue": "GuardDuty Detector should be Enabled",
		"keyActualValue": "GuardDuty Detector is not Enabled",
		"remediation": json.marshal({
			"before": "false",
			"after": "true"
		}),
		"remediationType": "replacement",
	}
}
