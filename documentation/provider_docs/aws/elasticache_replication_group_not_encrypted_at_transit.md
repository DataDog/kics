---
title: "ElastiCache Replication Group Not Encrypted At Transit"
meta:
  name: "aws/elasticache_replication_group_not_encrypted_at_transit"
  id: "1afbb3fa-cf6c-4a3d-b730-95e9f4df343e"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Encryption"
---

## Metadata
**Name:** `aws/elasticache_replication_group_not_encrypted_at_transit`

**Id:** `1afbb3fa-cf6c-4a3d-b730-95e9f4df343e`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Encryption

## Description
ElastiCache Replication Group encryption should be enabled at Transit

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elasticache_replication_group#transit_encryption_enabled)

## Non-Compliant Code Examples
```terraform
resource "aws_elasticache_replication_group" "example" {
  automatic_failover_enabled    = true
  availability_zones            = ["us-west-2a", "us-west-2b"]
  replication_group_id          = "tf-rep-group-1"
  replication_group_description = "test description"
  node_type                     = "cache.m4.large"
  number_cache_clusters         = 2
  port                          = 6379
  transit_encryption_enabled    = false
}

```

```terraform
resource "aws_elasticache_replication_group" "example" {
  automatic_failover_enabled    = true
  availability_zones            = ["us-west-2a", "us-west-2b"]
  replication_group_id          = "tf-rep-group-1"
  replication_group_description = "test description"
  node_type                     = "cache.m4.large"
  number_cache_clusters         = 2
  port                          = 6379
}

```

## Compliant Code Examples
```terraform
resource "aws_elasticache_replication_group" "example3" {
  automatic_failover_enabled    = true
  availability_zones            = ["us-west-2a", "us-west-2b"]
  replication_group_id          = "tf-rep-group-1"
  replication_group_description = "test description"
  node_type                     = "cache.m4.large"
  number_cache_clusters         = 2
  port                          = 6379
  at_rest_encryption_enabled    = true
  transit_encryption_enabled    = true
}

```