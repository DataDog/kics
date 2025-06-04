---
title: "CloudTrail Log Files Not Encrypted With KMS"
meta:
  name: "aws/cloudtrail_log_files_not_encrypted_with_kms"
  id: "5d9e3164-9265-470c-9a10-57ae454ac0c7"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Encryption"
---

## Metadata
**Name:** `aws/cloudtrail_log_files_not_encrypted_with_kms`

**Id:** `5d9e3164-9265-470c-9a10-57ae454ac0c7`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Encryption

## Description
Logs delivered by CloudTrail should be encrypted using KMS to increase security of your CloudTrail

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudtrail#kms_key_id)

## Non-Compliant Code Examples
```terraform
resource "aws_cloudtrail" "positive1" {
  name                          = "npositive_1"
  s3_bucket_name                = "bucketlog_1"
}

```

## Compliant Code Examples
```terraform
resource "aws_cloudtrail" "negative1" {
  name                          = "negative1"
  s3_bucket_name                = "bucketlog1"
  kms_key_id                    = "arn:aws:kms:us-east-2:123456789012:key/12345678-1234-1234-1234-123456789012"
}

```