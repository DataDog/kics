package Cx

import data.generic.terraform as tf_lib
import data.generic.common as common_lib

CxPolicy[result] {
    resource := input.document[i].resource.google_cloud_run_service_iam_member[name]
    common_lib.valid_key(resource, "member")
    member := resource.member
    member == "allAuthenticatedUsers"
    result := {
        "documentId": input.document[i].id,


        "searchKey": sprintf("google_cloud_run_service_iam_member[{{%s}}].member", [name]),
        "searchLine": common_lib.build_search_line(["resource", "google_cloud_run_service_iam_member", name, "member"], []),
        "issueType": "IncorrectValue",
        "keyExpectedValue": "IAM member should not be a public principal",
        "keyActualValue": sprintf("Found public principal: %s", [member]),
        "remediation": json.marshal({
            "before": "member = \"allUsers\" or \"allAuthenticatedUsers\"",
            "after": "Use a non-public principal"
        }),
        "remediationType": "replacement"
    }
}

CxPolicy[result] {
    resource := input.document[i].resource.google_cloud_run_service_iam_member[name]
    common_lib.valid_key(resource, "member")
    member := resource.member
    member == "allUsers"
    result := {
        "documentId": input.document[i].id,


        "searchKey": sprintf("google_cloud_run_service_iam_member[{{%s}}].member", [name]),
        "searchLine": common_lib.build_search_line(["resource", "google_cloud_run_service_iam_member", name, "member"], []),
        "issueType": "IncorrectValue",
        "keyExpectedValue": "IAM member should not be a public principal",
        "keyActualValue": sprintf("Found public principal: %s", [member]),
        "remediation": json.marshal({
            "before": "member = \"allUsers\" or \"allAuthenticatedUsers\"",
            "after": "Use a non-public principal"
        }),
        "remediationType": "replacement"
    }
}

CxPolicy[result] {
    resource := input.document[i].resource.google_cloud_run_service_iam_binding[name]
    common_lib.valid_key(resource, "members")
    member := resource.members[_]
    member == "allAuthenticatedUsers"
    result := {
        "documentId": input.document[i].id,


        "searchKey": sprintf("google_cloud_run_service_iam_binding[{{%s}}].members", [name]),
        "searchLine": common_lib.build_search_line(["resource", "google_cloud_run_service_iam_binding", name, "members"], []),
        "issueType": "IncorrectValue",
        "keyExpectedValue": "IAM binding should not include public principals",
        "keyActualValue": "Public principal found in members",
        "remediation": json.marshal({
            "before": "members includes \"allUsers\" or \"allAuthenticatedUsers\"",
            "after": "Remove public principals from members"
        }),
        "remediationType": "replacement"
    }
}

CxPolicy[result] {
    resource := input.document[i].resource.google_cloud_run_service_iam_binding[name]
    common_lib.valid_key(resource, "members")
    member := resource.members[_]
    member == "allUsers"
    result := {
        "documentId": input.document[i].id,


        "searchKey": sprintf("google_cloud_run_service_iam_binding[{{%s}}].members", [name]),
        "searchLine": common_lib.build_search_line(["resource", "google_cloud_run_service_iam_binding", name, "members"], []),
        "issueType": "IncorrectValue",
        "keyExpectedValue": "IAM binding should not include public principals",
        "keyActualValue": "Public principal found in members",
        "remediation": json.marshal({
            "before": "members includes \"allUsers\" or \"allAuthenticatedUsers\"",
            "after": "Remove public principals from members"
        }),
        "remediationType": "replacement"
    }
}
