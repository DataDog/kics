---
title: "API Gateway Without Configured Authorizer"
meta:
  name: "aws/api_gateway_without_configured_authorizer"
  id: "0a96ce49-4163-4ee6-8169-eb3b0797d694"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `aws/api_gateway_without_configured_authorizer`

**Id:** `0a96ce49-4163-4ee6-8169-eb3b0797d694`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Access Control

## Description
API Gateway REST API should have an API Gateway Authorizer

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_authorizer)

## Non-Compliant Code Examples
```terraform
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
```terraform
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