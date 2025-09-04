package Cx

import data.generic.common as common_lib
import data.generic.terraform as tf_lib

CxPolicy[result] {
	app := input.document[i].resource.azurerm_app_service[name]

	is_number(app.site_config.min_tls_version)
    versions := {1.2, "1.2", 1.3, "1.3"}
	not versions[app.site_config.min_tls_version]

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("azurerm_app_service[%s].site_config.min_tls_version", [name]),
		"issueType": "IncorrectValue",
		"keyExpectedValue": sprintf("'azurerm_app_service[%s].site_config.min_tls_version' should be set to '1.2'", [name]),
		"keyActualValue": sprintf("'azurerm_app_service[%s].site_config.min_tls_version' is not set to '1.2'", [name]),
		"searchLine": common_lib.build_search_line(["resource", "azurerm_app_service", name, "site_config", "min_tls_version"], []),
		"remediation": json.marshal({
			"before": sprintf("%.1f", [app.site_config.min_tls_version]),
			"after": "1.2"
		}),
		"remediationType": "replacement",
	}
}
