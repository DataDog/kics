---
title: "API Gateway Without SSL Certificate"
meta:
  name: "aws/api_gateway_without_ssl_certificate"
  id: "0b4869fc-a842-4597-aa00-1294df425440"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `aws/api_gateway_without_ssl_certificate`

**Id:** `0b4869fc-a842-4597-aa00-1294df425440`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Insecure Configurations

## Description
SSL Client Certificate should be enabled

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_stage#client_certificate_id)

## Non-Compliant Code Examples
```terraform
resource "aws_api_gateway_stage" "positive1" {
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


  client_certificate_id = "12131323"

}

```