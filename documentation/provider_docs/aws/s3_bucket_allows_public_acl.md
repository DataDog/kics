---
title: "S3 Bucket Allows Public ACL"
meta:
  name: "aws/s3_bucket_allows_public_acl"
  id: "d0cc8694-fcad-43ff-ac86-32331d7e867f"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `aws/s3_bucket_allows_public_acl`

**Id:** `d0cc8694-fcad-43ff-ac86-32331d7e867f`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Access Control

## Description
S3 bucket allows public ACL

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_public_access_block)

## Non-Compliant Code Examples
```terraform
module "s3_bucket" {
  source = "terraform-aws-modules/s3-bucket/aws"
  version = "3.7.0"

  bucket = "my-s3-bucket"
  acl    = "private"
  restrict_public_buckets = true
  block_public_acls = false

  versioning = {
    enabled = true
  }

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Id": "MYBUCKETPOLICY",
  "Statement": [
    {
      "Sid": "IPAllow",
      "Effect": "Deny",
      "Action": "s3:*",
      "Resource": "arn:aws:s3:::my_tf_test_bucket/*",
      "Condition": {
         "IpAddress": {"aws:SourceIp": "8.8.8.8/32"}
      }
    }
  ]
}
POLICY
}

```

```terraform
module "s3_bucket" {
  source = "terraform-aws-modules/s3-bucket/aws"
  version = "3.7.0"

  bucket = "my-s3-bucket"
  acl    = "private"
  restrict_public_buckets = true

  versioning = {
    enabled = true
  }

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Id": "MYBUCKETPOLICY",
  "Statement": [
    {
      "Sid": "IPAllow",
      "Effect": "Deny",
      "Action": "s3:*",
      "Resource": "arn:aws:s3:::my_tf_test_bucket/*",
      "Condition": {
         "IpAddress": {"aws:SourceIp": "8.8.8.8/32"}
      }
    }
  ]
}
POLICY
}

```

```terraform
resource "aws_s3_bucket" "positive1" {
  bucket = "example"
}

resource "aws_s3_bucket_public_access_block" "positive2" {
  bucket = aws_s3_bucket.example.id

  block_public_acls   = false
  block_public_policy = true
  ignore_public_acls  = false
}

// comment
// comment
// comment
// comment
// comment
resource "aws_s3_bucket_public_access_block" "positive3" {
  bucket = aws_s3_bucket.example.id

  block_public_policy = true
  ignore_public_acls  = false
}
```

## Compliant Code Examples
```terraform
module "s3_bucket" {
  source = "terraform-aws-modules/s3-bucket/aws"
  version = "3.7.0"

  bucket = "my-s3-bucket"
  acl    = "private"
  restrict_public_buckets = true
  block_public_acls = true

  versioning = {
    enabled = true
  }

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Id": "MYBUCKETPOLICY",
  "Statement": [
    {
      "Sid": "IPAllow",
      "Effect": "Deny",
      "Action": "s3:*",
      "Resource": "arn:aws:s3:::my_tf_test_bucket/*",
      "Condition": {
         "IpAddress": {"aws:SourceIp": "8.8.8.8/32"}
      }
    }
  ]
}
POLICY
}

```

```terraform
resource "aws_s3_bucket" "negative1" {
  bucket = "example"
}

resource "aws_s3_bucket_public_access_block" "negative2" {
  bucket = aws_s3_bucket.example.id

  block_public_acls   = true
  block_public_policy = true
  ignore_public_acls  = false
}
```