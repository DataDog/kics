---
title: "CloudWatch Without Retention Period Specified"
meta:
  name: "terraform/cloudwatch_without_retention_period_specified"
  id: "ef0b316a-211e-42f1-888e-64efe172b755"
  cloud_provider: "terraform"
  severity: "INFO"
  category: "Observability"
---
## Metadata
**Name:** `terraform/cloudwatch_without_retention_period_specified`
**Id:** `ef0b316a-211e-42f1-888e-64efe172b755`
**Cloud Provider:** terraform
**Severity:** Info
**Category:** Observability
## Description
CloudWatch Log Groups in AWS should explicitly specify the `retention_in_days` attribute to control how long log data is stored. If this attribute is omitted or set to `0`, as shown below, logs are retained indefinitely, increasing storage costs and the risk of sensitive data exposure over time:

```
resource "aws_cloudwatch_log_group" "example" {
  name = "application-logs"
  retention_in_days = 0
}
```

A secure configuration, specifying `retention_in_days`, ensures log data is automatically deleted after the defined period, reducing unnecessary data retention and helping meet compliance requirements:

```
resource "aws_cloudwatch_log_group" "example" {
  name = "application-logs"
  retention_in_days = 7
}
```


#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_log_group)

## Non-Compliant Code Examples
```aws
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
```aws
resource "aws_cloudwatch_log_group" "negative1" {
  name = "Yada"

  tags = {
    Environment = "production"
    Application = "serviceA"
  }

  retention_in_days = 1
}

```