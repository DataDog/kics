---
title: "API Gateway access logging disabled"
group_id: "rules/terraform/aws"
meta:
  name: "aws/api_gateway_access_logging_disabled"
  id: "1b6799eb-4a7a-4b04-9001-8cceb9999326"
  display_name: "API Gateway access logging disabled"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "MEDIUM"
  category: "Observability"
---
## Metadata

**Id:** `1b6799eb-4a7a-4b04-9001-8cceb9999326`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Medium

**Category:** Observability

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_stage#access_log_settings)

### Description

 When configuring an AWS API Gateway stage in Terraform, it is important to define the `access_log_settings` block to ensure that access logs are collected and sent to a centralized logging destination, such as an Amazon CloudWatch Logs group. Without specifying the `access_log_settings`, as shown in the configuration below, API requests and responses will not be logged, making it difficult to detect anomalous activity, debug issues, or perform security investigations:

```
resource "aws_api_gateway_stage" "example" {
  stage_name  = "dev"
  rest_api_id = "id"
  // Missing access_log_settings
}
```

The absence of access logging creates a blind spot in monitoring and incident response, potentially allowing malicious activities and API misuse to go unnoticed. To address this, always include the `access_log_settings` block in your API Gateway stage resource, specifying a valid `destination_arn`:

```
resource "aws_api_gateway_stage" "example" {
  stage_name  = "dev"
  rest_api_id = "id"

  access_log_settings {
    destination_arn = "arn:aws:logs:region:account-id:log-group:your-log-group"
  }
}
```

Enabling access logging helps meet compliance requirements and establishes a robust audit trail for all API interactions.


## Compliant Code Examples
```terraform
resource "aws_api_gateway_stage" "negative1" {
  stage_name    = "dev"
  rest_api_id   = "id"

  access_log_settings {
    destination_arn = "dest"
  }
}

resource "aws_api_gateway_method_settings" "all" {
  stage_name  = aws_api_gateway_stage.negative1.stage_name
  method_path = "*/*"

  settings {
    metrics_enabled = true
    logging_level   = "ERROR"
  }
}

resource "aws_apigatewayv2_stage" "negative2" {
  stage_name    = "dev"
  rest_api_id   = "id"

  access_log_settings {
    destination_arn = "dest"
  }

  default_route_settings {
    logging_level   = "ERROR"
  }
}


```

```terraform
module "apigatewayv2" {
  source  = "terraform-aws-modules/apigateway-v2/aws"
  version = "~> 2.0"

  name          = "example"
  description   = "HTTP API Gateway"
  protocol_type = "HTTP"

  cors_configuration = {
    allow_credentials = false
    allow_headers     = ["*"]
    allow_methods     = ["GET", "POST", "OPTIONS"]
    allow_origins     = ["*"]
    expose_headers    = ["*"]
    max_age           = 300
  }

  default_route_settings = {
    detailed_metrics_enabled = true
    throttling_burst_limit  = 10
    throttling_rate_limit   = 20
    "logging_level"         = "ERROR"
  }

  target = aws_lambda_function.example.arn

  access_log_settings = {
    destination_arn = module.log_group.cloudwatch_log_group_arn
    format          = "$context.identity.sourceIp $context.identity.caller $context.identity.user [$context.requestTime] \"$context.httpMethod $context.resourcePath $context.protocol\" $context.status $context.requestLength $context.responseLength $context.requestId"
  }

  depends_on = [aws_cloudwatch_log_group.example]
}
```
## Non-Compliant Code Examples
```terraform
resource "aws_api_gateway_stage" "postive1" {
  stage_name    = "dev"
  rest_api_id   = "id"

  access_log_settings {
    destination_arn = "dest"
  }
}

resource "aws_api_gateway_method_settings" "all" {
  stage_name  = aws_api_gateway_stage.postive1.stage_name
  method_path = "*/*"

  settings {
  }
}


resource "aws_apigatewayv2_stage" "postive2" {
  stage_name    = "dev"
  rest_api_id   = "id"

  access_log_settings {
    destination_arn = "dest"
  }

  default_route_settings {
  }
}

```

```terraform
resource "aws_api_gateway_stage" "postive1" {
  stage_name    = "dev"
  rest_api_id   = "id"

  access_log_settings {
    destination_arn = "dest"
  }
}

resource "aws_api_gateway_method_settings" "all" {
  stage_name  = aws_api_gateway_stage.postive1.stage_name
  method_path = "*/*"

  settings {
    logging_level   = ""
  }
}

resource "aws_apigatewayv2_stage" "postive2" {
  stage_name    = "dev"
  rest_api_id   = "id"

  access_log_settings {
    destination_arn = "dest"
  }

  default_route_settings {
    logging_level   = ""
  }
}

```

```terraform
resource "aws_api_gateway_stage" "postive1" {
  stage_name    = "dev"
  rest_api_id   = "id"

  access_log_settings {
    destination_arn = "dest"
  }
}

resource "aws_api_gateway_method_settings" "all" {
  stage_name  = aws_api_gateway_stage.postive1.stage_name
  method_path = "*/*"

  settings {
    logging_level   = "OFF"
  }
}

resource "aws_apigatewayv2_stage" "postive2" {
  stage_name    = "dev"
  rest_api_id   = "id"

  access_log_settings {
    destination_arn = "dest"
  }

  default_route_settings {
    logging_level   = "OFF"
  }
}

```