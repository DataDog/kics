---
title: "Redis Disabled"
meta:
  name: "aws/redis_disabled"
  id: "4bd15dd9-8d5e-4008-8532-27eb0c3706d3"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `aws/redis_disabled`

**Id:** `4bd15dd9-8d5e-4008-8532-27eb0c3706d3`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Insecure Configurations

## Description
ElastiCache should have Redis enabled, since it covers Compliance Certifications such as FedRAMP, HIPAA, and PCI DSS. For more information, take a look at 'https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/SelectEngine.html'

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elasticache_cluster#engine)

## Non-Compliant Code Examples
```terraform
#this is a problematic code where the query should report a result(s)
resource "aws_elasticache_cluster" "positive1" {
  cluster_id           = "cluster-example"
  engine               = "memcached"
  node_type            = "cache.m4.large"
  num_cache_nodes      = 1
  engine_version       = "3.2.10"
  port                 = 6379
}

```

## Compliant Code Examples
```terraform
#this code is a correct code for which the query should not find any result
resource "aws_elasticache_cluster" "negative1" {
  cluster_id           = "cluster-example"
  engine               = "redis"
  node_type            = "cache.m4.large"
  num_cache_nodes      = 2
  port                 = 11211
}

```