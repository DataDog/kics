---
title: "SNS topic is publicly accessible"
group_id: "rules/terraform/aws"
meta:
  name: "aws/sns_topic_is_publicly_accessible"
  id: "b26d2b7e-60f6-413d-a3a1-a57db24aa2b3"
  display_name: "SNS topic is publicly accessible"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "CRITICAL"
  category: "Access Control"
---
## Metadata

**Id:** `b26d2b7e-60f6-413d-a3a1-a57db24aa2b3`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Critical

**Category:** Access Control

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/sns_topic)

### Description

 This check verifies that Amazon SNS topic policies do not allow public access by having wildcard principals in their IAM policies. When an SNS topic policy includes a principal with wildcard (`*`) or allows anonymous access, it makes the topic publicly accessible to any AWS account, potentially exposing sensitive information or allowing unauthorized message publishing/consumption.

Secure configuration requires specifying explicit IAM principals rather than using wildcards. For example, instead of using `"AWS": "*"` which grants access to anyone, use a specific account ARN like `"AWS": "arn:aws:iam::account_number:root"` to limit access to authorized entities only. This prevents unauthorized access to your SNS topics and their messages.


## Compliant Code Examples
```terraform
resource "aws_sns_topic" "negative1" {
policy = <<EOF
{
"Version": "2012-10-17",
"Statement": [
{
"Action": "*",
"Principal": {
  "AWS": "arn:aws:iam::##account_number##:root"
},
"Effect": "Allow",
"Sid": ""
}
]
}
EOF
}


```

```terraform
module "s3_bucket" {
  source = "terraform-aws-modules/s3-bucket/aws"
  version = "3.7.0"

  bucket = "my-s3-bucket"
  acl    = "private"

  versioning = {
    enabled = true
  }

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Id": "default_sns_policy",
  "Statement": [
    {
      "Sid": "default-statement-id",
      "Effect": "Allow",
      "Action": [
        "SNS:Publish"
      ],
      "Resource": "${aws_sns_topic.example.arn}",
      "Condition": {
        "ArnEquals": {
          "AWS": "${aws_sns_topic.example.arn}"
        }
      },
      "Principal": {
        "AWS": "010:role/test"
      }
    }
  ]
}
POLICY

}
```
## Non-Compliant Code Examples
```terraform
resource "aws_sns_topic" "positive1" {
policy = <<EOF
{
"Version": "2012-10-17",
"Statement": [
{
"Action": "*",
"Principal": {
  "AWS": "*"
},
"Effect": "Allow",
"Sid": ""
}
]
}
EOF
}

```

```terraform
module "s3_bucket" {
  source = "terraform-aws-modules/s3-bucket/aws"
  version = "3.7.0"

  bucket = "my-s3-bucket"
  acl    = "private"

  versioning = {
    enabled = true
  }

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Id": "default_sns_policy",
  "Statement": [
    {
      "Sid": "default-statement-id",
      "Effect": "Allow",
      "Action": [
        "SNS:Publish"
      ],
      "Resource": "${aws_sns_topic.example.arn}",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "${aws_kms_key.example.arn}"
        }
      },
      "Principal": {
        "AWS": "*"
      }
    }
  ]
}
POLICY

}
```