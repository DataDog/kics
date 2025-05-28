---
title: "API Gateway Endpoint Config is Not Private"
meta:
  name: "aws/api_gateway_endpoint_config_is_not_private"
  id: "6b2739db-9c49-4db7-b980-7816e0c248c1"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `aws/api_gateway_endpoint_config_is_not_private`

**Id:** `6b2739db-9c49-4db7-b980-7816e0c248c1`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Networking and Firewall

## Description
The API Endpoint type in API Gateway should be set to PRIVATE so it's not exposed to the public internet

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

```

## Compliant Code Examples
```terraform
resource "aws_api_gateway_rest_api" "negative1" {
  name = "regional-example"

  endpoint_configuration {
    types = ["PRIVATE"]
  }
}

```