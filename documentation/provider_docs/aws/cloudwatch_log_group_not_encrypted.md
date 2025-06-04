---
title: "CloudWatch Log Group Without KMS"
meta:
  name: "aws/cloudwatch_log_group_not_encrypted"
  id: "0afbcfe9-d341-4b92-a64c-7e6de0543879"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Encryption"
---

## Metadata
**Name:** `aws/cloudwatch_log_group_not_encrypted`

**Id:** `0afbcfe9-d341-4b92-a64c-7e6de0543879`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Encryption

## Description
AWS CloudWatch Log groups should be encrypted using KMS

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_log_group)

## Non-Compliant Code Examples
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

## Compliant Code Examples
```terraform
resource "aws_cloudwatch_log_group" "negative1" {
  name = "Yada"

  tags = {
    Environment = "production"
    Application = "serviceA"
  }

  retention_in_days = 1
  kms_key_id = "arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"
}

```