---
title: "CloudWatch Without Retention Period Specified"
meta:
  name: "aws/cloudwatch_without_retention_period_specified"
  id: "ef0b316a-211e-42f1-888e-64efe172b755"
  cloud_provider: "aws"
  severity: "INFO"
  category: "Observability"
---

## Metadata
**Name:** `aws/cloudwatch_without_retention_period_specified`

**Id:** `ef0b316a-211e-42f1-888e-64efe172b755`

**Cloud Provider:** aws

**Severity:** Info

**Category:** Observability

## Description
AWS CloudWatch Log groups should have retention days specified

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_log_group)

## Non-Compliant Code Examples
```terraform
resource "aws_cloudwatch_log_group" "positive1" {
  name = "Yada"

  tags = {
    Environment = "production"
    Application = "serviceA"
  }
}

resource "aws_cloudwatch_log_group" "positive2" {
  name = "Yada"

  tags = {
    Environment = "production"
    Application = "serviceA"
  }

  retention_in_days = 0
}

```

## Compliant Code Examples
```terraform
resource "aws_cloudwatch_log_group" "negative1" {
  name = "Yada"

  tags = {
    Environment = "production"
    Application = "serviceA"
  }

  retention_in_days = 1
}

```