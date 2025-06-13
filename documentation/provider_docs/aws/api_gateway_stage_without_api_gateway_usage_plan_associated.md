---
title: "API Gateway Stage Without API Gateway UsagePlan Associated"
meta:
  name: "aws/api_gateway_stage_without_api_gateway_usage_plan_associated"
  id: "c999cf62-0920-40f8-8dda-0caccd66ed7e"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Resource Management"
---

## Metadata
**Name:** `aws/api_gateway_stage_without_api_gateway_usage_plan_associated`

**Id:** `c999cf62-0920-40f8-8dda-0caccd66ed7e`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Resource Management

## Description
API Gateway Stage should have API Gateway UsagePlan defined and associated.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_stage)

## Non-Compliant Code Examples
```terraform
resource "aws_api_gateway_stage" "positive1" {
  rest_api_id   = "some deployment id"
  deployment_id = "some rest api id"
  stage_name = "some name"
  tags {
    project = "ProjectName"
  }
}

resource "aws_api_gateway_stage" "positive2" {
  deployment_id = "some deployment id"
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
resource "aws_api_gateway_stage" "negative1" {
  deployment_id = "some deployment id"
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
}

```