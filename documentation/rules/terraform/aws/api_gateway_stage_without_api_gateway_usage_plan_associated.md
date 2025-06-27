---
title: "API Gateway Stage Without API Gateway UsagePlan Associated"
meta:
  name: "aws/api_gateway_stage_without_api_gateway_usage_plan_associated"
  id: "c999cf62-0920-40f8-8dda-0caccd66ed7e"
  display_name: "API Gateway Stage Without API Gateway UsagePlan Associated"
  cloud_provider: "aws"
  platform: "Terraform"
  severity: "LOW"
  category: "Resource Management"
---
## Metadata

**Name:** `aws/api_gateway_stage_without_api_gateway_usage_plan_associated`

**Query Name** `API Gateway Stage Without API Gateway UsagePlan Associated`

**Id:** `c999cf62-0920-40f8-8dda-0caccd66ed7e`

**Cloud Provider:** aws

**Platform** Terraform

**Severity:** Low

**Category:** Resource Management

## Description
API Gateway stages should always be associated with an API Gateway UsagePlan, which enforces throttling and quota limits for clients accessing your APIs. Without a defined `aws_api_gateway_usage_plan` resource and its association via the `api_stages` block, as shown below,

```
resource "aws_api_gateway_stage" "example" {
  deployment_id = "some deployment id"
  rest_api_id   = "some rest api id"
  stage_name    = "development"
}
```

the API stage can be accessed without usage restrictions, leading to potential misuse, abuse, or denial of service due to unlimited traffic. Configuring a UsagePlan like

```
resource "aws_api_gateway_usage_plan" "example" {
  name        = "my-usage-plan"
  description = "usage plan description"
  api_stages {
    api_id = "some rest api id"
    stage  = "development"
  }
}
```

helps mitigate these risks by controlling consumption through quotas and throttling, protecting backend resources and maintaining predictable API performance.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_stage)


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