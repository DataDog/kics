---
title: "S3 Bucket Without Ignore Public ACL"
meta:
  name: "terraform/s3_bucket_without_ignore_public_acl"
  id: "4fa66806-0dd9-4f8d-9480-3174d39c7c91"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `terraform/s3_bucket_without_ignore_public_acl`
**Id:** `4fa66806-0dd9-4f8d-9480-3174d39c7c91`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Insecure Configurations
## Description
When the `ignore_public_acls` attribute in the `aws_s3_bucket_public_access_block` resource is set to `false`, S3 buckets may still honor any public access permissions granted by existing or future bucket or object ACLs. This creates a risk of sensitive data being inadvertently exposed to the public internet if a user or process applies a permissive ACL. To prevent this, `ignore_public_acls` should be set to `true`, which ensures that all public ACLs are ignored and cannot be used to grant public access.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_public_access_block)

## Non-Compliant Code Examples
```aws
resource "aws_s3_bucket" "positive1" {
  bucket = "example"
}

resource "aws_s3_bucket_public_access_block" "positive2" {
  bucket = aws_s3_bucket.example.id

  block_public_acls   = true
  block_public_policy = true
  ignore_public_acls  = false
}

```

```aws
resource "aws_s3_bucket" "positive1" {
  bucket = "example"
}

resource "aws_s3_bucket_public_access_block" "positive2" {
  bucket = aws_s3_bucket.example.id

  block_public_acls   = true
  block_public_policy = true
}

```

```aws
module "s3_bucket" {
  source = "terraform-aws-modules/s3-bucket/aws"
  version = "3.7.0"

  bucket = "my-s3-bucket"
  acl    = "private"

  versioning = {
    enabled = true
  }
}

```

## Compliant Code Examples
```aws
resource "aws_s3_bucket" "negative1" {
  bucket = "example"
}

resource "aws_s3_bucket_public_access_block" "negative2" {
  bucket = aws_s3_bucket.example.id

  block_public_acls   = true
  block_public_policy = true
  ignore_public_acls  = true
}

```

```aws
module "s3_bucket" {
  source = "terraform-aws-modules/s3-bucket/aws"
  version = "3.7.0"

  bucket = "my-s3-bucket"
  acl    = "private"
  ignore_public_acls  = true

  versioning = {
    enabled = true
  }
}

```