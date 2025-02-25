package Cx

import data.generic.common as common_lib
import data.generic.terraform as tf_lib

flag_exists(flags) {
    f := flags[_]
    f.name == "skip_show_database"
}

CxPolicy[result] {
    resource := input.document[i].resource.google_sql_database_instance[name]

    common_lib.valid_key(resource, "settings")
    settings := resource.settings

    # For each flag in resource.settings.database_flags
    flag := object.get(settings, "database_flags", "")

    flag.name == "skip_show_database"
    flag.value == "off"

    result := {
        "documentId": input.document[i].id,
        "resourceType": "google_sql_database_instance",
        "resourceName": tf_lib.get_resource_name(resource, name),
        "searchKey": sprintf("google_sql_database_instance[%s].settings.database_flags.skip_show_database", [name]),
        "searchLine": common_lib.build_search_line(["resource", "google_sql_database_instance", name], ["settings", "database_flags", "skip_show_database"]),
        "issueType": "MissingValue",
        "keyExpectedValue": "'settings.database_flags.skip_show_database' must be present",
        "keyActualValue": "'skip_show_database' flag is missing",
        "remediation": json.marshal({
            "before": "database_flags without 'skip_show_database'",
            "after":  "Add { name = \"skip_show_database\", value = \"<desired_value>\" } to settings.database_flags"
        }),
        "remediationType": "addition"
    }
}

CxPolicy[result] {
    resource := input.document[i].resource.google_sql_database_instance[name]
    common_lib.valid_key(resource, "settings")
    settings := common_lib.as_array(resource.settings)[_]

    # Fail if the settings block does not contain a database_flags key
    not common_lib.valid_key(settings, "database_flags")

    result := {
        "documentId": input.document[i].id,
        "resourceType": "google_sql_database_instance",
        "resourceName": tf_lib.get_resource_name(resource, name),
        "searchKey": sprintf("google_sql_database_instance[%s].settings.database_flags", [name]),
        "searchLine": common_lib.build_search_line(["resource", "google_sql_database_instance", name], ["settings", "database_flags"]),
        "issueType": "MissingValue",
        "keyExpectedValue": "'settings.database_flags.skip_show_database' must be present",
        "keyActualValue": "'database_flags' block is missing",
        "remediation": json.marshal({
            "before": "No database_flags block",
            "after": "Add database_flags block with { name = \"skip_show_database\", value = \"<desired_value>\" }"
        }),
        "remediationType": "addition"
    }
}

CxPolicy[result] {
    resource := input.document[i].resource.google_sql_database_instance[name]
    common_lib.valid_key(resource, "settings")
    settings := common_lib.as_array(resource.settings)[_]
    common_lib.valid_key(settings, "database_flags")
    flags := common_lib.as_array(settings.database_flags)

    # Fail if no flag in the database_flags list has the name "skip_show_database"
    not flag_exists(flags)

    result := {
        "documentId": input.document[i].id,
        "resourceType": "google_sql_database_instance",
        "resourceName": tf_lib.get_resource_name(resource, name),
        "searchKey": sprintf("google_sql_database_instance[%s].settings.database_flags", [name]),
        "searchLine": common_lib.build_search_line(["resource", "google_sql_database_instance", name], ["settings", "database_flags", "skip_show_database"]),
        "issueType": "MissingValue",
        "keyExpectedValue": "'settings.database_flags.skip_show_database' must be present",
        "keyActualValue": "'skip_show_database' flag is missing",
        "remediation": json.marshal({
            "before": "database_flags missing 'skip_show_database'",
            "after": "Add { name = \"skip_show_database\", value = \"<desired_value>\" } to settings.database_flags"
        }),
        "remediationType": "addition"
    }
}
