---
title: "ElastiCache Using Default Port"
meta:
  name: "aws/elasticache_using_default_port"
  id: "5d89db57-8b51-4b38-bb76-b9bd42bd40f0"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `aws/elasticache_using_default_port`

**Id:** `5d89db57-8b51-4b38-bb76-b9bd42bd40f0`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Networking and Firewall

## Description
ElastiCache should not use the default port (an attacker can easily guess the port). For engine set to Redis, the default port is 6379. The Memcached default port is 11211

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elasticache_cluster#port)

## Non-Compliant Code Examples
```terraform
resource "aws_elasticache_cluster" "positive2" {
  cluster_id           = "cluster"
  engine               = "memcached"
  node_type            = "cache.m5.large"
  num_cache_nodes      = 1
  parameter_group_name = aws_elasticache_parameter_group.default.id
}

```

```terraform
resource "aws_elasticache_cluster" "positive3" {
  cluster_id           = "cluster"
  engine               = "redis"
  node_type            = "cache.m5.large"
  num_cache_nodes      = 1
  parameter_group_name = aws_elasticache_parameter_group.default.id
  port                 = 6379
}

```

```terraform
resource "aws_elasticache_cluster" "positive2" {
  cluster_id           = "cluster"
  engine               = "memcached"
  node_type            = "cache.m5.large"
  num_cache_nodes      = 1
  parameter_group_name = aws_elasticache_parameter_group.default.id
  port                 = 11211
}

```

```terraform
resource "aws_elasticache_cluster" "positive1" {
  cluster_id           = "cluster"
  engine               = "redis"
  node_type            = "cache.m5.large"
  num_cache_nodes      = 1
  parameter_group_name = aws_elasticache_parameter_group.default.id
}

```

## Compliant Code Examples
```terraform
resource "aws_elasticache_cluster" "negative2" {
  cluster_id           = "cluster"
  engine               = "memcached"
  node_type            = "cache.m5.large"
  num_cache_nodes      = 1
  parameter_group_name = aws_elasticache_parameter_group.default.id
  port                 = 11212
}

```

```terraform
resource "aws_elasticache_cluster" "negative1" {
  cluster_id           = "cluster"
  engine               = "redis"
  node_type            = "cache.m5.large"
  num_cache_nodes      = 1
  parameter_group_name = aws_elasticache_parameter_group.default.id
  port                 = 6380
}

```