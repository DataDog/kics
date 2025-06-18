---
title: "S3 Bucket ACL Allows Read to Any Authenticated User"
meta:
  name: "terraform/s3_bucket_acl_allows_read_to_any_authenticated_user"
  id: "57b9893d-33b1-4419-bcea-a717ea87e139"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/s3_bucket_acl_allows_read_to_any_authenticated_user`
**Id:** `57b9893d-33b1-4419-bcea-a717ea87e139`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Access Control
## Description
This check identifies S3 buckets that use the 'authenticated-read' ACL, which grants read access to any authenticated AWS user, regardless of their account. This configuration creates a significant security risk as it exposes your data to all authenticated AWS users worldwide, potentially leading to unauthorized data access and information disclosure. 

Secure configuration example:
```
resource "aws_s3_bucket" "example" {
  bucket = "my-tf-test-bucket"
  acl    = "private"
}
```

Insecure configuration example:
```
resource "aws_s3_bucket" "example" {
  bucket = "my-tf-test-bucket"
  acl    = "authenticated-read"
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket#acl)

## Non-Compliant Code Examples
```aws
provider "aws" {
  region = "us-east-1"
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

resource "aws_s3_bucket" "positive1" {
  bucket = "my-tf-test-bucket"
  acl    = "authenticated-read"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}

```

```aws
provider "aws" {
  region = "us-east-1"
}

terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "4.2.0"
    }
  }
}

resource "aws_s3_bucket" "example1" {
  bucket = "my-tf-example-bucket"
}

resource "aws_s3_bucket_acl" "example_bucket_acl" {
  bucket = aws_s3_bucket.example1.id
  acl    = "authenticated-read"
}

```

```aws
module "s3_bucket" {
  source = "terraform-aws-modules/s3-bucket/aws"
  version = "3.7.0"

  bucket = "my-s3-bucket"
  acl    = "authenticated-read"

  versioning = {
    enabled = true
  }
}

```

## Compliant Code Examples
```aws
provider "aws" {
  region = "us-east-1"
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

resource "aws_s3_bucket" "negative1" {
  bucket = "my-tf-test-bucket"
  acl    = "private"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}

```

```aws
provider "aws" {
  region = "us-east-1"
}

terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "4.2.0"
    }
  }
}

resource "aws_s3_bucket" "example0" {
  bucket = "my-tf-example-bucket"
}

resource "aws_s3_bucket_acl" "example_bucket_acl" {
  bucket = aws_s3_bucket.example0.id
  acl    = "private"
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