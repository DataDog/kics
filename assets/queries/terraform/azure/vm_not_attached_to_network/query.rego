package Cx

import data.generic.terraform as tf_lib

CxPolicy[result] {
	vm := input.document[i].resource.azurerm_virtual_machine[name]

	count(vm.network_interface_ids) == 0

	result := {
		"documentId": input.document[i].id,


		"searchKey": sprintf("azurerm_virtual_machine[%s].network_interface_ids", [name]),
		"issueType": "IncorrectValue",
		"keyExpectedValue": sprintf("'azurerm_virtual_machine[%s].network_interface_ids' list should not be empty", [name]),
		"keyActualValue": sprintf("'azurerm_virtual_machine[%s].network_interface_ids' list is empty", [name]),
	}
}
