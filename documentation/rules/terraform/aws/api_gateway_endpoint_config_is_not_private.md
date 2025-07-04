---
title: "API Gateway Endpoint Config is Not Private"
group-id: "rules/terraform/aws"
meta:
  name: "aws/api_gateway_endpoint_config_is_not_private"
  id: "6b2739db-9c49-4db7-b980-7816e0c248c1"
  display_name: "API Gateway Endpoint Config is Not Private"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---
## Metadata

**Id:** `6b2739db-9c49-4db7-b980-7816e0c248c1`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Medium

**Category:** Networking and Firewall

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/api_gateway_rest_api)

### Description

 This check ensures that the `endpoint_configuration.types` attribute for `aws_api_gateway_rest_api` resources is set to `"PRIVATE"`, rather than `"REGIONAL"` or `"EDGE"`. By exposing API Gateway endpoints to the public internet (e.g., with `"REGIONAL"`), sensitive services can be accessed or exploited by unauthorized parties. Setting the endpoint type to `"PRIVATE"` restricts access to only sources within your VPC, mitigating the risk of data exposure or malicious traffic.

```
resource "aws_api_gateway_rest_api" "secure_example" {
  name = "private-endpoint"

  endpoint_configuration {
    types = ["PRIVATE"]
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
## Non-Compliant Code Examples
```terraform
resource "aws_api_gateway_rest_api" "positive1" {
  name = "regional-example"

  endpoint_configuration {
    types = ["REGIONAL"]
  }
}

```