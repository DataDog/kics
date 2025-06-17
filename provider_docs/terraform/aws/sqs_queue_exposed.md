---
title: "SQS Queue Exposed"
meta:
  name: "terraform/sqs_queue_exposed"
  id: "abb06e5f-ef9a-4a99-98c6-376d396bfcdf"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/sqs_queue_exposed`
**Id:** `abb06e5f-ef9a-4a99-98c6-376d396bfcdf`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Access Control
## Description
An SQS Queue is considered exposed when its policy allows access from any principal by setting 'Principal' to '*'. This creates a security vulnerability where unauthorized users can send messages to your queue, potentially leading to information disclosure or denial of service attacks.

To secure your SQS queue, avoid using wildcard principals in your policy statements. Instead, explicitly specify the ARNs of services or resources that should have access to the queue, as shown in the following secure example:

```
{
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "sqs:SendMessage",
      "Resource": "${aws_sqs_queue.q.arn}",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "${aws_sns_topic.example.arn}"
        }
      }
    }
  ]
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/sqs_queue#policy)

## Non-Compliant Code Examples
```aws
resource "aws_sqs_queue" "positive1" {
  name = "examplequeue"

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Id": "sqspolicy",
  "Statement": [
    {
      "Sid": "First",
      "Effect": "Allow",
      "Principal": "*",
      "Action": "sqs:SendMessage",
      "Resource": "${aws_sqs_queue.q.arn}",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "${aws_sns_topic.example.arn}"
        }
      }
    }
  ]
}
POLICY
}

```

```aws
module "user_queue" {
  source  = "terraform-aws-modules/sqs/aws"
  version = "~> 2.0"

  name = "user"

  tags = {
    Service     = "user"
    Environment = "dev"
  }

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Id": "sqspolicy",
  "Statement": [
    {
      "Sid": "First",
      "Effect": "Allow",
      "Principal": "*",
      "Action": "sqs:SendMessage",
      "Resource": "${aws_sqs_queue.q.arn}",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "${aws_sns_topic.example.arn}"
        }
      }
    }
  ]
}
POLICY

}

```

## Compliant Code Examples
```aws
resource "aws_sqs_queue" "negative1" {
  name = "examplequeue"

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Id": "sqspolicy",
  "Statement": [
    {
      "Sid": "First",
      "Effect": "Allow",
      "Action": "sqs:SendMessage",
      "Resource": "${aws_sqs_queue.q.arn}",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "${aws_sns_topic.example.arn}"
        }
      }
    }
  ]
}
POLICY

}

```

```aws
module "user_queue" {
  source  = "terraform-aws-modules/sqs/aws"
  version = "~> 2.0"

  name = "user"

  tags = {
    Service     = "user"
    Environment = "dev"
  }

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Id": "sqspolicy",
  "Statement": [
    {
      "Sid": "First",
      "Effect": "Allow",
      "Action": "sqs:SendMessage",
      "Resource": "${aws_sqs_queue.q.arn}",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "${aws_sns_topic.example.arn}"
        }
      }
    }
  ]
}
POLICY

}

```