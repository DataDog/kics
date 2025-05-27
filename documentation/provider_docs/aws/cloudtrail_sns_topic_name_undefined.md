---
title: "CloudTrail SNS Topic Name Undefined"
meta:
  name: "aws/cloudtrail_sns_topic_name_undefined"
  id: "482b7d26-0bdb-4b5f-bf6f-545826c0a3dd"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Observability"
---

## Metadata
**Name:** `aws/cloudtrail_sns_topic_name_undefined`

**Id:** `482b7d26-0bdb-4b5f-bf6f-545826c0a3dd`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Observability

## Description
Check if SNS topic name is set for CloudTrail

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudtrail)

## Non-Compliant Code Examples
```terraform
resource "aws_cloudtrail" "positive1" {
  # ... other configuration ...
}

resource "aws_cloudtrail" "positive2" {
  # ... other configuration ...

  sns_topic_name = null
}
```

## Compliant Code Examples
```terraform
resource "aws_cloudtrail" "negative1" {
  # ... other configuration ...

  sns_topic_name = "some-topic"
}
```