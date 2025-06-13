---
title: "API Gateway With CloudWatch Logging Disabled"
meta:
  name: "aws/api_gateway_with_cloudwatch_logging_disabled"
  id: "982aa526-6970-4c59-8b9b-2ce7e019fe36"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Observability"
---

## Metadata
**Name:** `aws/api_gateway_with_cloudwatch_logging_disabled`

**Id:** `982aa526-6970-4c59-8b9b-2ce7e019fe36`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Observability

## Description
AWS CloudWatch Logs for APIs should be enabled and using the naming convention described in documentation

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_stage#managing-the-api-logging-cloudwatch-log-group)

## Non-Compliant Code Examples
```terraform
variable "stage_name" {
  default = "example"
  type    = string
}
variable "stage_names" {
  default = "examples"
  type    = string
}

resource "aws_api_gateway_rest_api" "example" {
  # ... other configuration ...
}

resource "aws_api_gateway_stage" "example" {
  depends_on = [aws_cloudwatch_log_group.example]

  stage_name = var.stage_name
  # ... other configuration ...
}

resource "aws_cloudwatch_log_group" "example" {
  name              = "API-Gateway-Execution-Logs_${aws_api_gateway_rest_api.example.id}/${var.stage_names}"
  retention_in_days = 7
  # ... potentially other configuration ...
}

```

## Compliant Code Examples
```terraform
module "env" {
  source = "./env"
}

resource "aws_api_gateway_rest_api" "example" {
  # ... other configuration ...
}

resource "aws_api_gateway_stage" "example" {
  depends_on = [aws_cloudwatch_log_group.example]

  stage_name = module.env.vars.stage_name
  # ... other configuration ...
}

resource "aws_cloudwatch_log_group" "example" {
  name              = "API-Gateway-Execution-Logs_${aws_api_gateway_rest_api.example.id}/${module.env.vars.stage_name}"
  retention_in_days = 7
  # ... potentially other configuration ...
}

```

```terraform
variable "stage_name" {
  default = "example"
  type    = string
}

resource "aws_api_gateway_rest_api" "example" {
  # ... other configuration ...
}

resource "aws_api_gateway_stage" "example" {
  depends_on = [aws_cloudwatch_log_group.example]

  stage_name = var.stage_name
  # ... other configuration ...
}

resource "aws_cloudwatch_log_group" "example" {
  name              = "API-Gateway-Execution-Logs_${aws_api_gateway_rest_api.example.id}/${var.stage_name}"
  retention_in_days = 7
  # ... potentially other configuration ...
}

```