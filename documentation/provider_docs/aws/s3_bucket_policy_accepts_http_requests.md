---
title: "S3 Bucket Policy Accepts HTTP Requests"
meta:
  name: "aws/s3_bucket_policy_accepts_http_requests"
  id: "4bc4dd4c-7d8d-405e-a0fb-57fa4c31b4d9"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Encryption"
---

## Metadata
**Name:** `aws/s3_bucket_policy_accepts_http_requests`

**Id:** `4bc4dd4c-7d8d-405e-a0fb-57fa4c31b4d9`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Encryption

## Description
S3 Bucket policy should not accept HTTP Requests

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_policy#policy)

## Non-Compliant Code Examples
```terraform
resource "aws_s3_bucket" "b2" {
  bucket = "my-tf-test-bucket"

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Id": "MYBUCKETPOLICY",
    "Statement": [
      {
        "Sid": "IPAllow",
        "Effect": "Deny",
        "Principal": "*",
        "Action": "s3:*",
        "Resource": [
          "aws_s3_bucket.b.arn"
        ],
        "Condition": {
          "IpAddress": {
            "aws:SourceIp": "8.8.8.8/32"
          }
        }
      }
    ]
}
EOF
}

```

```terraform
module "s3_bucket" {
  source = "terraform-aws-modules/s3-bucket/aws"
  version = "3.7.0"

  bucket = "my-s3-bucket"
  acl    = "private"

  versioning = {
    enabled = true
  }

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Id": "MYBUCKETPOLICY",
    "Statement": [
      {
        "Sid": "IPAllow",
        "Effect": "Deny",
        "Principal": "*",
        "Action": "s3:*",
        "Resource": [
          "aws_s3_bucket.b.arn"
        ],
        "Condition": {
          "IpAddress": {
            "aws:SourceIp": "8.8.8.8/32"
          }
        }
      }
    ]
}
EOF
}

```

```terraform

data "aws_iam_policy_document" "pos4" {

  statement {
    effect = "Deny"

    principals {
      type        = "*"
      identifiers = ["*"]
    }

    actions = [
      "s3:*",
    ]


    resources = [
      "arn:aws:s3:::a/*",
      "arn:aws:s3:::a",
    ]
    condition {
      test     = "Bool"
      variable = "aws:SecureTransport"
      values   = ["true"]
    }
  }
}


resource "aws_s3_bucket" "pos4" {
  bucket = "a"
  policy = data.aws_iam_policy_document.pos4.json
}

```

```terraform

data "aws_iam_policy_document" "pos5" {

  statement {
    effect = "Allow"

    principals {
      type        = "*"
      identifiers = ["*"]
    }

    actions = [
      "s3:*",
    ]


    resources = [
      "arn:aws:s3:::a/*",
      "arn:aws:s3:::a",
    ]
    condition {
      test     = "Bool"
      variable = "aws:SecureTransport"
      values   = ["false"]
    }
  }
}


resource "aws_s3_bucket" "pos5" {
  bucket = "a"
  policy = data.aws_iam_policy_document.pos5.json
}

```

```terraform
resource "aws_s3_bucket" "b" {
  bucket = "my-tf-test-bucket"
}

resource "aws_s3_bucket_policy" "b" {
  bucket = aws_s3_bucket.b.id

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Id": "MYBUCKETPOLICY",
    "Statement": [
      {
        "Sid": "IPAllow",
        "Effect": "Deny",
        "Principal": "*",
        "Action": "s3:*",
        "Resource": [
          "aws_s3_bucket.b.arn"
        ],
        "Condition": {
          "IpAddress": {
            "aws:SourceIp": "8.8.8.8/32"
          }
        }
      }
    ]
}
EOF
}





```

## Compliant Code Examples
```terraform
resource "aws_s3_bucket" "negative7" {
  bucket = "my-tf-test-bucket"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}

data "aws_iam_policy_document" "policy" {
  statement {
    sid    = "https"
    effect = "Deny"
    principals {
      type        = "*"
      identifiers = ["*"]
    }
    actions = [
      "s3:*"
    ]
    resources = [
      aws_s3_bucket.negative7.arn,
      "${aws_s3_bucket.negative7.arn}/*"
    ]
    condition {
      test     = "Bool"
      variable = "aws:SecureTransport"
      values = [
        "false"
      ]
    }
  }
}


resource "aws_s3_bucket_policy" "bucket_policy" {
  bucket = aws_s3_bucket.negative7.id
  policy = data.aws_iam_policy_document.policy.json
}

```

```terraform
module "s3_bucket" {
  source = "terraform-aws-modules/s3-bucket/aws"
  version = "3.7.0"

  bucket = "my-s3-bucket"
  acl    = "private"

  versioning = {
    enabled = true
  }

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Id": "MYBUCKETPOLICY",
    "Statement": [
      {
        "Sid": "IPAllow",
        "Effect": "Deny",
        "Principal": "*",
        "Action": "s3:*",
        "Resource": [
          "aws_s3_bucket.b.arn"
        ],
        "Condition": {
          "Bool": {
            "aws:SecureTransport": "false"
          }
        }
      },
      {
        "Sid": "IPAllow",
        "Effect": "Deny",
        "Principal": "*",
        "Action": "s3:*",
        "Resource": [
          "aws_s3_bucket.c.arn"
        ],
        "Condition": {
          "Bool": {
            "aws:SecureTransport": "false"
          }
        }
      }
    ]
}
EOF
}

```

```terraform
resource "aws_s3_bucket" "b2" {
  bucket = "my-tf-test-bucket"

   policy = <<EOF
{
    "Version": "2012-10-17",
    "Id": "MYBUCKETPOLICY",
    "Statement": [
      {
        "Sid": "IPAllow",
        "Effect": "Deny",
        "Principal": "*",
        "Action": "s3:*",
        "Resource": [
          "aws_s3_bucket.b.arn"
        ],
        "Condition": {
          "Bool": {
            "aws:SecureTransport": "false"
          }
        }
      }
    ]
}
EOF
}

```

```terraform

data "aws_iam_policy_document" "neg6" {

  statement {
    effect = "Allow"

    principals {
      type        = "*"
      identifiers = ["*"]
    }

    actions = [
      "s3:*",
    ]


    resources = [
      "arn:aws:s3:::a/*",
      "arn:aws:s3:::a",
    ]
    condition {
      test     = "Bool"
      variable = "aws:SecureTransport"
      values   = ["true"]
    }
  }
}


resource "aws_s3_bucket" "neg6" {
  bucket = "a"
  policy = data.aws_iam_policy_document.neg6.json
}

```

```terraform
resource "aws_s3_bucket" "b" {
  bucket = "my-tf-test-bucket"
}

resource "aws_s3_bucket_policy" "b" {
  bucket = aws_s3_bucket.b.id

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Id": "MYBUCKETPOLICY",
    "Statement": [
      {
        "Sid": "IPAllow",
        "Effect": "Deny",
        "Principal": "*",
        "Action": "s3:*",
        "Resource": [
          "aws_s3_bucket.b.arn"
        ],
        "Condition": {
          "Bool": {
            "aws:SecureTransport": "false"
          }
        }
      }
    ]
}
EOF
}





```

```terraform

data "aws_iam_policy_document" "neg5" {

  statement {
    effect = "Deny"

    principals {
      type        = "*"
      identifiers = ["*"]
    }

    actions = [
      "s3:*",
    ]


    resources = [
      "arn:aws:s3:::a/*",
      "arn:aws:s3:::a",
    ]
    condition {
      test     = "Bool"
      variable = "aws:SecureTransport"
      values   = ["false"]
    }
  }
}


resource "aws_s3_bucket" "neg5" {
  bucket = "a"
  policy = data.aws_iam_policy_document.neg5.json
}

```

```terraform
module "s3_bucket" {
  source = "terraform-aws-modules/s3-bucket/aws"
  version = "3.7.0"

  bucket = "my-s3-bucket"
  acl    = "private"

  versioning = {
    enabled = true
  }

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Id": "MYBUCKETPOLICY",
    "Statement": [
      {
        "Sid": "IPAllow",
        "Effect": "Allow",
        "Principal": "*",
        "Action": "s3:*",
        "Resource": [
          "aws_s3_bucket.b.arn"
        ],
        "Condition": {
          "Bool": {
            "aws:SecureTransport": "true"
          }
        }
      }
    ]
}
EOF
}

```