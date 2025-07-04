---
title: "Certificate Has Expired"
group-id: "rules/terraform/aws"
meta:
  name: "aws/certificate_has_expired"
  id: "c3831315-5ae6-4fa8-b458-3d4d5ab7a3f6"
  display_name: "Certificate Has Expired"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata

**Id:** `c3831315-5ae6-4fa8-b458-3d4d5ab7a3f6`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Medium

**Category:** Access Control

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_rest_api)

### Description

 Expired SSL/TLS certificates should be removed from cloud resources to prevent the risk of exposing users to insecure or untrusted connections. When a resource, such as an AWS API Gateway custom domain, is configured with an expired certificate (e.g., `certificate_body = file("expiredCertificate.pem")`), clients attempting to access the API will receive security warnings, and automated clients may reject the connection entirely. This vulnerability undermines the integrity and trust of the service, potentially leading to denial of service, data interception, or man-in-the-middle attacks. Regularly updating and ensuring only valid certificates are used helps maintain secure encrypted communications between clients and services.


## Compliant Code Examples
```terraform
resource "aws_api_gateway_domain_name" "example" {
  certificate_body = file("validCertificate.pem")
  domain_name     = "api.example.com"
}


```
## Non-Compliant Code Examples
```terraform
resource "aws_api_gateway_domain_name" "example2" {
  certificate_body = file("expiredCertificate.pem")
  domain_name     = "api.example.com"
}


```