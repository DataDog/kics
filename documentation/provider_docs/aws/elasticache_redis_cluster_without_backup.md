---
title: "ElastiCache Redis Cluster Without Backup"
meta:
  name: "aws/elasticache_redis_cluster_without_backup"
  id: "8fdb08a0-a868-4fdf-9c27-ccab0237f1ab"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Backup"
---

## Metadata
**Name:** `aws/elasticache_redis_cluster_without_backup`

**Id:** `8fdb08a0-a868-4fdf-9c27-ccab0237f1ab`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Backup

## Description
ElastiCache Redis cluster should have 'snapshot_retention_limit' higher than 0

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elasticache_cluster#snapshot_retention_limit)

## Non-Compliant Code Examples
```terraform
resource "aws_elasticache_cluster" "positive1" {
  cluster_id           = "cluster"
  engine               = "redis"
  node_type            = "cache.m5.large"
  num_cache_nodes      = 1
  parameter_group_name = aws_elasticache_parameter_group.default.id
}

resource "aws_elasticache_cluster" "positive2" {
  cluster_id           = "cluster"
  engine               = "redis"
  node_type            = "cache.m5.large"
  num_cache_nodes      = 1
  parameter_group_name = aws_elasticache_parameter_group.default.id

  snapshot_retention_limit = 0
}

resource "aws_elasticache_parameter_group" "default" {
  name   = "cache-params"
  family = "redis2.8"

  parameter {
    name  = "activerehashing"
    value = "yes"
  }

  parameter {
    name  = "min-slaves-to-write"
    value = "2"
  }
}

```

## Compliant Code Examples
```terraform
resource "aws_elasticache_cluster" "negative1" {
  cluster_id           = "cluster"
  engine               = "redis"
  node_type            = "cache.m5.large"
  num_cache_nodes      = 1
  parameter_group_name = aws_elasticache_parameter_group.default.id

  snapshot_retention_limit = 5
}

resource "aws_elasticache_parameter_group" "default" {
  name   = "cache-params"
  family = "redis2.8"

  parameter {
    name  = "activerehashing"
    value = "yes"
  }

  parameter {
    name  = "min-slaves-to-write"
    value = "2"
  }
}

```