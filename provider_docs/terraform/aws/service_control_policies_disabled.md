---
title: "Service Control Policies Disabled"
meta:
  name: "terraform/service_control_policies_disabled"
  id: "5ba6229c-8057-433e-91d0-21cf13569ca9"
  cloud_provider: "terraform"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---
## Metadata
**Name:** `terraform/service_control_policies_disabled`
**Id:** `5ba6229c-8057-433e-91d0-21cf13569ca9`
**Cloud Provider:** terraform
**Severity:** Medium
**Category:** Insecure Configurations
## Description
This check verifies whether the Amazon Organizations configuration has the `feature_set` attribute set to `"ALL"`, which enables all features, including the use of Service Control Policies (SCPs). If `feature_set` is set only to `"CONSOLIDATED_BILLING"`, as in:

```
resource "aws_organizations_organization" "example" {
  feature_set = "CONSOLIDATED_BILLING"
}
```

then organizations cannot use SCPs for centralized governance, making it difficult to enforce security and compliance policies across AWS accounts. This leaves accounts within the organization more vulnerable to misconfigurations and unauthorized access, as critical controls cannot be imposed at the organization level.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/organizations_policy)

## Non-Compliant Code Examples
```aws
resource "aws_organizations_organization" "positive1" {
  aws_service_access_principals = [
    "cloudtrail.amazonaws.com",
    "config.amazonaws.com",
  ]

  feature_set = "CONSOLIDATED_BILLING"
}

```

## Compliant Code Examples
```aws
resource "aws_organizations_organization" "negative1" {
  aws_service_access_principals = [
    "cloudtrail.amazonaws.com",
    "config.amazonaws.com",
  ]

  feature_set = "ALL"
}

```