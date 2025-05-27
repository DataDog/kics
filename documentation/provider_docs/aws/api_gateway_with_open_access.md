---
title: "API Gateway With Open Access"
meta:
  name: "aws/api_gateway_with_open_access"
  id: "15ccec05-5476-4890-ad19-53991eba1db8"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `aws/api_gateway_with_open_access`

**Id:** `15ccec05-5476-4890-ad19-53991eba1db8`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Insecure Configurations

## Description
API Gateway Method should restrict the authorization type, except for the HTTP OPTIONS method.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_method)

## Non-Compliant Code Examples
```terraform
resource "aws_api_gateway_method" "positive1" {
  rest_api_id   = aws_api_gateway_rest_api.this.id
  resource_id   = aws_api_gateway_resource.this.id
  http_method   = "GET"
  authorization = "NONE"
  authorizer_id = aws_api_gateway_authorizer.this.id

  request_parameters = {
    "method.request.path.proxy" = true
  }
}

```

## Compliant Code Examples
```terraform
resource "aws_api_gateway_method" "negative1" {
  rest_api_id   = aws_api_gateway_rest_api.this.id
  resource_id   = aws_api_gateway_resource.this.id
  http_method   = "OPTIONS"
  authorization = "NONE"
  authorizer_id = aws_api_gateway_authorizer.this.id

  request_parameters = {
    "method.request.path.proxy" = true
  }
}

```