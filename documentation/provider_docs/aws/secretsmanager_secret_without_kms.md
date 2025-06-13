---
title: "Secretsmanager Secret Without KMS"
meta:
  name: "aws/secretsmanager_secret_without_kms"
  id: "a2f548f2-188c-4fff-b172-e9a6acb216bd"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Encryption"
---

## Metadata
**Name:** `aws/secretsmanager_secret_without_kms`

**Id:** `a2f548f2-188c-4fff-b172-e9a6acb216bd`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Encryption

## Description
AWS Secretmanager should use AWS KMS customer master key (CMK) to encrypt the secret values in the versions stored in the secret

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/secretsmanager_secret#kms_key_id)

## Non-Compliant Code Examples
```terraform
resource "aws_secretsmanager_secret" "example" {
  name = "example"
}

```

## Compliant Code Examples
```terraform
resource "aws_secretsmanager_secret" "example" {
  name = "example"
  kms_key_id = "arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"
}

```