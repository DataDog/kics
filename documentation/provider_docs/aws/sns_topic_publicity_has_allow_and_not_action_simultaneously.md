---
title: "SNS Topic Publicity Has Allow and NotAction Simultaneously"
meta:
  name: "aws/sns_topic_publicity_has_allow_and_not_action_simultaneously"
  id: "5ea624e4-c8b1-4bb3-87a4-4235a776adcc"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `aws/sns_topic_publicity_has_allow_and_not_action_simultaneously`

**Id:** `5ea624e4-c8b1-4bb3-87a4-4235a776adcc`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Access Control

## Description
SNS topic Publicity should not have 'Effect: Allow' and argument 'NotAction' at the same time. If it has 'Effect: Allow', the argument stated should be 'Action'.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/sns_topic_policy)

## Non-Compliant Code Examples
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
  "Id": "MYPOLICYTEST",
  "Statement": [
    {
      "NotAction": "s3:DeleteBucket",
      "Resource": "arn:aws:s3:::*",
      "Sid": "MyStatementId",
      "Effect": "Allow"
    }
  ]
}
POLICY

   server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        kms_master_key_id = aws_kms_key.mykey.arn
        sse_algorithm     = "aws:kms"
      }
    }
  }
}

```

```terraform
resource "aws_sns_topic" "positive1" {
  name = "my-topic-with-policy"
}

resource "aws_sns_topic_policy" "positive2" {
  arn = aws_sns_topic.test.arn

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Id": "MYPOLICYTEST",
  "Statement": [
    {
      "NotAction": "s3:DeleteBucket",
      "Resource": "arn:aws:s3:::*",
      "Sid": "MyStatementId",
      "Effect": "Allow"
    }
  ]
}
POLICY
}

```

## Compliant Code Examples
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
  "Id": "MYPOLICYTEST",
  "Statement": [
    {
      "Action": "s3:DeleteBucket",
      "Resource": "arn:aws:s3:::*",
      "Sid": "MyStatementId",
      "Effect": "Allow"
    }
  ]
}
POLICY

   server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        kms_master_key_id = aws_kms_key.mykey.arn
        sse_algorithm     = "aws:kms"
      }
    }
  }
}

```

```terraform
resource "aws_sns_topic" "negative1" {
  name = "my-topic-with-policy"
}

resource "aws_sns_topic_policy" "negative2" {
  arn = aws_sns_topic.test.arn

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Id": "MYPOLICYTEST",
  "Statement": [
    {
      "Action": "s3:DeleteBucket",
      "Resource": "arn:aws:s3:::*",
      "Sid": "MyStatementId",
      "Effect": "Allow"
    }
  ]
}
POLICY
}

```