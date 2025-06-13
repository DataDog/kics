---
title: "ECS Task Definition Network Mode Not Recommended"
meta:
  name: "aws/ecs_task_definition_network_mode_not_recommended"
  id: "9f4a9409-9c60-4671-be96-9716dbf63db1"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `aws/ecs_task_definition_network_mode_not_recommended`

**Id:** `9f4a9409-9c60-4671-be96-9716dbf63db1`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Insecure Configurations

## Description
Network_Mode should be 'awsvpc' in ecs_task_definition. AWS VPCs provides the controls to facilitate a formal process for approving and testing all network connections and changes to the firewall and router configurations

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs_task_definition#network_mode)

## Non-Compliant Code Examples
```terraform
resource "aws_ecs_task_definition" "positive1" {
  family                = "service"
  network_mode = "none"

  volume {
    name      = "service-storage"
    host_path = "/ecs/service-storage"
  }

  placement_constraints {
    type       = "memberOf"
    expression = "attribute:ecs.availability-zone in [us-west-2a, us-west-2b]"
  }
}
```

## Compliant Code Examples
```terraform
resource "aws_ecs_task_definition" "negative1" {
  family                = "service"
  network_mode = "awsvpc"

  volume {
    name      = "service-storage"
    host_path = "/ecs/service-storage"
  }

  placement_constraints {
    type       = "memberOf"
    expression = "attribute:ecs.availability-zone in [us-west-2a, us-west-2b]"
  }
}
```