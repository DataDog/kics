---
title: "API Gateway Deployment Without Access Log Setting"
meta:
  name: "terraform/api_gateway_deployment_without_access_log_setting"
  id: "625abc0e-f980-4ac9-a775-f7519ee34296"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Observability"
---
## Metadata
**Name:** `terraform/api_gateway_deployment_without_access_log_setting`
**Id:** `625abc0e-f980-4ac9-a775-f7519ee34296`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Observability
## Description
API Gateway deployments should have `access_log_settings` defined for each connected API Gateway Stage to ensure proper logging of API requests and responses. Without these settings, as shown below, critical API activity may go unlogged, making it difficult to track access patterns, detect malicious requests, or troubleshoot operational issues:

```
resource "aws_api_gateway_stage" "example" {
  ...
  access_log_settings {
    destination_arn = "arn:aws:logs:region:account-id:log-group:log-group-name"
    format          = "format"
  }
}
```

Leaving this misconfiguration unaddressed can hinder security monitoring and auditing, potentially exposing the environment to undetected abuse or data exfiltration.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_deployment)

## Non-Compliant Code Examples
```aws
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

```aws
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

```aws
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

## Compliant Code Examples
```aws
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