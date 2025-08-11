---
title: "API Gateway with open access"
group_id: "rules/terraform/aws"
meta:
  name: "aws/api_gateway_with_open_access"
  id: "15ccec05-5476-4890-ad19-53991eba1db8"
  display_name: "API Gateway with open access"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata

**Id:** `15ccec05-5476-4890-ad19-53991eba1db8`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Medium

**Category:** Insecure Configurations

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_method)

### Description

 API Gateway methods should restrict the `authorization` type to prevent unauthenticated access, except for the `OPTIONS` method used in CORS preflight requests. If you configure an API Gateway method without specifying authorization, as shown in the example below, it allows open, unauthenticated access to your API, increasing the risk of data exposure and abuse.

```
resource "aws_api_gateway_method" "example" {
  http_method   = "GET"
  authorization = "NONE"
  // ...
}
```

Proper configuration requires setting `authorization = "NONE"` only for the `OPTIONS` method. For example:

```
resource "aws_api_gateway_method" "example" {
  http_method   = "OPTIONS"
  authorization = "NONE"
  // ...
}
```

This ensures that only preflight CORS requests remain unauthenticated, while all other methods require proper authorization, reducing the attack surface of your API.


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

```terraform
module "api_gateway_method" {
  source          = "terraform-aws-modules/apigateway-v2/aws"
  version         = "3.0.0"

  name            = "my-apigateway"
  description     = "My awesome HTTP API Gateway"
  protocol_type   = "HTTP"

  custom_domain {
    domain_name = "api.${module.acm.acm_domain_name}"
    certificate_arn = module.acm.acm_certificate_arn

    endpoint_configuration = "REGIONAL"
    security_policy        = "TLS_1_2"
    stage                 = "*"
  }

  cors_configuration {
    allow_headers = ["authorization", "Content-Type", "X-Amz-Date", "X-Api-Key", "X-Amz-Security-Token"]
    allow_methods = ["OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE"]
    allow_origins = ["*"]
  }
  
  target      = "integrations"
  authorization        = "NONE"
  http_method          = "OPTIONS"
}
```
## Non-Compliant Code Examples
```terraform
module "api_gateway_method" {
  source          = "terraform-aws-modules/apigateway-v2/aws"
  version         = "3.0.0"

  name            = "my-apigateway"
  description     = "My awesome HTTP API Gateway"
  protocol_type   = "HTTP"

  custom_domain {
    domain_name = "api.${module.acm.acm_domain_name}"
    certificate_arn = module.acm.acm_certificate_arn

    endpoint_configuration = "REGIONAL"
    security_policy        = "TLS_1_2"
    stage                 = "*"
  }

  cors_configuration {
    allow_headers = ["authorization", "Content-Type", "X-Amz-Date", "X-Api-Key", "X-Amz-Security-Token"]
    allow_methods = ["OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE"]
    allow_origins = ["*"]
  }
  
  target      = "integrations"
  authorization        = "NONE"
  http_method          = "GET"
}
```

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