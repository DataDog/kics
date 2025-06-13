---
title: "CloudTrail Multi Region Disabled"
meta:
  name: "aws/cloudtrail_multi_region_disabled"
  id: "8173d5eb-96b5-4aa6-a71b-ecfa153c123d"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Observability"
---

## Metadata
**Name:** `aws/cloudtrail_multi_region_disabled`

**Id:** `8173d5eb-96b5-4aa6-a71b-ecfa153c123d`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Observability

## Description
CloudTrail multi region should be enabled, which means attributes 'is_multi_region_trail' and 'include_global_service_events' should be enabled

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudtrail#is_multi_region_trail)

## Non-Compliant Code Examples
```terraform
resource "aws_cloudtrail" "positive2" {
  name                          = "npositive_2"
  s3_bucket_name                = "bucketlog_2"
  is_multi_region_trail         = false
}

```

```terraform
resource "aws_cloudtrail" "positive3" {
  name                          = "npositive_3"
  s3_bucket_name                = "bucketlog_3"
  is_multi_region_trail         = true
  include_global_service_events = false
}

```

```terraform
#this is a problematic code where the query should report a result(s)
resource "aws_cloudtrail" "positive1" {
  name                          = "npositive_1"
  s3_bucket_name                = "bucketlog_1"
}

```

## Compliant Code Examples
```terraform
#this code is a correct code for which the query should not find any result
resource "aws_cloudtrail" "negative1" {
  name                          = "negative"
  s3_bucket_name                = "bucketlog"
  is_multi_region_trail         = true
}

```