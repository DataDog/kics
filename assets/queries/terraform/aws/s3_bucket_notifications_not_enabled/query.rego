package Cx

import data.generic.common as common_lib
import data.generic.terraform as tf_lib

CxPolicy[result] {
    bucket := input.document[i].resource.aws_s3_bucket[bucketName]

    not tf_lib.has_target_resource(bucketName, "aws_s3_bucket_notification")

    result := {
        "documentId": input.document[i].id,
        "resourceType": "aws_s3_bucket",
        "resourceName": tf_lib.get_specific_resource_name(bucket, "aws_s3_bucket", bucketName),
        "searchKey": sprintf("aws_s3_bucket[%s]", [bucketName]),
        "issueType": "MissingBlock",
        "keyExpectedValue": "S3 bucket should be associated with an aws_s3_bucket_notification resource",
        "keyActualValue": "No aws_s3_bucket_notification resource refers to this bucket",
        "searchLine": common_lib.build_search_line(["resource", "aws_s3_bucket", bucketName], [])
    }
}
