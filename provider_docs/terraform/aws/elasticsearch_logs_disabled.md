---
title: "Elasticsearch Log Disabled"
meta:
  name: "terraform/elasticsearch_logs_disabled"
  id: "acb6b4e2-a086-4f35-aefd-4db6ea51ada2"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Observability"
---
## Metadata
**Name:** `terraform/elasticsearch_logs_disabled`
**Id:** `acb6b4e2-a086-4f35-aefd-4db6ea51ada2`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Observability
## Description
Elasticsearch domains in AWS should have logging enabled to capture important audit and performance data. The attribute `enabled` within the `log_publishing_options` block should be set to `true` to ensure that logs, such as `"INDEX_SLOW_LOGS"`, are published to the associated CloudWatch log group. If log publishing is disabled (e.g. `enabled = false`), critical operational and security events may go undetected, making it difficult to troubleshoot issues, monitor for suspicious activity, or meet compliance requirements. Failing to enable logging increases the risk of undetected data breaches or operational failures due to limited visibility into Elasticsearch domain activity.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elasticsearch_domain#log_publishing_options)

## Non-Compliant Code Examples
```aws
resource "aws_elasticsearch_domain" "positive1" {

  log_publishing_options {
    cloudwatch_log_group_arn = aws_cloudwatch_log_group.example.arn
    log_type                 = "INDEX_SLOW_LOGS"
    enabled                  = false
  }
}

```

```aws
resource "aws_elasticsearch_domain" "positive2" {
  domain_name           = "example"
  elasticsearch_version = "1.5"

  cluster_config {
    instance_type = "r4.large.elasticsearch"
  }

  snapshot_options {
    automated_snapshot_start_hour = 23
  }

  tags = {
    Domain = "TestDomain"
  }
}

```

## Compliant Code Examples
```aws
resource "aws_elasticsearch_domain" "negative1" {

  log_publishing_options {
    cloudwatch_log_group_arn = aws_cloudwatch_log_group.example.arn
    log_type                 = "INDEX_SLOW_LOGS"
    enabled                  = true //for default its true
  }
}

```