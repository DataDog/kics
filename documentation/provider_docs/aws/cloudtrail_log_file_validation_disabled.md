---
title: "CloudTrail Log File Validation Disabled"
meta:
  name: "aws/cloudtrail_log_file_validation_disabled"
  id: "52ffcfa6-6c70-4ea6-8376-d828d3961669"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Observability"
---

## Metadata
**Name:** `aws/cloudtrail_log_file_validation_disabled`

**Id:** `52ffcfa6-6c70-4ea6-8376-d828d3961669`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Observability

## Description
CloudTrail log file validation should be enabled to determine whether a log file has not been tampered

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudtrail#enable_log_file_validation)

## Non-Compliant Code Examples
```terraform
resource "aws_cloudtrail" "positive1" {
  name                          = "positive1"
  s3_bucket_name                = "bucketlog1"
}

resource "aws_cloudtrail" "positive2" {
  name                          = "positive2"
  s3_bucket_name                = "bucketlog2"
  enable_log_file_validation    = false
}

```

## Compliant Code Examples
```terraform
resource "aws_cloudtrail" "negative1" {
  name                          = "negative1"
  s3_bucket_name                = "bucketlog1"
  enable_log_file_validation    = true
}

```