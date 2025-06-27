---
title: "Global Accelerator Flow Logs Disabled"
meta:
  name: "aws/global_accelerator_flow_logs_disabled"
  id: "96e8183b-e985-457b-90cd-61c0503a3369"
  display_name: "Global Accelerator Flow Logs Disabled"
  cloud_provider: "aws"
  platform: "Terraform"
  severity: "MEDIUM"
  category: "Observability"
---
## Metadata

**Name:** `aws/global_accelerator_flow_logs_disabled`

**Query Name** `Global Accelerator Flow Logs Disabled`

**Id:** `96e8183b-e985-457b-90cd-61c0503a3369`

**Cloud Provider:** aws

**Platform** Terraform

**Severity:** Medium

**Category:** Observability

## Description
Enabling flow logs for AWS Global Accelerator allows visibility into all traffic that traverses the accelerator, providing critical data for monitoring, security auditing, and detecting anomalous activity. If the Terraform attribute `flow_logs_enabled` is not set to `true` and related fields such as `flow_logs_s3_bucket` are not specified, administrators lose valuable insight into network events, significantly hindering threat detection and incident response. Without flow logs enabled, malicious or unauthorized activity could go undetected, increasing the risk of security breaches and data exfiltration.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/globalaccelerator_accelerator#flow_logs_enabled)


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