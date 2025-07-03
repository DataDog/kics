---
title: "ECS Cluster with Container Insights Disabled"
meta:
  name: "aws/ecs_cluster_container_insights_disabled"
  id: "97cb0688-369a-4d26-b1f7-86c4c91231bc"
  display_name: "ECS Cluster with Container Insights Disabled"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "LOW"
  category: "Observability"
---
## Metadata

**Id:** `97cb0688-369a-4d26-b1f7-86c4c91231bc`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Low

**Category:** Observability

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs_cluster#setting)

### Description

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