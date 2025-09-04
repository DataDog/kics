package Cx

import data.generic.terraform as tf_lib
import data.generic.common as common_lib

# Check for google_pubsub_topic_iam_member
CxPolicy[result] {
    resource := input.document[i].resource.google_pubsub_topic_iam_member[name]
    common_lib.valid_key(resource, "member")
    member := resource.member
    member == "allUsers"

    result := {
        "documentId": input.document[i].id,


        "searchKey": sprintf("google_pubsub_topic_iam_member[{{%s}}].member", [name]),
        "searchLine": common_lib.build_search_line(["resource", "google_pubsub_topic_iam_member", name, "member"], []),
        "issueType": "IncorrectValue",
        "keyExpectedValue": "member should not be a public principal",
        "keyActualValue": sprintf("member is '%s'", [member]),
        "remediation": json.marshal({
            "before": "member = \"allUsers\" or \"allAuthenticatedUsers\"",
            "after": "Use a non-public principal"
        }),
        "remediationType": "replacement"
    }
}

# Check for google_pubsub_topic_iam_member
CxPolicy[result] {
    resource := input.document[i].resource.google_pubsub_topic_iam_member[name]
    common_lib.valid_key(resource, "member")
    member := resource.member
    member == "allAuthenticatedUsers"

    result := {
        "documentId": input.document[i].id,


        "searchKey": sprintf("google_pubsub_topic_iam_member[{{%s}}].member", [name]),
        "searchLine": common_lib.build_search_line(["resource", "google_pubsub_topic_iam_member", name, "member"], []),
        "issueType": "IncorrectValue",
        "keyExpectedValue": "member should not be a public principal",
        "keyActualValue": sprintf("member is '%s'", [member]),
        "remediation": json.marshal({
            "before": "member = \"allUsers\" or \"allAuthenticatedUsers\"",
            "after": "Use a non-public principal"
        }),
        "remediationType": "replacement"
    }
}

# Check for google_pubsub_topic_iam_binding
CxPolicy[result] {
    resource := input.document[i].resource.google_pubsub_topic_iam_binding[name]
    common_lib.valid_key(resource, "members")
    member := resource.members[_]
    member == "allUsers"

    result := {
        "documentId": input.document[i].id,


        "searchKey": sprintf("google_pubsub_topic_iam_binding[{{%s}}].members", [name]),
        "searchLine": common_lib.build_search_line(["resource", "google_pubsub_topic_iam_binding", name, "members"], []),
        "issueType": "IncorrectValue",
        "keyExpectedValue": "members should not include public principals",
        "keyActualValue": "Public principal found in members",
        "remediation": json.marshal({
            "before": "members includes \"allUsers\" or \"allAuthenticatedUsers\"",
            "after": "Remove public principals from members"
        }),
        "remediationType": "replacement"
    }
}

# Check for google_pubsub_topic_iam_binding
CxPolicy[result] {
    resource := input.document[i].resource.google_pubsub_topic_iam_binding[name]
    common_lib.valid_key(resource, "members")
    member := resource.members[_]
    member == "allAuthenticatedUsers"

    result := {
        "documentId": input.document[i].id,


        "searchKey": sprintf("google_pubsub_topic_iam_binding[{{%s}}].members", [name]),
        "searchLine": common_lib.build_search_line(["resource", "google_pubsub_topic_iam_binding", name, "members"], []),
        "issueType": "IncorrectValue",
        "keyExpectedValue": "members should not include public principals",
        "keyActualValue": "Public principal found in members",
        "remediation": json.marshal({
            "before": "members includes \"allUsers\" or \"allAuthenticatedUsers\"",
            "after": "Remove public principals from members"
        }),
        "remediationType": "replacement"
    }
}
