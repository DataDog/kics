---
title: "MSK Cluster Logging Disabled"
meta:
  name: "terraform/msk_cluster_logging_disabled"
  id: "2f56b7ab-7fba-4e93-82f0-247e5ddeb239"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Observability"
---
## Metadata
**Name:** `terraform/msk_cluster_logging_disabled`
**Id:** `2f56b7ab-7fba-4e93-82f0-247e5ddeb239`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Observability
## Description
Amazon MSK cluster broker logging should be enabled to capture important audit and troubleshooting information. If the `logging_info.broker_logs.cloudwatch_logs.enabled` attribute is set to `false` or omitted, as shown below, critical activity and error logs from the MSK brokers will not be captured: 

```
resource "aws_msk_cluster" "example" {
  logging_info {
    broker_logs {
      cloudwatch_logs {
        enabled = false
      }
    }
  }
}
```

Without logging enabled, security events, operational issues, and access patterns will be invisible, making detection and investigation of incidents significantly more difficult. To ensure proper visibility, logging should be enabled as below:

```
resource "aws_msk_cluster" "example" {
  logging_info {
    broker_logs {
      cloudwatch_logs {
        enabled   = true
        log_group = aws_cloudwatch_log_group.example.name
      }
    }
  }
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/msk_cluster#broker_logs)

## Non-Compliant Code Examples
```aws
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
```aws
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