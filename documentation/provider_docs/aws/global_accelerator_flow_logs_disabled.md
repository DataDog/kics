---
title: "Global Accelerator Flow Logs Disabled"
meta:
  name: "aws/global_accelerator_flow_logs_disabled"
  id: "96e8183b-e985-457b-90cd-61c0503a3369"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Observability"
---

## Metadata
**Name:** `aws/global_accelerator_flow_logs_disabled`

**Id:** `96e8183b-e985-457b-90cd-61c0503a3369`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Observability

## Description
Global Accelerator should have flow logs enabled

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/globalaccelerator_accelerator#flow_logs_enabled)

## Non-Compliant Code Examples
```terraform
resource "aws_globalaccelerator_accelerator" "positive2" {
  name            = "Example"
  ip_address_type = "IPV4"
  enabled         = true

  attributes {
    flow_logs_s3_bucket = "example-bucket"
    flow_logs_s3_prefix = "flow-logs/"
  }
}

```

```terraform
resource "aws_globalaccelerator_accelerator" "positive3" {
  name            = "Example"
  ip_address_type = "IPV4"
  enabled         = true

  attributes {
    flow_logs_enabled   = false
  }
}

```

```terraform
resource "aws_globalaccelerator_accelerator" "positive1" {
  name            = "Example"
  ip_address_type = "IPV4"
  enabled         = true
}

```

## Compliant Code Examples
```terraform
resource "aws_globalaccelerator_accelerator" "negative1" {
  name            = "Example"
  ip_address_type = "IPV4"
  enabled         = true

  attributes {
    flow_logs_enabled   = true
    flow_logs_s3_bucket = "example-bucket"
    flow_logs_s3_prefix = "flow-logs/"
  }
}

```