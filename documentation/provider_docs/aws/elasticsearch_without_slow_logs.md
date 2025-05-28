---
title: "ElasticSearch Without Slow Logs"
meta:
  name: "aws/elasticsearch_without_slow_logs"
  id: "e979fcbc-df6c-422d-9458-c33d65e71c45"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Observability"
---

## Metadata
**Name:** `aws/elasticsearch_without_slow_logs`

**Id:** `e979fcbc-df6c-422d-9458-c33d65e71c45`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Observability

## Description
Ensure that AWS Elasticsearch enables support for slow logs

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elasticsearch_domain#log_publishing_options)

## Non-Compliant Code Examples
```terraform
resource "aws_elasticsearch_domain" "positive1" {
  log_publishing_options {
    cloudwatch_log_group_arn = aws_cloudwatch_log_group.example.arn
    log_type                 = "ES_APPLICATION_LOGS"
    enabled                  = true
  }
}

```

## Compliant Code Examples
```terraform
resource "aws_elasticsearch_domain" "negative1" {

  log_publishing_options {
    cloudwatch_log_group_arn = aws_cloudwatch_log_group.example.arn
    log_type                 = "INDEX_SLOW_LOGS"
    enabled                  = true //for default its true
  }
}

```