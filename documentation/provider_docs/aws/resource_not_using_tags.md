---
title: "Resource Not Using Tags"
meta:
  name: "aws/resource_not_using_tags"
  id: "e38a8e0a-b88b-4902-b3fe-b0fcb17d5c10"
  cloud_provider: "aws"
  severity: "INFO"
  category: "Best Practices"
---

## Metadata
**Name:** `aws/resource_not_using_tags`

**Id:** `e38a8e0a-b88b-4902-b3fe-b0fcb17d5c10`

**Cloud Provider:** aws

**Severity:** Info

**Category:** Best Practices

## Description
AWS services resource tags are an essential part of managing components. As a best practice, the field 'tags' should have additional tags defined other than 'Name'

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/guides/resource-tagging)

## Non-Compliant Code Examples
```terraform
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
```terraform
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