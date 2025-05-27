---
title: "API Gateway Deployment Without Access Log Setting"
meta:
  name: "aws/api_gateway_deployment_without_access_log_setting"
  id: "625abc0e-f980-4ac9-a775-f7519ee34296"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Observability"
---

## Metadata
**Name:** `aws/api_gateway_deployment_without_access_log_setting`

**Id:** `625abc0e-f980-4ac9-a775-f7519ee34296`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Observability

## Description
API Gateway Deployment should have access log setting defined when connected to an API Gateway Stage.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_deployment)

## Non-Compliant Code Examples
```terraform
resource "aws_api_gateway_deployment" "example3" {
  rest_api_id   = "some rest api id"
  stage_name = "some name"
  tags {
    project = "ProjectName"
  }
}

resource "aws_api_gateway_stage" "example000" {
  deployment_id = aws_api_gateway_deployment.example3.id
  rest_api_id   = aws_api_gateway_rest_api.example.id
  stage_name    = "example"
}

```

```terraform
resource "aws_api_gateway_deployment" "example4" {
  rest_api_id   = "some rest api id"
  stage_name = "some name"
  tags {
    project = "ProjectName"
  }
}

resource "aws_api_gateway_stage" "example0000" {
  deployment_id = aws_api_gateway_deployment.example4.id
  rest_api_id   = aws_api_gateway_rest_api.example.id
  stage_name    = "example"

  access_log_settings {
    destination_arn = "dest"
    format = "format"
  }
}

```

```terraform
resource "aws_api_gateway_deployment" "examplee" {
  rest_api_id   = "some rest api id"
  stage_name = "some name"
  tags {
    project = "ProjectName"
  }
}

resource "aws_api_gateway_stage" "example00" {
  deployment_id = aws_api_gateway_deployment.example.id
  rest_api_id   = aws_api_gateway_rest_api.example.id
  stage_name    = "example"
}

```

## Compliant Code Examples
```terraform
resource "aws_api_gateway_deployment" "example5" {
  rest_api_id   = "some rest api id"
  stage_name = "some name"
  stage_description = "some description"

  tags {
    project = "ProjectName"
  }
}

resource "aws_api_gateway_stage" "example0" {
  deployment_id = aws_api_gateway_deployment.example5.id
  rest_api_id   = aws_api_gateway_rest_api.example.id
  stage_name    = "example"

  access_log_settings {
    destination_arn = "dest"
    format = "format"
  }
}

```