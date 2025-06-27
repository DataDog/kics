---
title: "ElastiCache Replication Group Not Encrypted At Transit"
meta:
  name: "aws/elasticache_replication_group_not_encrypted_at_transit"
  id: "1afbb3fa-cf6c-4a3d-b730-95e9f4df343e"
  display_name: "ElastiCache Replication Group Not Encrypted At Transit"
  cloud_provider: "aws"
  platform: "Terraform"
  severity: "MEDIUM"
  category: "Encryption"
---
## Metadata

**Name:** `aws/elasticache_replication_group_not_encrypted_at_transit`

**Query Name** `ElastiCache Replication Group Not Encrypted At Transit`

**Id:** `1afbb3fa-cf6c-4a3d-b730-95e9f4df343e`

**Cloud Provider:** aws

**Platform** Terraform

**Severity:** Medium

**Category:** Encryption

## Description
When `transit_encryption_enabled` is not set to `true` in the `aws_elasticache_replication_group` resource, data transmitted between ElastiCache nodes is not encrypted, increasing the risk of data interception or unauthorized access while data is in motion. Without encryption in transit, sensitive information can be exposed to attackers with network access, potentially leading to data breaches. Enabling transit encryption ensures all traffic between nodes is protected, safeguarding against eavesdropping and man-in-the-middle attacks.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elasticache_replication_group#transit_encryption_enabled)


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