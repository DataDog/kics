---
title: "ECS Cluster with Container Insights Disabled"
meta:
  name: "aws/ecs_cluster_container_insights_disabled"
  id: "97cb0688-369a-4d26-b1f7-86c4c91231bc"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Observability"
---

## Metadata
**Name:** `aws/ecs_cluster_container_insights_disabled`

**Id:** `97cb0688-369a-4d26-b1f7-86c4c91231bc`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Observability

## Description
ECS Cluster should enable container insights

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs_cluster#setting)

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