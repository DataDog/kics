---
title: "API Gateway With Invalid Compression"
meta:
  name: "aws/api_gateway_with_invalid_compression"
  id: "ed35928e-195c-4405-a252-98ccb664ab7b"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Encryption"
---

## Metadata
**Name:** `aws/api_gateway_with_invalid_compression`

**Id:** `ed35928e-195c-4405-a252-98ccb664ab7b`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Encryption

## Description
API Gateway should have valid compression, which means attribute 'minimum_compression_size' should be set and its value should be greater than -1 and smaller than 10485760.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_rest_api)

## Non-Compliant Code Examples
```terraform
resource "aws_api_gateway_rest_api" "positive1" {
  name = "regional-example"

  endpoint_configuration {
    types = ["REGIONAL"]
  }
}


resource "aws_api_gateway_rest_api" "positive2" {
  name = "regional-example"

  endpoint_configuration {
    types = ["REGIONAL"]
  }

  minimum_compression_size = -1
}


resource "aws_api_gateway_rest_api" "positive3" {
  name = "regional-example"

  endpoint_configuration {
    types = ["REGIONAL"]
  }

  minimum_compression_size = 10485760
}
```

## Compliant Code Examples
```terraform
resource "aws_api_gateway_rest_api" "negative1" {
  name = "regional-example"

  endpoint_configuration {
    types = ["REGIONAL"]
  }

  minimum_compression_size = 0
}
```