---
title: "Glue Data Catalog Encryption Disabled"
meta:
  name: "terraform/glue_data_catalog_encryption_disabled"
  id: "01d50b14-e933-4c99-b314-6d08cd37ad35"
  cloud_provider: "terraform"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata
**Name:** `terraform/glue_data_catalog_encryption_disabled`
**Id:** `01d50b14-e933-4c99-b314-6d08cd37ad35`
**Cloud Provider:** terraform
**Severity:** High
**Category:** Encryption
## Description
AWS Glue Data Catalog contains metadata about AWS resources and should be properly encrypted to protect sensitive information. When encryption is disabled for connection passwords or data at rest, it could expose sensitive connection credentials and metadata to unauthorized access, potentially leading to data breaches or unauthorized resource access. Enabling both connection password encryption (with return_connection_password_encrypted set to true) and encryption at rest with SSE-KMS ensures that all sensitive metadata is properly protected with AWS KMS keys.

Example of secure configuration:
```
resource "aws_glue_data_catalog_encryption_settings" "secure_example" {
  data_catalog_encryption_settings {
    connection_password_encryption {
      aws_kms_key_id                       = aws_kms_key.test.arn
      return_connection_password_encrypted = true
    }

    encryption_at_rest {
      catalog_encryption_mode = "SSE-KMS"
      sse_aws_kms_key_id      = aws_kms_key.test.arn
    }
  }
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/glue_data_catalog_encryption_settings#data_catalog_encryption_settings)

## Non-Compliant Code Examples
```aws
resource "aws_glue_data_catalog_encryption_settings" "positive1" {
  data_catalog_encryption_settings {
    connection_password_encryption {
      aws_kms_key_id                       = aws_kms_key.test.arn
      return_connection_password_encrypted = false
    }

    encryption_at_rest {
      catalog_encryption_mode = "SSE-KMS"
      sse_aws_kms_key_id      = aws_kms_key.test.arn
    }
  }
}

```

```aws
resource "aws_glue_data_catalog_encryption_settings" "positive4" {
  data_catalog_encryption_settings {
    connection_password_encryption {
      aws_kms_key_id                       = aws_kms_key.test.arn
      return_connection_password_encrypted = true
    }

    encryption_at_rest {
      catalog_encryption_mode = "SSE-KMS"
    }
  }
}

```

```aws
resource "aws_glue_data_catalog_encryption_settings" "positive3" {
  data_catalog_encryption_settings {
    connection_password_encryption {
      aws_kms_key_id                       = aws_kms_key.test.arn
      return_connection_password_encrypted = true
    }

    encryption_at_rest {
      catalog_encryption_mode = "DISABLED"
      sse_aws_kms_key_id      = aws_kms_key.test.arn
    }
  }
}

```

## Compliant Code Examples
```aws
resource "aws_glue_data_catalog_encryption_settings" "negative1" {
  data_catalog_encryption_settings {
    connection_password_encryption {
      aws_kms_key_id                       = aws_kms_key.test.arn
      return_connection_password_encrypted = true
    }

    encryption_at_rest {
      catalog_encryption_mode = "SSE-KMS"
      sse_aws_kms_key_id      = aws_kms_key.test.arn
    }
  }
}

```