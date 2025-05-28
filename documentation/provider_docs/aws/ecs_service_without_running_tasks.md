---
title: "ECS Service Without Running Tasks"
meta:
  name: "aws/ecs_service_without_running_tasks"
  id: "91f16d09-689e-4926-aca7-155157f634ed"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Availability"
---

## Metadata
**Name:** `aws/ecs_service_without_running_tasks`

**Id:** `91f16d09-689e-4926-aca7-155157f634ed`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Availability

## Description
ECS Service should have at least 1 task running

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs_service)

## Non-Compliant Code Examples
```terraform
resource "aws_ecs_service" "positive1" {
  name    = "positive1"
  cluster = aws_ecs_cluster.example.id
  desired_count   = 0
}

```

## Compliant Code Examples
```terraform
resource "aws_ecs_service" "negative1" {
  name    = "negative1"
  cluster = aws_ecs_cluster.example.id

  deployment_maximum_percent         = 200
  deployment_minimum_healthy_percent = 100
}

resource "aws_ecs_service" "km_ecs_service" {
  name            = "km_ecs_service_${var.environment}"
  cluster         = aws_ecs_cluster.km_ecs_cluster.id
  task_definition = aws_ecs_task_definition.km_ecs_task.arn
  desired_count   = 1
  launch_type     = "FARGATE"

  load_balancer {
    target_group_arn = var.elb_target_group_arn
    container_name   = "km-frontend"
    container_port   = 80
  }
  network_configuration {
    assign_public_ip = true
    subnets          = var.private_subnet
    security_groups  = [ var.elb_sg ]
  }
  tags = merge(var.default_tags, {
  })
}

```