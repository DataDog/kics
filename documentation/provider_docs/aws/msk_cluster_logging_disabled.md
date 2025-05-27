---
title: "MSK Cluster Logging Disabled"
meta:
  name: "aws/msk_cluster_logging_disabled"
  id: "2f56b7ab-7fba-4e93-82f0-247e5ddeb239"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Observability"
---

## Metadata
**Name:** `aws/msk_cluster_logging_disabled`

**Id:** `2f56b7ab-7fba-4e93-82f0-247e5ddeb239`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Observability

## Description
Ensure MSK Cluster Logging is enabled

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/msk_cluster#broker_logs)

## Non-Compliant Code Examples
```terraform
resource "aws_msk_cluster" "positive1" {
  logging_info {
    broker_logs {
      cloudwatch_logs {
        enabled   = false
        log_group = aws_cloudwatch_log_group.test.name
      }
      firehose {
        delivery_stream = aws_kinesis_firehose_delivery_stream.test_stream.name
      }
    }
  }
}

resource "aws_msk_cluster" "positive2" {

}

```

## Compliant Code Examples
```terraform
resource "aws_msk_cluster" "negative1" {  
  logging_info {
    broker_logs {
      cloudwatch_logs {
        enabled   = true
        log_group = aws_cloudwatch_log_group.test.name
      }
    }
  }
}
```