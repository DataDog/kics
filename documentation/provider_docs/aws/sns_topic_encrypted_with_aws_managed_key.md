---
title: "SNS Topic Encrypted With AWS Managed Key"
meta:
  name: "aws/sns_topic_encrypted_with_aws_managed_key"
  id: "b1a72f66-2236-4f3b-87ba-0da1b366956f"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Encryption"
---

## Metadata
**Name:** `aws/sns_topic_encrypted_with_aws_managed_key`

**Id:** `b1a72f66-2236-4f3b-87ba-0da1b366956f`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Encryption

## Description
SNS (Simple Notification Service) Topic should be encrypted with customer-managed KMS keys instead of AWS managed keys

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/sns_topic#kms_master_key_id)

## Non-Compliant Code Examples
```terraform
provider "aws" {
  region = "us-east-1"
}

data "aws_kms_key" "by_alias" {
  key_id = "alias/aws/sns"
}

resource "aws_sns_topic" "test" {
  name              = "sns_ecnrypted"
  kms_master_key_id = data.aws_kms_key.by_alias.arn
}

```

```terraform
resource "aws_sns_topic" "user_updates" {
  name              = "user-updates-topic"
  kms_master_key_id = "alias/aws/sns"
}

```

## Compliant Code Examples
```terraform
provider "aws2" {
  region = "us-east-1"
}

resource "aws_sns_topic" "test2" {
  name              = "sns_ecnrypted"
  kms_master_key_id = "alias/MyAlias"
}

```