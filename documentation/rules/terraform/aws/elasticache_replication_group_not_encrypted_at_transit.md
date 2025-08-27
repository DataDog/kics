---
title: "ElastiCache replication group not encrypted at transit"
group_id: "rules/terraform/aws"
meta:
  name: "aws/elasticache_replication_group_not_encrypted_at_transit"
  id: "1afbb3fa-cf6c-4a3d-b730-95e9f4df343e"
  display_name: "ElastiCache replication group not encrypted at transit"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "MEDIUM"
  category: "Encryption"
---
## Metadata

**Id:** `1afbb3fa-cf6c-4a3d-b730-95e9f4df343e`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Medium

**Category:** Encryption

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elasticache_replication_group#transit_encryption_enabled)

### Description

 When `transit_encryption_enabled` is not set to `true` in the `aws_elasticache_replication_group` resource, data transmitted between ElastiCache nodes is not encrypted, increasing the risk of data interception or unauthorized access while data is in motion. Without encryption in transit, sensitive information can be exposed to attackers with network access, potentially leading to data breaches. Enabling transit encryption ensures all traffic between nodes is protected, safeguarding against eavesdropping and man-in-the-middle attacks.


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

```terraform
module "user_example_redis" {
  source  = "terraform-aws-modules/elasticache/aws"
  version = "~> 2.0"

  name = "cluster-example"

  engine_version = "3.2.10"
  node_type      = "cache.t2.small"
  number_cache_clusters = 2
  parameter_group_name = "default.redis3.2"

  port = 6379

  subnet_group_name = "${module.user_example_redis.cache_subnet_group_name}"
  family            = "${module.user_example_redis.cache_subnet_group_family}"

  maintenance_window = "sun:00:00-sun:03:00"
  notification_topic_arn = "arn:aws:sns:us-east-1:012345678900:my_sns_topic_name"

  security_group_ids = ["${module.sg.sg_id}"]

  transit_encryption_enabled = true

  apply_immediately     = true
  automatic_failover    = true
  auto_minor_version_upgrade = true

  tags = {
    Owner       = "user"
    Environment = "dev"
  }

  parameter {
    name  = "notify-keyspace-events"
    value = "Kg"
  }
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

```terraform
module "elasticache_redis" {
  source = "terraform-aws-modules/elasticache/aws"
  version = "0.49.0"

  namespace     = "eg"
  stage         = "prod"
  name          = "redis"
  instance_type = "cache.t2.micro"
  availability_zones = ["us-west-2a", "us-west-2b", "us-west-2c"]
  cluster_size  = 2

  description     = "Elasticache redis cluster for the foo service"
  alarm_actions   = ["${aws_kms_key.example.arn}"]
  ok_actions      = ["${aws_kms_key.example.arn}"]

  subnets         = ["subnet-a4dc27fb", "subnet-b7c9c1c5", "subnet-c9345686"]
  vpc_id          = "vpc-xxxxxx"
  allowed_cidr    = "10.10.10.10/32"
  region          = "us-west-2"
}

```