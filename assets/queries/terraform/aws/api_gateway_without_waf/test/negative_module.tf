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