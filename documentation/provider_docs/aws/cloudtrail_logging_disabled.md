---
title: "CloudTrail Logging Disabled"
meta:
  name: "aws/cloudtrail_logging_disabled"
  id: "4bb76f17-3d63-4529-bdca-2b454529d774"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Observability"
---

## Metadata
**Name:** `aws/cloudtrail_logging_disabled`

**Id:** `4bb76f17-3d63-4529-bdca-2b454529d774`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Observability

## Description
Checks if logging is enabled for CloudTrail.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudtrail#enable_logging)

## Non-Compliant Code Examples
```terraform
#this is a problematic code where the query should report a result(s)
resource "aws_cloudtrail" "positive1" {
  name                          = "positive"
  s3_bucket_name                = "bucketlog"
  enable_logging                = false
}
```

## Compliant Code Examples
```terraform
resource "aws_cloudtrail" "negative1" {
  name                          = "negative_1"
  s3_bucket_name                = "bucketlog"
  enable_logging                = true
}

resource "aws_cloudtrail" "negative2" {
  name                          = "negative_2"
  s3_bucket_name                = "bucketlog"
}
```