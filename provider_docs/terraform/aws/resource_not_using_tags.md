---
title: "Resource Not Using Tags"
meta:
  name: "terraform/resource_not_using_tags"
  id: "e38a8e0a-b88b-4902-b3fe-b0fcb17d5c10"
  cloud_provider: "terraform"
  severity: "INFO"
  category: "Best Practices"
---
## Metadata
**Name:** `terraform/resource_not_using_tags`
**Id:** `e38a8e0a-b88b-4902-b3fe-b0fcb17d5c10`
**Cloud Provider:** terraform
**Severity:** Info
**Category:** Best Practices
## Description
AWS resource tagging is an essential best practice that supports resource management, cost tracking, automation, and security enforcement. If only the default `Name` tag is applied and no additional metadata tags (such as `Environment`, `Owner`, or `Project`) are defined, resources may become difficult to identify, audit, and manage at scale, increasing the risk of misconfiguration or unintended resource exposure. For example, a secure configuration should use a `tags` block with descriptive tags:

```
tags = {
  Name        = "example-cert"
  Environment = "production"
  Owner       = "devops-team"
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/guides/resource-tagging)

## Non-Compliant Code Examples
```aws
resource "aws_acm_certificate" "cert" {
  domain_name       = "example.com"
  validation_method = "DNS"

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_acm_certificate" "cert_2" {
  domain_name       = "example.com"
  validation_method = "DNS"

  tags = {
    Name = "test"
  }

  lifecycle {
    create_before_destroy = true
  }
}

```

## Compliant Code Examples
```aws
resource "aws_acm_certificate" "cert" {
  domain_name       = "example.com"
  validation_method = "DNS"

  tags = {
    Environment = "test"
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_acm_certificate_validation" "example" {
  certificate_arn         = aws_acm_certificate.example.arn
  validation_record_fqdns = [for record in aws_route53_record.example : record.fqdn]
}

```