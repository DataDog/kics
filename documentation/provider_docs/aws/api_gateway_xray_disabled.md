---
title: "API Gateway X-Ray Disabled"
meta:
  name: "aws/api_gateway_xray_disabled"
  id: "5813ef56-fa94-406a-b35d-977d4a56ff2b"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Observability"
---

## Metadata
**Name:** `aws/api_gateway_xray_disabled`

**Id:** `5813ef56-fa94-406a-b35d-977d4a56ff2b`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Observability

## Description
API Gateway should have X-Ray Tracing enabled

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_stage#xray_tracing_enabled)

## Non-Compliant Code Examples
```terraform
resource "aws_api_gateway_stage" "positive1" {
  stage_name    = "prod"
  rest_api_id   = aws_api_gateway_rest_api.test.id
  deployment_id = aws_api_gateway_deployment.test.id
  xray_tracing_enabled = false
}

resource "aws_api_gateway_stage" "positive2" {
  stage_name    = "prod"
  rest_api_id   = aws_api_gateway_rest_api.test.id
  deployment_id = aws_api_gateway_deployment.test.id
}
```

## Compliant Code Examples
```terraform
resource "aws_api_gateway_stage" "negative1" {
  stage_name    = "prod"
  rest_api_id   = aws_api_gateway_rest_api.test.id
  deployment_id = aws_api_gateway_deployment.test.id
  xray_tracing_enabled = true
}
```