---
title: "API Gateway without WAF"
group_id: "rules/terraform/aws"
meta:
  name: "aws/api_gateway_without_waf"
  id: "a186e82c-1078-4a7b-85d8-579561fde884"
  display_name: "API Gateway without WAF"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata

**Id:** `a186e82c-1078-4a7b-85d8-579561fde884`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Medium

**Category:** Networking and Firewall

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/wafregional_web_acl_association#resource_arn)

### Description

 This check ensures that AWS API Gateway stages are protected by an associated Web Application Firewall (WAF) using the `aws_wafregional_web_acl_association` resource. Without a WAF enabled, exposed API endpoints are left unprotected against common web threats such as SQL injection, cross-site scripting (XSS), and other HTTP-based attacks, increasing the risk of data breaches or service disruption. To secure your API Gateway, associate a WAF Web ACL to your stage, as shown below:

```
resource "aws_wafregional_web_acl_association" "association" {
  resource_arn = aws_api_gateway_stage.negative1.arn
  web_acl_id   = aws_wafregional_web_acl.foo.id
}
```


## Compliant Code Examples
```terraform
resource "aws_wafregional_ipset" "ipset" {
  name = "tfIPSet"

  ip_set_descriptor {
    type  = "IPV4"
    value = "192.0.7.0/24"
  }
}

resource "aws_wafregional_rule" "foo" {
  name        = "tfWAFRule"
  metric_name = "tfWAFRule"

  predicate {
    data_id = aws_wafregional_ipset.ipset.id
    negated = false
    type    = "IPMatch"
  }
}

resource "aws_wafregional_web_acl" "foo" {
  name        = "foo"
  metric_name = "foo"

  default_action {
    type = "ALLOW"
  }

  rule {
    action {
      type = "BLOCK"
    }

    priority = 1
    rule_id  = aws_wafregional_rule.foo.id
  }
}

resource "aws_api_gateway_rest_api" "example" {
  body = jsonencode({
    openapi = "3.0.1"
    info = {
      title   = "example"
      version = "1.0"
    }
    paths = {
      "/path1" = {
        get = {
          x-amazon-apigateway-integration = {
            httpMethod           = "GET"
            payloadFormatVersion = "1.0"
            type                 = "HTTP_PROXY"
            uri                  = "https://ip-ranges.amazonaws.com/ip-ranges.json"
          }
        }
      }
    }
  })

  name = "example"
}

resource "aws_api_gateway_deployment" "example" {
  rest_api_id = aws_api_gateway_rest_api.example.id

  triggers = {
    redeployment = sha1(jsonencode(aws_api_gateway_rest_api.example.body))
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_api_gateway_stage" "negative1" {
  deployment_id = aws_api_gateway_deployment.example.id
  rest_api_id   = aws_api_gateway_rest_api.example.id
  stage_name    = "example"
}

resource "aws_wafregional_web_acl_association" "association" {
  resource_arn = aws_api_gateway_stage.negative1.arn
  web_acl_id   = aws_wafregional_web_acl.foo.id
}

```

```terraform
module "api_gateway_stage" {
  source     = "terraform-aws-modules/apigateway-v2/aws"
  version    = "2.1.0"

  name          = "my-api-stage"
  protocol_type = "HTTP"
  description   = "My awesome HTTP API Gateway"

  cors_configuration {
    allow_headers = ["*"]
    allow_methods = ["POST", "OPTIONS", "GET"]
    allow_origins = ["*"]
  }

  # Use default route ($default)
  default_route {
    lambda_arn             = "arn:aws:lambda:eu-west-1:135479305108:function:lambda-function-allowed-1"
    lambda_payload_format  = "2.0"
    authorization_type     = "NONE"
    authorizer_id          = "AUTHORIZER-ID"
    authorization_scopes   = ["admin:all"]
    route_key              = "$default"
    timeout_milliseconds   = 12000
  }

  # Use request/response models and api key source for 'POST /image'
  additional_routes = [
    {
      action_type            = "lambda"
      lambda_arn             = "arn:aws:lambda:eu-west-1:135479305108:function:lambda-function-allowed-2"
      lambda_payload_format  = "2.0"
      authorization_type     = "NONE"
      route_key              = "POST /image"
      timeout_milliseconds   = 12000
      model_content_type     = "multipart/form-data"
      model_schema           = "payload: {image: string}"
      api_key_source         = "HEADER"
    }
  ]

  waf_acl_id = "waf12345"

  access_log_settings = {
    destination_arn = aws_cloudwatch_log_group.example.arn
    format          = "$context.requestId"
  }

  tags = {
    ThisModule = "Coooooooooooooooooooooooool"
  }
}
```

```terraform
﻿resource "aws_api_gateway_rest_api" "example" {
  body = jsonencode({
    openapi = "3.0.1"
    info = {
      title   = "example"
      version = "1.0"
    }
    paths = {
      "/path1" = {
        get = {
          x-amazon-apigateway-integration = {
            httpMethod           = "GET"
            payloadFormatVersion = "1.0"
            type                 = "HTTP_PROXY"
            uri                  = "https://ip-ranges.amazonaws.com/ip-ranges.json"
          }
        }
      }
    }
  })

  name = "example"
}

resource "aws_api_gateway_deployment" "example" {
  rest_api_id = aws_api_gateway_rest_api.example.id

  triggers = {
    redeployment = sha1(jsonencode(aws_api_gateway_rest_api.example.body))
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_api_gateway_stage" "negative2" {
  deployment_id = aws_api_gateway_deployment.example.id
  rest_api_id   = aws_api_gateway_rest_api.example.id
  stage_name    = "example"
}

resource "aws_wafv2_web_acl" "foo" {
  name = "foo"
  scope = "REGIONAL"

  default_action {
    allow {}
  }

  visibility_config {
    cloudwatch_metrics_enabled = false
    metric_name = "foo"
    sampled_requests_enabled = false
  }
}

resource "aws_wafv2_web_acl_association" "association" {
  resource_arn = aws_api_gateway_stage.negative2.arn
  web_acl_arn   = aws_wafv2_web_acl.foo.arn
}

```
## Non-Compliant Code Examples
```terraform
resource "aws_wafregional_ipset" "ipset" {
  name = "tfIPSet"

  ip_set_descriptor {
    type  = "IPV4"
    value = "192.0.7.0/24"
  }
}

resource "aws_wafregional_rule" "foo" {
  name        = "tfWAFRule"
  metric_name = "tfWAFRule"

  predicate {
    data_id = aws_wafregional_ipset.ipset.id
    negated = false
    type    = "IPMatch"
  }
}

resource "aws_wafregional_web_acl" "foo" {
  name        = "foo"
  metric_name = "foo"

  default_action {
    type = "ALLOW"
  }

  rule {
    action {
      type = "BLOCK"
    }

    priority = 1
    rule_id  = aws_wafregional_rule.foo.id
  }
}

resource "aws_api_gateway_rest_api" "example" {
  body = jsonencode({
    openapi = "3.0.1"
    info = {
      title   = "example"
      version = "1.0"
    }
    paths = {
      "/path1" = {
        get = {
          x-amazon-apigateway-integration = {
            httpMethod           = "GET"
            payloadFormatVersion = "1.0"
            type                 = "HTTP_PROXY"
            uri                  = "https://ip-ranges.amazonaws.com/ip-ranges.json"
          }
        }
      }
    }
  })

  name = "example"
}

resource "aws_api_gateway_deployment" "example" {
  rest_api_id = aws_api_gateway_rest_api.example.id

  triggers = {
    redeployment = sha1(jsonencode(aws_api_gateway_rest_api.example.body))
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_api_gateway_stage" "positive1" {
  deployment_id = aws_api_gateway_deployment.example.id
  rest_api_id   = aws_api_gateway_rest_api.example.id
  stage_name    = "example"
}

resource "aws_wafregional_web_acl_association" "association" {
  resource_arn = aws_api_gateway_stage.positive.arn
  web_acl_id   = aws_wafregional_web_acl.foo.id
}

```

```terraform
module "api_gateway_stage" {
  source     = "terraform-aws-modules/apigateway-v2/aws"
  version    = "2.1.0"

  name          = "my-api-stage"
  protocol_type = "HTTP"
  description   = "My awesome HTTP API Gateway"

  cors_configuration {
    allow_headers = ["*"]
    allow_methods = ["POST", "OPTIONS", "GET"]
    allow_origins = ["*"]
  }

  # Use default route ($default)
  default_route {
    lambda_arn             = "arn:aws:lambda:eu-west-1:135479305108:function:lambda-function-allowed-1"
    lambda_payload_format  = "2.0"
    authorization_type     = "NONE"
    authorizer_id          = "AUTHORIZER-ID"
    authorization_scopes   = ["admin:all"]
    route_key              = "$default"
    timeout_milliseconds   = 12000
  }

  # Use request/response models and api key source for 'POST /image'
  additional_routes = [
    {
      action_type            = "lambda"
      lambda_arn             = "arn:aws:lambda:eu-west-1:135479305108:function:lambda-function-allowed-2"
      lambda_payload_format  = "2.0"
      authorization_type     = "NONE"
      route_key              = "POST /image"
      timeout_milliseconds   = 12000
      model_content_type     = "multipart/form-data"
      model_schema           = "payload: {image: string}"
      api_key_source         = "HEADER"
    }
  ]

  access_log_settings = {
    destination_arn = aws_cloudwatch_log_group.example.arn
    format          = "$context.requestId"
  }

  tags = {
    ThisModule = "Coooooooooooooooooooooooool"
  }
}
```