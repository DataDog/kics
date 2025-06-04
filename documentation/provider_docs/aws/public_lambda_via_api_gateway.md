---
title: "Public Lambda via API Gateway"
meta:
  name: "aws/public_lambda_via_api_gateway"
  id: "3ef8696c-e4ae-4872-92c7-520bb44dfe77"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `aws/public_lambda_via_api_gateway`

**Id:** `3ef8696c-e4ae-4872-92c7-520bb44dfe77`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Access Control

## Description
Allowing to run lambda function using public API Gateway

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_permission)

## Non-Compliant Code Examples
```terraform
resource "aws_lambda_permission" "apigw" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.example.function_name
  principal     = "apigateway.amazonaws.com"

  # The "/*/*" portion grants access from any method on any resource
  # within the API Gateway REST API.
  source_arn = "${aws_api_gateway_rest_api.example.execution_arn}/*/*"
}

resource "aws_lambda_function" "example" {
  function_name = "ServerlessPerson"

  handler = "MyHandler::handleRequest"
  runtime = "java11"

  role = aws_iam_role.lambda_exec.arn
}


```

## Compliant Code Examples
```terraform
resource "aws_lambda_permission" "apigw" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.example.function_name
  principal     = "apigateway.amazonaws.com"

  # The "/*/*" portion grants access from any method on any resource
  # within the API Gateway REST API.
  source_arn = "${aws_api_gateway_rest_api.example.execution_arn}/test/test"
}

resource "aws_lambda_function" "example" {
  function_name = "ServerlessPerson"

  handler = "MyHandler::handleRequest"
  runtime = "java11"

  role = aws_iam_role.lambda_exec.arn
}


```