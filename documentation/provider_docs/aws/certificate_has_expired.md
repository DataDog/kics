---
title: "Certificate Has Expired"
meta:
  name: "aws/certificate_has_expired"
  id: "c3831315-5ae6-4fa8-b458-3d4d5ab7a3f6"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `aws/certificate_has_expired`

**Id:** `c3831315-5ae6-4fa8-b458-3d4d5ab7a3f6`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Access Control

## Description
Expired SSL/TLS certificates should be removed

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_rest_api)

## Non-Compliant Code Examples
```terraform
resource "aws_api_gateway_domain_name" "example2" {
  certificate_body = file("expiredCertificate.pem")
  domain_name     = "api.example.com"
}


```

## Compliant Code Examples
```terraform
resource "aws_api_gateway_domain_name" "example" {
  certificate_body = file("validCertificate.pem")
  domain_name     = "api.example.com"
}


```