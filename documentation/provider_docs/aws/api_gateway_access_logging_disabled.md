---
title: "API Gateway Access Logging Disabled"
meta:
  name: "aws/api_gateway_access_logging_disabled"
  id: "1b6799eb-4a7a-4b04-9001-8cceb9999326"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Observability"
---

## Metadata
**Name:** `aws/api_gateway_access_logging_disabled`

**Id:** `1b6799eb-4a7a-4b04-9001-8cceb9999326`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Observability

## Description
API Gateway Stage should have Access Logging Settings defined

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_stage#access_log_settings)

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
    metrics_enabled = true
  }
}

resource "aws_apigatewayv2_stage" "postive2" {
  stage_name    = "dev"
  rest_api_id   = "id"

  access_log_settings {
    destination_arn = "dest"
  }

  default_route_settings {
    data_trace_enabled = "true"
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
}

resource "aws_apigatewayv2_stage" "postive2" {
  stage_name    = "dev"
  rest_api_id   = "id"

  access_log_settings {
    destination_arn = "dest"
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

```terraform
resource "aws_api_gateway_stage" "postive1" {
  stage_name    = "dev"
  rest_api_id   = "id"
}

resource "aws_api_gateway_method_settings" "all" {
  stage_name  = aws_api_gateway_stage.postive1.stage_name
  method_path = "*/*"

  settings {
    logging_level   = "ERROR"
  }
}

resource "aws_apigatewayv2_stage" "postive2" {
  stage_name    = "dev"
  rest_api_id   = "id"

  default_route_settings {
    logging_level   = "ERROR"
  }
}

```

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