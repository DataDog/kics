---
title: "Public Lambda via API Gateway"
meta:
  name: "terraform/public_lambda_via_api_gateway"
  id: "3ef8696c-e4ae-4872-92c7-520bb44dfe77"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/public_lambda_via_api_gateway`
**Id:** `3ef8696c-e4ae-4872-92c7-520bb44dfe77`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Access Control
## Description
Allowing a Lambda function to be invoked through a public API Gateway endpoint can introduce security vulnerabilities by exposing the function to the internet. This configuration grants invocation permissions on the Lambda function from any HTTP method and any resource within the associated API Gateway, as illustrated by the use of the wildcard `"/*/*"` in the ARN. If left unaddressed, attackers could exploit the public endpoint to trigger the Lambda function, potentially resulting in unauthorized data access, denial-of-service attacks, or increased costs due to unwanted invocations. To mitigate this risk, permissions should be restricted to only necessary resources and specific methods, thereby limiting the surface area exposed to the public.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_permission)

## Non-Compliant Code Examples
```aws
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
```aws
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