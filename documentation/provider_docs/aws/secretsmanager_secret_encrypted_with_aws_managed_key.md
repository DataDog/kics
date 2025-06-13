---
title: "Secretsmanager Secret Encrypted With AWS Managed Key"
meta:
  name: "aws/secretsmanager_secret_encrypted_with_aws_managed_key"
  id: "b0d3ef3f-845d-4b1b-83d6-63a5a380375f"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Encryption"
---

## Metadata
**Name:** `aws/secretsmanager_secret_encrypted_with_aws_managed_key`

**Id:** `b0d3ef3f-845d-4b1b-83d6-63a5a380375f`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Encryption

## Description
Secrets Manager secret should be encrypted with customer-managed KMS keys instead of AWS managed keys

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/secretsmanager_secret#kms_key_id)

## Non-Compliant Code Examples
```terraform
provider "aws" {
  region = "us-east-1"
}

data "aws_kms_key" "by_alias" {
  key_id = "alias/aws/secretsmanager"
}

resource "aws_secretsmanager_secret" "test" {
  name       = "test-cloudrail-1"
  kms_key_id = data.aws_kms_key.by_alias.arn
}

```

```terraform
resource "aws_secretsmanager_secret" "test2" {
  name       = "test-cloudrail-1"
  kms_key_id = "alias/aws/secretsmanager"
}

```

## Compliant Code Examples
```terraform
resource "aws_secretsmanager_secret" "test222" {
  name       = "test-cloudrail-1"
  kms_key_id = "alias/MyAlias"
}


```