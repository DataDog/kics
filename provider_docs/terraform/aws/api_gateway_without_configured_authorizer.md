---
title: "API Gateway Without Configured Authorizer"
meta:
  name: "terraform/api_gateway_without_configured_authorizer"
  id: "0a96ce49-4163-4ee6-8169-eb3b0797d694"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/api_gateway_without_configured_authorizer`
**Id:** `0a96ce49-4163-4ee6-8169-eb3b0797d694`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Access Control
## Description
This check ensures that all AWS API Gateway REST APIs have an associated API Gateway Authorizer, which is responsible for validating incoming requests before granting access to backend resources. Without specifying an authorizer using the `aws_api_gateway_authorizer` resource and linking it to the `rest_api_id` attribute in your Terraform configuration, APIs may be left unprotected, allowing unauthenticated and potentially malicious users to access sensitive endpoints. Failure to enforce proper authorization can result in unauthorized access, data exposure, or abuse of backend services.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_authorizer)

## Non-Compliant Code Examples
```aws
resource "aws_api_gateway_authorizer" "demo" {
  name                   = "demo"
  rest_api_id            = aws_api_gateway_rest_api.demo.id
  authorizer_uri         = aws_lambda_function.authorizer.invoke_arn
  authorizer_credentials = aws_iam_role.invocation_role.arn
}

resource "aws_api_gateway_rest_api" "demo2" {
  name = "auth-demo"
}

```

## Compliant Code Examples
```aws
resource "aws_api_gateway_authorizer" "demo" {
  name                   = "demo"
  rest_api_id            = aws_api_gateway_rest_api.demo.id
  authorizer_uri         = aws_lambda_function.authorizer.invoke_arn
  authorizer_credentials = aws_iam_role.invocation_role.arn
}

resource "aws_api_gateway_rest_api" "demo" {
  name = "auth-demo"
}

```