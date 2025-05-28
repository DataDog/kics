---
title: "API Gateway Deployment Without API Gateway UsagePlan Associated"
meta:
  name: "aws/api_gateway_deployment_without_api_gateway_usage_plan_associated"
  id: "b3a59b8e-94a3-403e-b6e2-527abaf12034"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Observability"
---

## Metadata
**Name:** `aws/api_gateway_deployment_without_api_gateway_usage_plan_associated`

**Id:** `b3a59b8e-94a3-403e-b6e2-527abaf12034`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Observability

## Description
API Gateway Deployment should have API Gateway UsagePlan defined and associated.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_deployment)

## Non-Compliant Code Examples
```terraform
resource "aws_api_gateway_deployment" "positive1" {
  rest_api_id   = "some rest api id"
  stage_name = "some name"
  tags {
    project = "ProjectName"
  }
}

resource "aws_api_gateway_deployment" "positive2" {
  rest_api_id   = "some rest api id"
  stage_name    = "development"
}

resource "aws_api_gateway_usage_plan" "positive3" {
  name         = "my-usage-plan"
  description  = "my description"
  product_code = "MYCODE"

  api_stages {
    api_id = "another id"
    stage  = "development"
  }
}

```

## Compliant Code Examples
```terraform
resource "aws_api_gateway_deployment" "negative1" {
  rest_api_id   = "rest_api_1"
  stage_name    = "development"
}

resource "aws_api_gateway_usage_plan" "negative2" {
  name         = "my-usage-plan"
  description  = "my description"
  product_code = "MYCODE"

  api_stages {
    api_id = "rest_api_1"
    stage  = "development"
  }

  api_stages {
    api_id = "rest_api_2"
    stage  = "development_2"
  }
}

```