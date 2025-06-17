---
title: "S3 Bucket Without Versioning"
meta:
  name: "terraform/s3_bucket_without_versioning"
  id: "568a4d22-3517-44a6-a7ad-6a7eed88722c"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Backup"
---
## Metadata
**Name:** `terraform/s3_bucket_without_versioning`
**Id:** `568a4d22-3517-44a6-a7ad-6a7eed88722c`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Backup
## Description
Amazon S3 buckets should have versioning enabled to protect against accidental or malicious deletion and overwriting of objects. Without versioning (`versioning { enabled = false }`), deleted or overwritten files cannot be recovered, potentially leading to permanent data loss or loss of critical information. Enabling versioning (`versioning { enabled = true }`) allows you to preserve, retrieve, and restore every version of every object stored in the bucket, significantly improving data resiliency.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket#versioning)

## Non-Compliant Code Examples
```aws
terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "4.2.0"
    }
  }
}

provider "aws" {
  # Configuration options
}

resource "aws_s3_bucket" "b0" {
  bucket = "my-tf-test-bucket"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}

resource "aws_s3_bucket_versioning" "example" {
  bucket = aws_s3_bucket.b0.id

  versioning_configuration {
    status = "Suspended"
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
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

resource "aws_s3_bucket" "positive1" {
  bucket = "my-tf-test-bucket"
  acl    = "private"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }

  versioning {
    enabled = false
  }
}

```

```aws
module "s3_bucket" {
  source = "terraform-aws-modules/s3-bucket/aws"

  version = "3.7.0"

  bucket = "my-s3-bucket"
  acl    = "private"

  versioning = {
    enabled = false
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

  versioning {
    enabled = true
  }
}

```

```aws
terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "4.2.0"
    }
  }
}

provider "aws" {
  # Configuration options
}

resource "aws_s3_bucket" "b" {
  bucket = "my-tf-test-bucket"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}

resource "aws_s3_bucket_versioning" "example" {
  bucket = aws_s3_bucket.b.id

  versioning_configuration {
    status = "Enabled"
  }
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