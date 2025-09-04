package Cx

import data.generic.terraform as tf_lib
import data.generic.common as common_lib

# Check for IAM member on KMS key rings
CxPolicy[result] {
    resource := input.document[i].resource.google_kms_key_ring_iam_member[name]
    common_lib.valid_key(resource, "member")
    member := resource.member
    member == "allUsers"

    result := {
        "documentId": input.document[i].id,


        "searchKey": sprintf("google_kms_key_ring_iam_member[{{%s}}].member", [name]),
        "searchLine": common_lib.build_search_line(["resource", "google_kms_key_ring_iam_member", name, "member"], []),
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

# Check for IAM member on KMS key rings
CxPolicy[result] {
    resource := input.document[i].resource.google_kms_key_ring_iam_member[name]
    common_lib.valid_key(resource, "member")
    member := resource.member
    member == "allAuthenticatedUsers"

    result := {
        "documentId": input.document[i].id,


        "searchKey": sprintf("google_kms_key_ring_iam_member[{{%s}}].member", [name]),
        "searchLine": common_lib.build_search_line(["resource", "google_kms_key_ring_iam_member", name, "member"], []),
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

# Check for IAM binding on KMS key rings
CxPolicy[result] {
    resource := input.document[i].resource.google_kms_key_ring_iam_binding[name]
    common_lib.valid_key(resource, "members")
    m := resource.members[_]
    m == "allUsers"

    result := {
        "documentId": input.document[i].id,


        "searchKey": sprintf("google_kms_key_ring_iam_binding[{{%s}}].members", [name]),
        "searchLine": common_lib.build_search_line(["resource", "google_kms_key_ring_iam_binding", name, "members"], []),
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

# Check for IAM binding on KMS key rings
CxPolicy[result] {
    resource := input.document[i].resource.google_kms_key_ring_iam_binding[name]
    common_lib.valid_key(resource, "members")
    m := resource.members[_]
    m == "allAuthenticatedUsers"

    result := {
        "documentId": input.document[i].id,


        "searchKey": sprintf("google_kms_key_ring_iam_binding[{{%s}}].members", [name]),
        "searchLine": common_lib.build_search_line(["resource", "google_kms_key_ring_iam_binding", name, "members"], []),
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
