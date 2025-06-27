---
title: "ECS Cluster with Container Insights Disabled"
meta:
  name: "aws/ecs_cluster_container_insights_disabled"
  id: "97cb0688-369a-4d26-b1f7-86c4c91231bc"
  display_name: "ECS Cluster with Container Insights Disabled"
  cloud_provider: "aws"
  platform: "Terraform"
  severity: "LOW"
  category: "Observability"
---
## Metadata

**Name:** `aws/ecs_cluster_container_insights_disabled`

**Query Name** `ECS Cluster with Container Insights Disabled`

**Id:** `97cb0688-369a-4d26-b1f7-86c4c91231bc`

**Cloud Provider:** aws

**Platform** Terraform

**Severity:** Low

**Category:** Observability

## Description
ECS clusters should have the `containerInsights` setting enabled to provide enhanced monitoring and observability for container workloads. Without enabling container insights, as in the configuration below:

```
resource "aws_ecs_cluster" "foo" {
  name = "white-hart"
}
```

critical metrics and logs about cluster and task performance will not be collected, making it more difficult to detect anomalies, troubleshoot issues, and ensure operational health. Enabling container insights by specifying:

```
setting {
  name  = "containerInsights"
  value = "enabled"
}
```

helps provide visibility into resource utilization, failures, and capacity planning, reducing operational risk.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs_cluster#setting)


## Compliant Code Examples
```terraform
resource "aws_ecs_cluster" "foo" {
  name = "white-hart"

  setting {
    name  = "containerInsights"
    value = "enabled"
  }
}

```
## Non-Compliant Code Examples
```terraform
resource "aws_ecs_cluster" "foo" {
  name = "white-hart"

#  setting {
#    name  = "containerInsights"
#    value = "enabled"
#  }
}

```