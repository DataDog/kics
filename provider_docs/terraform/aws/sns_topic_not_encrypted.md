---
title: "SNS Topic Not Encrypted"
meta:
  name: "terraform/sns_topic_not_encrypted"
  id: "28545147-2fc6-42d5-a1f9-cf226658e591"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata
**Name:** `terraform/sns_topic_not_encrypted`
**Id:** `28545147-2fc6-42d5-a1f9-cf226658e591`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Encryption
## Description
Amazon SNS topics should be encrypted at rest using AWS KMS to protect sensitive message content. Without encryption, data stored in SNS topics is vulnerable to unauthorized access if the service or underlying storage is compromised. Use the `kms_master_key_id` attribute to specify a KMS key for encryption as shown in the secure example: `kms_master_key_id = "alias/MyAlias"`. Leaving this attribute empty or unspecified (as in `kms_master_key_id = ""`) results in unencrypted SNS topics, exposing sensitive data to potential security breaches.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/sns_topic#kms_master_key_id)

## Non-Compliant Code Examples
```aws
resource "aws_sns_topic" "user_updates" {
  name              = "user-updates-topic"
  kms_master_key_id = ""
}

```

```aws
provider "aws" {
  region = "us-east-1"
}

resource "aws_sns_topic" "test" {
  name = "sns_not_ecnrypted"
}

```

## Compliant Code Examples
```aws
provider "aws2" {
  region = "us-east-1"
}

resource "aws_sns_topic" "test2" {
  name              = "sns_ecnrypted"
  kms_master_key_id = "alias/MyAlias"
}

```