---
title: "CloudWatch Logging Disabled"
meta:
  name: "terraform/cloudwatch_logging_disabled"
  id: "7dbba512-e244-42dc-98bb-422339827967"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Observability"
---
## Metadata
**Name:** `terraform/cloudwatch_logging_disabled`
**Id:** `7dbba512-e244-42dc-98bb-422339827967`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Observability
## Description
This check verifies whether Amazon Route53 hosted zones have CloudWatch logging enabled by ensuring an `aws_route53_query_log` resource is associated with the hosted zone and correctly configured. Without query logging, DNS queries to the hosted zone are not captured, making it difficult to detect potential misuse, troubleshoot DNS issues, or comply with audit requirements. Enable logging using the `aws_route53_query_log` resource, and specify the `cloudwatch_log_group_arn` and `zone_id` attributes, as in:

```
resource "aws_route53_query_log" "example_com" {
  cloudwatch_log_group_arn = aws_cloudwatch_log_group.aws_route53_example_com.arn
  zone_id                  = aws_route53_zone.example_com.zone_id
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_query_log)

## Non-Compliant Code Examples
```aws
resource "aws_route53_zone" "no_query_log" {
  name = "example.com"
}

resource "aws_route53_zone" "log_group_mismatch" {
  name = "example.com"
}

resource "aws_route53_query_log" "log_group_mismatch" {
  cloudwatch_log_group_arn = aws_cloudwatch_log_group.aws_route53_log_mismatch.arn
}
```

## Compliant Code Examples
```aws
resource "aws_route53_zone" "example_com" {
  name = "example.com"
}

resource "aws_route53_query_log" "example_com" {
  depends_on = [aws_cloudwatch_log_resource_policy.route53-query-logging-policy]

  cloudwatch_log_group_arn = aws_cloudwatch_log_group.aws_route53_example_com.arn
  zone_id                  = aws_route53_zone.example_com.zone_id
}
```