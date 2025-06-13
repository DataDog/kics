---
title: "API Gateway Without Security Policy"
meta:
  name: "aws/api_gateway_without_security_policy"
  id: "4e1cc5d3-2811-4fb2-861c-ee9b3cb7f90b"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `aws/api_gateway_without_security_policy`

**Id:** `4e1cc5d3-2811-4fb2-861c-ee9b3cb7f90b`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Insecure Configurations

## Description
API Gateway should have a Security Policy defined and use TLS 1.2.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_domain_name#security_policy)

## Non-Compliant Code Examples
```terraform
resource "aws_api_gateway_domain_name" "example2" {
  domain_name              = "api.example.com"
  security_policy = "TLS_1_0"
}

```

```terraform
resource "aws_api_gateway_domain_name" "example" {
  domain_name              = "api.example.com"
}

```

## Compliant Code Examples
```terraform
resource "aws_api_gateway_domain_name" "example4" {
  domain_name              = "api.example.com"
  security_policy = "TLS_1_2"
}

```