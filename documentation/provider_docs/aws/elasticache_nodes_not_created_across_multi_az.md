---
title: "ElastiCache Nodes Not Created Across Multi AZ"
meta:
  name: "aws/elasticache_nodes_not_created_across_multi_az"
  id: "6db03a91-f933-4f13-ab38-a8b87a7de54d"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Availability"
---

## Metadata
**Name:** `aws/elasticache_nodes_not_created_across_multi_az`

**Id:** `6db03a91-f933-4f13-ab38-a8b87a7de54d`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Availability

## Description
ElastiCache Nodes should be created across multi az, which means 'az_mode' should be set to 'cross-az' in multi nodes cluster

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elasticache_cluster)

## Non-Compliant Code Examples
```terraform
resource "aws_elasticache_cluster" "positive1" {
  cluster_id = "cluster-example"
  engine = "memcached"
  num_cache_nodes = 3
}

resource "aws_elasticache_cluster" "positive2" {
  cluster_id = "cluster-example"
  engine = "memcached"
  num_cache_nodes = 3

  az_mode = "single-az"
}
```

## Compliant Code Examples
```terraform
resource "aws_elasticache_cluster" "negative1" {
  cluster_id           = "cluster-example"
  engine = "memcached"

  num_cache_nodes = 3

  az_mode = "cross-az"
}
```