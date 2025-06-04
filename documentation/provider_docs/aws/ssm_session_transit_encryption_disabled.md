---
title: "SSM Session Transit Encryption Disabled"
meta:
  name: "aws/ssm_session_transit_encryption_disabled"
  id: "ce60cc6b-6831-4bd7-84a2-cc7f8ee71433"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Encryption"
---

## Metadata
**Name:** `aws/ssm_session_transit_encryption_disabled`

**Id:** `ce60cc6b-6831-4bd7-84a2-cc7f8ee71433`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Encryption

## Description
SSM Session should be encrypted in transit

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ssm_document#content)

## Non-Compliant Code Examples
```terraform
resource "aws_ssm_document" "positive2" {
  name          = "test_document"
  document_type = "Session"

  content = <<DOC
  {
    "schemaVersion": "1.2",
    "description": "Check ip configuration of a Linux instance.",
    "inputs": {
      "s3EncryptionEnabled": true,
      "cloudWatchEncryptionEnabled": true,
      "cloudWatchStreamingEnabled": true,
      "runAsEnabled": false
    }
  }
DOC
}

```

```terraform
resource "aws_ssm_document" "positive1" {
  name          = "test_document"
  document_type = "Session"

  content = <<DOC
  {
    "schemaVersion": "1.2",
    "description": "Check ip configuration of a Linux instance."
  }
DOC
}

```

## Compliant Code Examples
```terraform
resource "aws_ssm_document" "negative" {
  name          = "test_document"
  document_type = "Session"

  content = <<DOC
  {
    "schemaVersion": "1.2",
    "description": "Check ip configuration of a Linux instance.",
    "inputs": {
      "s3EncryptionEnabled": true,
      "cloudWatchEncryptionEnabled": true,
      "cloudWatchStreamingEnabled": true,
      "runAsEnabled": false,
      "kmsKeyId": "${var.kms_key_id}"
    }
  }
DOC
}

```