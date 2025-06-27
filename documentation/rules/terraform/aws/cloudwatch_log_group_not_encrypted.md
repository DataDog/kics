---
title: "CloudWatch Log Group Without KMS"
meta:
  name: "aws/cloudwatch_log_group_not_encrypted"
  id: "0afbcfe9-d341-4b92-a64c-7e6de0543879"
  display_name: "CloudWatch Log Group Without KMS"
  cloud_provider: "aws"
  platform: "Terraform"
  severity: "MEDIUM"
  category: "Encryption"
---
## Metadata

**Name:** `aws/cloudwatch_log_group_not_encrypted`

**Query Name** `CloudWatch Log Group Without KMS`

**Id:** `0afbcfe9-d341-4b92-a64c-7e6de0543879`

**Cloud Provider:** aws

**Platform** Terraform

**Severity:** Medium

**Category:** Encryption

## Description
AWS CloudWatch Log groups should use KMS encryption to protect sensitive log data at rest. When the `aws_cloudwatch_log_group` resource is defined without the `kms_key_id` attribute, as in:

```
resource "aws_cloudwatch_log_group" "negative1" {
  name = "Yada"
  retention_in_days = 1
}
```

logs are stored unencrypted or with default encryption, making them vulnerable to unauthorized access if AWS accounts or storage is compromised. By specifying the `kms_key_id` attribute, you ensure that log data is encrypted with a customer-managed AWS KMS key, reducing the risk of data exposure.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_log_group)


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