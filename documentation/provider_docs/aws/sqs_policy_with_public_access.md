---
title: "SQS Policy With Public Access"
meta:
  name: "aws/sqs_policy_with_public_access"
  id: "730675f9-52ed-49b6-8ead-0acb5dd7df7f"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `aws/sqs_policy_with_public_access`

**Id:** `730675f9-52ed-49b6-8ead-0acb5dd7df7f`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Access Control

## Description
Checks for dangerous permissions in Action statements in an SQS Queue Policy. This is deemed a potential security risk as it would allow various attacks to the queue

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/sqs_queue_policy)

## Non-Compliant Code Examples
```terraform
resource "aws_sqs_queue" "q" {
  name = "examplequeue"
}

resource "aws_sqs_queue_policy" "test" {
  queue_url = aws_sqs_queue.q.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Id": "Queue1_Policy_UUID",
  "Statement": [{
      "Sid":"Queue1_AnonymousAccess_AllActions_AllowlistIP",
      "Effect": "Allow",
      "Principal": "*",
      "Action": "sqs:*",
      "Resource": "arn:aws:sqs:*:111122223333:queue1",
      "Condition" : {
        "IpAddress" : {
            "aws:SourceIp":"192.168.143.0/24"
        }
      }
  }]
}
EOF
}

resource "aws_sqs_queue" "q_aws_array" {
  name = "examplequeue_aws_array"
}

resource "aws_sqs_queue" "q_aws" {
  name = "examplequeue_aws"
}

resource "aws_sqs_queue_policy" "test_aws" {
  queue_url = aws_sqs_queue.q_aws.id

  policy = <<EOF
{
   "Version": "2012-10-17",
   "Id": "Queue1_Policy_UUID",
   "Statement": [{
      "Sid":"Queue1_AnonymousAccess_AllActions_AllowlistIP",
      "Effect": "Allow",
      "Principal": {
        "AWS": "*"
      },
      "Action": "sqs:*",
      "Resource": "arn:aws:sqs:*:111122223333:queue1",
      "Condition" : {
         "IpAddress" : {
            "aws:SourceIp":"192.168.143.0/24"
         }
      }
   }]
}
EOF
}

resource "aws_sqs_queue_policy" "test_aws_array" {
  queue_url = aws_sqs_queue.q_aws_array.id

  policy = <<EOF
{
   "Version": "2012-10-17",
   "Id": "Queue1_Policy_UUID",
   "Statement": [{
      "Sid":"Queue1_AnonymousAccess_AllActions_AllowlistIP",
      "Effect": "Allow",
      "Principal": {
        "AWS": ["*"]
      },
      "Action": "sqs:*",
      "Resource": "arn:aws:sqs:*:111122223333:queue1",
      "Condition" : {
         "IpAddress" : {
            "aws:SourceIp":"192.168.143.0/24"
         }
      }
   }]
}
EOF
}

```

## Compliant Code Examples
```terraform
resource "aws_sqs_queue" "q" {
  name = "examplequeue"
}

resource "aws_sqs_queue_policy" "test" {
  queue_url = aws_sqs_queue.q.id

  policy = <<POLICY
{
   "Version": "2012-10-17",
   "Id": "Queue1_Policy_UUID",
   "Statement": [{
      "Sid":"Queue1_AnonymousAccess_AllActions_AllowlistIP",
      "Effect": "Allow",
      "Principal": "SOMETHING",
      "Action": "sqs:*",
      "Resource": "arn:aws:sqs:*:111122223333:queue1",
      "Condition" : {
         "IpAddress" : {
            "aws:SourceIp":"192.168.143.0/24"
         }
      }
   }]
}
POLICY
}
```