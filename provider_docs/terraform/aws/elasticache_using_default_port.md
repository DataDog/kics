---
title: "ElastiCache Using Default Port"
meta:
  name: "terraform/elasticache_using_default_port"
  id: "5d89db57-8b51-4b38-bb76-b9bd42bd40f0"
  cloud_provider: "terraform"
  severity: "LOW"
  category: "Networking and Firewall"
---
## Metadata
**Name:** `terraform/elasticache_using_default_port`
**Id:** `5d89db57-8b51-4b38-bb76-b9bd42bd40f0`
**Cloud Provider:** terraform
**Severity:** Low
**Category:** Networking and Firewall
## Description
ElastiCache clusters should avoid using the default ports (`6379` for Redis, `11211` for Memcached), as attackers commonly scan these ports to find and exploit exposed services. By explicitly configuring a non-default port in the `port` attribute of the `aws_elasticache_cluster` resource, you reduce the risk of automated attacks or unauthorized access. Leaving the default port unchanged makes it easier for malicious actors to guess the service endpoint and attempt brute force or exploitation attempts.

A secure Terraform example:

```
resource "aws_elasticache_cluster" "secure_example" {
  cluster_id           = "cluster"
  engine               = "redis"
  node_type            = "cache.m5.large"
  num_cache_nodes      = 1
  parameter_group_name = aws_elasticache_parameter_group.default.id
  port                 = 6380
}
```

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elasticache_cluster#port)

## Non-Compliant Code Examples
```aws
resource "aws_elasticache_cluster" "positive1" {
  cluster_id           = "cluster"
  engine               = "redis"
  node_type            = "cache.m5.large"
  num_cache_nodes      = 1
  parameter_group_name = aws_elasticache_parameter_group.default.id
}

```

```aws
resource "aws_elasticache_cluster" "positive2" {
  cluster_id           = "cluster"
  engine               = "memcached"
  node_type            = "cache.m5.large"
  num_cache_nodes      = 1
  parameter_group_name = aws_elasticache_parameter_group.default.id
  port                 = 11211
}

```

```aws
resource "aws_elasticache_cluster" "positive3" {
  cluster_id           = "cluster"
  engine               = "redis"
  node_type            = "cache.m5.large"
  num_cache_nodes      = 1
  parameter_group_name = aws_elasticache_parameter_group.default.id
  port                 = 6379
}

```

## Compliant Code Examples
```aws
resource "aws_elasticache_cluster" "negative1" {
  cluster_id           = "cluster"
  engine               = "redis"
  node_type            = "cache.m5.large"
  num_cache_nodes      = 1
  parameter_group_name = aws_elasticache_parameter_group.default.id
  port                 = 6380
}

```

```aws
resource "aws_elasticache_cluster" "negative2" {
  cluster_id           = "cluster"
  engine               = "memcached"
  node_type            = "cache.m5.large"
  num_cache_nodes      = 1
  parameter_group_name = aws_elasticache_parameter_group.default.id
  port                 = 11212
}

```