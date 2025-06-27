---
title: "Glue Security Configuration Encryption Disabled"
meta:
  name: "aws/glue_security_configuration_encryption_disabled"
  id: "ad5b4e97-2850-4adf-be17-1d293e0b85ee"
  display_name: "Glue Security Configuration Encryption Disabled"
  cloud_provider: "aws"
  platform: "Terraform"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata

**Name:** `aws/glue_security_configuration_encryption_disabled`

**Query Name** `Glue Security Configuration Encryption Disabled`

**Id:** `ad5b4e97-2850-4adf-be17-1d293e0b85ee`

**Cloud Provider:** aws

**Platform** Terraform

**Severity:** High

**Category:** Encryption

## Description
AWS Glue Security Configuration requires proper encryption settings for all three components (CloudWatch, job bookmarks, and S3) with valid KMS key ARNs to ensure comprehensive data protection. When any of these components lacks proper encryption configuration or is missing the required KMS key ARN, it creates security vulnerabilities that could expose sensitive data. The impact of this misconfiguration includes potential unauthorized access to data, compliance violations, and increased risk of data breaches.

Secure configuration example:
```terraform
encryption_configuration {
  cloudwatch_encryption {
    cloudwatch_encryption_mode = "SSE-KMS"
    kms_key_arn = data.aws_kms_key.example.arn
  }
  job_bookmarks_encryption {
    job_bookmarks_encryption_mode = "CSE-KMS"
    kms_key_arn = data.aws_kms_key.example.arn
  }
  s3_encryption {
    s3_encryption_mode = "SSE-KMS"
    kms_key_arn = data.aws_kms_key.example.arn
  }
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/glue_security_configuration#encryption_configuration)


## Compliant Code Examples
```terraform
resource "aws_glue_security_configuration" "negative1" {
  name = "example"

  encryption_configuration {
    cloudwatch_encryption {
      cloudwatch_encryption_mode = "SSE-KMS"
      kms_key_arn = data.aws_kms_key.example.arn
    }

    job_bookmarks_encryption {
      job_bookmarks_encryption_mode = "CSE-KMS"
      kms_key_arn = data.aws_kms_key.example.arn
    }

    s3_encryption {
      kms_key_arn        = data.aws_kms_key.example.arn
      s3_encryption_mode = "SSE-KMS"
    }
  }
}

```
## Non-Compliant Code Examples
```terraform
resource "aws_glue_security_configuration" "positive2" {
  name = "example"

  encryption_configuration {
    cloudwatch_encryption {
      cloudwatch_encryption_mode = "SSE-KMS"
      kms_key_arn = data.aws_kms_key.example.arn
    }

    job_bookmarks_encryption {
      job_bookmarks_encryption_mode = "DISABLED"
      kms_key_arn = data.aws_kms_key.example.arn
    }

    s3_encryption {
      kms_key_arn        = data.aws_kms_key.example.arn
      s3_encryption_mode = "SSE-KMS"
    }
  }
}

```

```terraform
resource "aws_glue_security_configuration" "positive2" {
  name = "example"

  encryption_configuration {
    cloudwatch_encryption {
      cloudwatch_encryption_mode = "SSE-KMS"
      kms_key_arn = data.aws_kms_key.example.arn
    }

    job_bookmarks_encryption {
      kms_key_arn = data.aws_kms_key.example.arn
    }

    s3_encryption {
      kms_key_arn        = data.aws_kms_key.example.arn
      s3_encryption_mode = "SSE-KMS"
    }
  }
}

```

```terraform
resource "aws_glue_security_configuration" "positive1" {
  name = "example"

  encryption_configuration {
    cloudwatch_encryption {
      cloudwatch_encryption_mode = "SSE-KMS"
    }

    job_bookmarks_encryption {
      job_bookmarks_encryption_mode = "CSE-KMS"
      kms_key_arn = data.aws_kms_key.example.arn
    }

    s3_encryption {
      kms_key_arn        = data.aws_kms_key.example.arn
      s3_encryption_mode = "SSE-KMS"
    }
  }
}

```