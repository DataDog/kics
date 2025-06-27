---
title: "API Gateway Without Security Policy"
meta:
  name: "aws/api_gateway_without_security_policy"
  id: "4e1cc5d3-2811-4fb2-861c-ee9b3cb7f90b"
  display_name: "API Gateway Without Security Policy"
  cloud_provider: "aws"
  platform: "Terraform"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `aws/api_gateway_without_security_policy`
**Query Name** `API Gateway Without Security Policy`
**Id:** `4e1cc5d3-2811-4fb2-861c-ee9b3cb7f90b`
**Cloud Provider:** aws
**Platform** Terraform
**Severity:** Medium
**Category:** Insecure Configurations
## Description
The AWS API Gateway custom domain resource should have a security policy explicitly defined to enforce the use of strong encryption protocols. By omitting the `security_policy` attribute or leaving it unset, as shown below:

```
resource "aws_api_gateway_domain_name" "example" {
  domain_name = "api.example.com"
}
```

the domain name may default to an older, less secure version of TLS, making the API vulnerable to downgrade attacks and exposure of sensitive data. Setting `security_policy = "TLS_1_2"` ensures that only connections using TLS 1.2 are allowed, significantly increasing the security posture of the API endpoint:

```
resource "aws_api_gateway_domain_name" "example" {
  domain_name     = "api.example.com"
  security_policy = "TLS_1_2"
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_domain_name#security_policy)


## Compliant Code Examples
```terraform
resource "aws_api_gateway_domain_name" "example4" {
  domain_name              = "api.example.com"
  security_policy = "TLS_1_2"
}

```
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