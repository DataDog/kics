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
API Gateway methods should restrict the `authorization` type to prevent unauthenticated access, except for the `OPTIONS` method used in CORS preflight requests. If you configure an API Gateway method such as:

```
resource "aws_api_gateway_method" "example" {
  http_method   = "GET"
  authorization = "NONE"
  // ...
}
```

without specifying authorization, it allows open, unauthenticated access to your API, increasing the risk of data exposure and abuse. Proper configuration requires setting `authorization = "NONE"` only for the `OPTIONS` method, for example:

```
resource "aws_api_gateway_method" "example" {
  http_method   = "OPTIONS"
  authorization = "NONE"
  // ...
}
```

This ensures that only preflight CORS requests remain unauthenticated, while all other methods require proper authorization, reducing the attack surface of your API.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_method)


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