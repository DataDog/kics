---
title: "Service Control Policies Disabled"
meta:
  name: "aws/service_control_policies_disabled"
  id: "5ba6229c-8057-433e-91d0-21cf13569ca9"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `aws/service_control_policies_disabled`

**Id:** `5ba6229c-8057-433e-91d0-21cf13569ca9`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Insecure Configurations

## Description
Check if the Amazon Organizations ensure that all features are enabled to achieve full control over the use of AWS services and actions across multiple AWS accounts using Service Control Policies (SCPs).

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/organizations_policy)

## Non-Compliant Code Examples
```terraform
resource "aws_organizations_organization" "positive1" {
  aws_service_access_principals = [
    "cloudtrail.amazonaws.com",
    "config.amazonaws.com",
  ]

  feature_set = "CONSOLIDATED_BILLING"
}

```

## Compliant Code Examples
```terraform
resource "aws_organizations_organization" "negative1" {
  aws_service_access_principals = [
    "cloudtrail.amazonaws.com",
    "config.amazonaws.com",
  ]

  feature_set = "ALL"
}

```