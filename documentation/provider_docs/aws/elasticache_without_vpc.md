---
title: "ElastiCache Without VPC"
meta:
  name: "aws/elasticache_without_vpc"
  id: "8c849af7-a399-46f7-a34c-32d3dc96f1fc"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `aws/elasticache_without_vpc`

**Id:** `8c849af7-a399-46f7-a34c-32d3dc96f1fc`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Networking and Firewall

## Description
ElastiCache should be launched in a Virtual Private Cloud (VPC)

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elasticache_cluster#subnet_group_name)

## Non-Compliant Code Examples
```terraform
resource "aws_elasticache_cluster" "positive1" {
  cluster_id           = "cluster-example"
  engine               = "memcached"
  node_type            = "cache.m4.large"
  num_cache_nodes      = 2
  parameter_group_name = aws_elasticache_parameter_group.default.id
  port                 = 11211
}

```

## Compliant Code Examples
```terraform
resource "aws_elasticache_cluster" "negative1" {
  cluster_id           = "cluster-example"
  engine               = "memcached"
  node_type            = "cache.m4.large"
  num_cache_nodes      = 2
  parameter_group_name = aws_elasticache_parameter_group.default.id
  port                 = 11211
  subnet_group_name    = var.subnet_group_name
}

```