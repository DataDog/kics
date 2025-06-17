---
title: "API Gateway Deployment Without API Gateway UsagePlan Associated"
meta:
  name: "terraform/api_gateway_deployment_without_api_gateway_usage_plan_associated"
  id: "b3a59b8e-94a3-403e-b6e2-527abaf12034"
  cloud_provider: "terraform"
  severity: "LOW"
  category: "Observability"
---
## Metadata
**Name:** `terraform/api_gateway_deployment_without_api_gateway_usage_plan_associated`
**Id:** `b3a59b8e-94a3-403e-b6e2-527abaf12034`
**Cloud Provider:** terraform
**Severity:** Low
**Category:** Observability
## Description
An API Gateway Deployment should have an associated UsagePlan defined using the `aws_api_gateway_usage_plan` resource, with the `api_stages` attribute referencing the relevant API and stage. Without a UsagePlan, API Gateway endpoints are left unprotected, allowing unlimited, unauthenticated access and risking abuse, denial of service, or unexpected cost overruns. To prevent this, always associate deployments with a UsagePlan as shown below:

```
resource "aws_api_gateway_usage_plan" "example" {
  name = "usage-plan"
  api_stages {
    api_id = aws_api_gateway_deployment.example.rest_api_id
    stage  = aws_api_gateway_deployment.example.stage_name
  }
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_deployment)

## Non-Compliant Code Examples
```aws
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
```aws
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