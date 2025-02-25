package Cx

import data.generic.common as common_lib
import data.generic.terraform as tf_lib

CxPolicy[result] {
    resource := input.document[i].resource[_][name]
    not common_lib.valid_key(resource, "labels")
    result := {
        "documentId": input.document[i].id,
        "resourceType": "any",
        "resourceName": tf_lib.get_resource_name(resource, name),
        "searchKey": sprintf("resource[%s].labels", [name]),
        "searchLine": common_lib.build_search_line(["resource", name, "labels"], []),
        "issueType": "MissingValue",
        "keyExpectedValue": "Resource must have a 'team' label",
        "keyActualValue": "'labels' block is missing",
        "remediation": json.marshal({
            "before": "No 'labels' block",
            "after": "Add labels = { team = \"YourTeamName\" }"
        }),
        "remediationType": "addition"
    }
}

CxPolicy[result] {
    resource := input.document[i].resource[_][name]
    common_lib.valid_key(resource, "labels")
    labels := resource.labels
    not common_lib.valid_key(labels, "team")
    result := {
        "documentId": input.document[i].id,
        "resourceType": "any",
        "resourceName": tf_lib.get_resource_name(resource, name),
        "searchKey": sprintf("resource[%s].labels.team", [name]),
        "searchLine": common_lib.build_search_line(["resource", name, "labels", "team"], []),
        "issueType": "MissingValue",
        "keyExpectedValue": "Resource must have a 'team' label",
        "keyActualValue": "'team' label is missing",
        "remediation": json.marshal({
            "before": "labels = { ... }",
            "after": "labels = { ..., team = \"YourTeamName\" }"
        }),
        "remediationType": "addition"
    }
}
