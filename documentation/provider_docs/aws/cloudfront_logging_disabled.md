---
title: "CloudFront Logging Disabled"
meta:
  name: "aws/cloudfront_logging_disabled"
  id: "94690d79-b3b0-43de-b656-84ebef5753e5"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Observability"
---

## Metadata
**Name:** `aws/cloudfront_logging_disabled`

**Id:** `94690d79-b3b0-43de-b656-84ebef5753e5`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Observability

## Description
AWS CloudFront distributions should have logging enabled to collect all viewer requests, which means the attribute 'logging_config' should be defined

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudfront_distribution)

## Non-Compliant Code Examples
```terraform
resource "aws_cloudfront_distribution" "positive1" {
  origin {
    domain_name = aws_s3_bucket.b.bucket_regional_domain_name
    origin_id   = local.s3_origin_id

    s3_origin_config {
      origin_access_identity = "origin-access-identity/cloudfront/ABCDEFG1234567"
    }
  }

  enabled             = true
  is_ipv6_enabled     = true
  comment             = "Some comment"
  default_root_object = "index.html"

  }
  
```

## Compliant Code Examples
```terraform
resource "aws_cloudfront_distribution" "negative1" {
  origin {
    domain_name = aws_s3_bucket.b.bucket_regional_domain_name
    origin_id   = local.s3_origin_id

    s3_origin_config {
      origin_access_identity = "origin-access-identity/cloudfront/ABCDEFG1234567"
    }
  }

  enabled             = true
  is_ipv6_enabled     = true
  comment             = "Some comment"
  default_root_object = "index.html"

  logging_config {
    include_cookies = false
    bucket          = "mylogs.s3.amazonaws.com"
    prefix          = "myprefix"
  }
  
}
```