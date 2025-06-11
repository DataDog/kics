package Cx

import data.generic.terraform as tf_lib
import data.generic.common as common_lib

domains := {
  "aws_opensearch_domain",
  "aws_elasticsearch_domain"
}

CxPolicy[result] {
    resource_type := domains[_]
    resource := input.document[i].resource[resource_type][name]

    not common_lib.valid_key(resource, "advanced_security_options")

    result := {
        "documentId": input.document[i].id,
        "resourceType": resource_type,
        "resourceName": tf_lib.get_resource_name(resource, name),
        "searchKey": sprintf("%s[{{%s}}].advanced_security_options", [resource_type, name]),
        "searchLine": common_lib.build_search_line(["resource", resource_type, name, "advanced_security_options"], []),
        "issueType": "MissingAttribute",
        "keyExpectedValue": "advanced_security_options block should be present with enabled = true and internal_user_database_enabled = true",
        "keyActualValue": "advanced_security_options block is missing"
    }
}

CxPolicy[result] {
    resource_type := domains[_]
    resource := input.document[i].resource[resource_type][name]

    common_lib.valid_key(resource.advanced_security_options, "enabled")
    not resource.advanced_security_options.enabled

    result := {
        "documentId": input.document[i].id,
        "resourceType": resource_type,
        "resourceName": tf_lib.get_resource_name(resource, name),
        "searchKey": sprintf("%s[{{%s}}].advanced_security_options", [resource_type, name]),
        "searchLine": common_lib.build_search_line(["resource", resource_type, name, "advanced_security_options"], []),
        "issueType": "IncorrectValue",
        "keyExpectedValue": "advanced_security_options.enabled should be true",
        "keyActualValue": "advanced_security_options.enabled is false"
    }
}

CxPolicy[result] {
    resource_type := domains[_]
    resource := input.document[i].resource[resource_type][name]

    common_lib.valid_key(resource.advanced_security_options, "internal_user_database_enabled")
    not resource.advanced_security_options.internal_user_database_enabled

    result := {
        "documentId": input.document[i].id,
        "resourceType": resource_type,
        "resourceName": tf_lib.get_resource_name(resource, name),
        "searchKey": sprintf("%s[{{%s}}].advanced_security_options.internal_user_database_enabled", [resource_type, name]),
        "searchLine": common_lib.build_search_line(["resource", resource_type, name, "advanced_security_options", "internal_user_database_enabled"], []),
        "issueType": "IncorrectValue",
        "keyExpectedValue": "advanced_security_options.internal_user_database_enabled should be true",
        "keyActualValue": "advanced_security_options.internal_user_database_enabled is false"
    }
}