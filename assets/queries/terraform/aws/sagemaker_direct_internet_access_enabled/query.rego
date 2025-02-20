package Cx

import data.generic.terraform as tf_lib
import data.generic.common as common_lib

CxPolicy[result] {
    resource := input.document[i].resource.aws_sagemaker_notebook_instance[name]
    common_lib.valid_key(resource, "direct_internet_access")

    resource.direct_internet_access != "Disabled"

    result := {
        "documentId": input.document[i].id,
        "resourceType": "aws_sagemaker_notebook_instance",
        "resourceName": tf_lib.get_resource_name(resource, name),
        "searchKey": sprintf("aws_sagemaker_notebook_instance[{{%s}}].direct_internet_access", [name]),
        "searchLine": common_lib.build_search_line(["resource", "aws_sagemaker_notebook_instance", name, "direct_internet_access"], []),
        "issueType": "IncorrectValue",
        "keyExpectedValue": "aws_sagemaker_notebook_instance.direct_internet_access should be 'Disabled'",
        "keyActualValue": sprintf("aws_sagemaker_notebook_instance.direct_internet_access is '%s'", [resource.direct_internet_access]),
        "remediation": json.marshal({
            "before": "direct_internet_access = \"Enabled\"",
            "after": "direct_internet_access = \"Disabled\""
        }),
        "remediationType": "replacement"
    }
}
