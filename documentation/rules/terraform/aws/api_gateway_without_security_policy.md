---
title: "API Gateway without security policy"
group_id: "rules/terraform/aws"
meta:
  name: "aws/api_gateway_without_security_policy"
  id: "4e1cc5d3-2811-4fb2-861c-ee9b3cb7f90b"
  display_name: "API Gateway without security policy"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `4e1cc5d3-2811-4fb2-861c-ee9b3cb7f90b`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Medium

**Category:** Insecure Configurations

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_domain_name#security_policy)

### Description

 The AWS API Gateway custom domain resource should have a security policy explicitly defined to enforce the use of strong encryption protocols. By omitting the `security_policy` attribute or leaving it unset, as shown below, the domain name may default to an older, less secure version of TLS, making the API vulnerable to downgrade attacks and exposure of sensitive data.

```
resource "aws_api_gateway_domain_name" "example" {
  domain_name = "api.example.com"
}
```

Setting `security_policy = "TLS_1_2"` ensures that only connections using TLS 1.2 are allowed, significantly increasing the security posture of the API endpoint:

```
resource "aws_api_gateway_domain_name" "example" {
  domain_name     = "api.example.com"
  security_policy = "TLS_1_2"
}
```


## Compliant Code Examples
```terraform
resource "aws_api_gateway_domain_name" "example4" {
  domain_name              = "api.example.com"
  security_policy = "TLS_1_2"
}

```

```terraform
module "api_gateway_custom_domain" {
  source   = "terraform-aws-modules/apigateway-v2/aws"
  version  = "3.0.2"
  security_policy = "TLS_1_2"
  description          = "My awesome HTTP API"
  create_api_domain_name = true
  domain_name = "api.${module.acm.acm_certificate_domain}"

  domain_name_certificate_arn = module.acm.acm_certificate_arn

  cors_configuration = {
    allow_credentials = false
    allow_headers     = ["date", "keep-alive"]
    allow_methods     = ["PUT", "POST"]
    allow_origins     = ["*"]
    expose_headers    = ["keep-alive"]
    max_age           = 5
  }

  integrations = {
    "ANY /{proxy+}" = {
      lambda_arn             = module.lambda_function.lambda_function_arn
      payload_format_version = "2.0"
      timeout_milliseconds  = 12000
      }
    }
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

```terraform
module "api_gateway_custom_domain" {
  source   = "terraform-aws-modules/apigateway-v2/aws"
  version  = "3.0.2"
  security_policy = "TLS_1_0"
  description          = "My awesome HTTP API"
  create_api_domain_name = true
  domain_name = "api.${module.acm.acm_certificate_domain}"

  domain_name_certificate_arn = module.acm.acm_certificate_arn

  cors_configuration = {
    allow_credentials = false
    allow_headers     = ["date", "keep-alive"]
    allow_methods     = ["PUT", "POST"]
    allow_origins     = ["*"]
    expose_headers    = ["keep-alive"]
    max_age           = 5
  }

  integrations = {
    "ANY /{proxy+}" = {
      lambda_arn             = module.lambda_function.lambda_function_arn
      payload_format_version = "2.0"
      timeout_milliseconds  = 12000
      }
    }
}
```