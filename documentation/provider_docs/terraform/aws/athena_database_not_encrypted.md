---
title: "Athena Database Not Encrypted"
meta:
  name: "aws/athena_database_not_encrypted"
  id: "b2315cae-b110-4426-81e0-80bb8640cdd3"
  cloud_provider: "aws"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata
**Name:** `aws/athena_database_not_encrypted`
**Id:** `b2315cae-b110-4426-81e0-80bb8640cdd3`
**Cloud Provider:** aws
**Severity:** High
**Category:** Encryption
## Description
AWS Athena Databases store query results in S3 buckets, and when not encrypted, sensitive data may be exposed to unauthorized access, potentially leading to data breaches and compliance violations. Encryption at rest protects this data using keys managed either by AWS or customer-managed KMS keys. To secure your implementation, add an encryption_configuration block to your aws_athena_database resource as shown below:

```terraform
resource "aws_athena_database" "secure_example" {
  name   = "database_name"
  bucket = aws_s3_bucket.example.bucket

  encryption_configuration {
    encryption_option = "SSE_KMS"
    kms_key           = "your_kms_key_arn"
  }
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/athena_database#encryption_configuration)


## Compliant Code Examples
```terraform
resource "aws_s3_bucket" "hoge" {
  bucket = "hoge"
}

resource "aws_athena_database" "hoge" {
  name   = "database_name"
  bucket = aws_s3_bucket.hoge.bucket

  encryption_configuration {
    encryption_option = "SSE_KMS"
    kms_key           = "SSE_KMS"
 }
}

```
## Non-Compliant Code Examples
```terraform
resource "aws_s3_bucket" "hoge" {
  bucket = "hoge"
}

resource "aws_athena_database" "hoge" {
  name   = "database_name"
  bucket = aws_s3_bucket.hoge.bucket
}

```