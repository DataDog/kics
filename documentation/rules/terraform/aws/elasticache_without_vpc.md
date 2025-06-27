---
title: "ElastiCache Without VPC"
meta:
  name: "aws/elasticache_without_vpc"
  id: "8c849af7-a399-46f7-a34c-32d3dc96f1fc"
  display_name: "ElastiCache Without VPC"
  cloud_provider: "aws"
  platform: "Terraform"
  severity: "LOW"
  category: "Networking and Firewall"
---
## Metadata

**Name:** `aws/elasticache_without_vpc`

**Query Name** `ElastiCache Without VPC`

**Id:** `8c849af7-a399-46f7-a34c-32d3dc96f1fc`

**Cloud Provider:** aws

**Platform** Terraform

**Severity:** Low

**Category:** Networking and Firewall

## Description
Amazon ElastiCache clusters should be launched within a Virtual Private Cloud (VPC) to ensure that network access is restricted and controlled. When the `subnet_group_name` attribute is omitted, as shown below, ElastiCache is deployed outside a VPC, making it potentially accessible over the public internet and exposing sensitive cached data to unauthorized actors:

```
resource "aws_elasticache_cluster" "example" {
  cluster_id           = "cluster-example"
  engine               = "memcached"
  node_type            = "cache.m4.large"
  num_cache_nodes      = 2
  parameter_group_name = aws_elasticache_parameter_group.default.id
  port                 = 11211
}
```

This misconfiguration can lead to increased risk of data breaches and unauthorized access to cached application data.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elasticache_cluster#subnet_group_name)


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