---
title: "API Gateway Method Does Not Contains An API Key"
meta:
  name: "aws/api_gateway_method_does_not_contains_an_api_key"
  id: "671211c5-5d2a-4e97-8867-30fc28b02216"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `aws/api_gateway_method_does_not_contains_an_api_key`

**Id:** `671211c5-5d2a-4e97-8867-30fc28b02216`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Access Control

## Description
An API Key should be required on a method request.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_method)

## Non-Compliant Code Examples
```terraform
resource "aws_api_gateway_method" "positive1" {
  rest_api_id       = aws_api_gateway_rest_api.MyDemoAPI.id
  resource_id       = aws_api_gateway_resource.MyDemoResource.id
  http_method       = "GET"
  authorization     = "NONE"
}

resource "aws_api_gateway_method" "positive2" {
  rest_api_id       = aws_api_gateway_rest_api.MyDemoAPI.id
  resource_id       = aws_api_gateway_resource.MyDemoResource.id
  http_method       = "GET"
  authorization     = "NONE"
  api_key_required  = false
}


```

## Compliant Code Examples
```terraform
resource "aws_api_gateway_method" "negative1" {
  rest_api_id       = aws_api_gateway_rest_api.MyDemoAPI.id
  resource_id       = aws_api_gateway_resource.MyDemoResource.id
  http_method       = "GET"
  authorization     = "NONE"
  api_key_required  = true
}


```