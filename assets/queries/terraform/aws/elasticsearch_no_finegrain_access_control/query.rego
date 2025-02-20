package Cx

import data.generic.terraform as tf_lib
import data.generic.common as common_lib

CxPolicy[result] {
    resource := input.document[i].resource.aws_opensearch_domain[name]

    resource.advanced_security_options.internal_user_database_enabled != true
    resource.advanced_security_options.enabled != true

    result := {
        "documentId": input.document[i].id,
        "resourceType": "aws_opensearch_domain",
        "resourceName": tf_lib.get_resource_name(resource, name),
        "searchKey": sprintf("aws_opensearch_domain[{{%s}}].advanced_security_options", [name]),
        "searchLine": common_lib.build_search_line(["resource", "aws_opensearch_domain", name, "advanced_security_options"], []),
        "issueType": "IncorrectValue",
        "keyExpectedValue": "advanced_security_options.internal_user_database_enabled and advanced_security_options.enabled should be true",
        "keyActualValue": "advanced_security_options.internal_user_database_enabled or advanced_security_options.enabled is false"
    }
}

CxPolicy[result] {
    resource := input.document[i].resource.aws_elasticsearch_domain[name]

    resource.advanced_security_options.internal_user_database_enabled != true
    resource.advanced_security_options.enabled != true

    result := {
        "documentId": input.document[i].id,
        "resourceType": "aws_elasticsearch_domain",
        "resourceName": tf_lib.get_resource_name(resource, name),
        "searchKey": sprintf("aws_elasticsearch_domain[{{%s}}].advanced_security_options", [name]),
        "searchLine": common_lib.build_search_line(["resource", "aws_elasticsearch_domain", name, "advanced_security_options"], []),
        "issueType": "IncorrectValue",
        "keyExpectedValue": "advanced_security_options.internal_user_database_enabled and advanced_security_options.enabled should be true",
        "keyActualValue": "advanced_security_options.internal_user_database_enabled or advanced_security_options.enabled is false"
    }
}
