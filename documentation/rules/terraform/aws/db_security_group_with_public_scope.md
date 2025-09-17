---
title: "DB security group with public scope"
group_id: "rules/terraform/aws"
meta:
  name: "aws/db_security_group_with_public_scope"
  id: "1e0ef61b-ad85-4518-a3d3-85eaad164885"
  display_name: "DB security group with public scope"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "CRITICAL"
  category: "Networking and Firewall"
---
## Metadata

**Id:** `1e0ef61b-ad85-4518-a3d3-85eaad164885`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** Critical

**Category:** Networking and Firewall

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/rgeraskin/aws3/latest/docs/resources/db_security_group)

### Description

 AWS DB security groups with overly permissive ingress rules (`0.0.0.0/0` or `::/0`) expose database instances to potential unauthorized access from any IP address on the internet. This critical security vulnerability could lead to data breaches, unauthorized data manipulation, or complete database compromise. Instead of using public CIDR ranges, restrict access to specific IP ranges that require database connectivity.

Insecure example:
```terraform
resource "aws_db_security_group" "insecure" {
  name = "rds_sg"
  ingress {
    cidr = "0.0.0.0/0"
  }
}
```

Secure example:
```terraform
resource "aws_db_security_group" "secure" {
  name = "rds_sg"
  ingress {
    cidr = "10.0.0.0/25"
  }
}
```


## Compliant Code Examples
```terraform
resource "aws_db_security_group" "negative1" {
  name = "rds_sg"

  ingress {
    cidr = "10.0.0.0/25"
  }
}
```

```terraform
module "rds" {
  source = "terraform-aws-modules/rds/aws"
  version = "~> 3.0"

  identifier = "demodb"

  engine            = "mysql"
  engine_version    = "5.7.19"
  instance_class    = "db.t2.large"
  allocated_storage = 5
  auto_minor_version_upgrade = true
  db_name  = "demodb"
  username = "user"
  password = "YourPwdShouldBeLongAndSecure!"
  port     = "3306"

  iam_database_authentication_enabled = true

  vpc_security_group_ids = ["sg-12345678"]

  maintenance_window = "Mon:00:00-Mon:03:00"
  backup_window      = "03:00-06:00"

  # Enhanced Monitoring - see example for details on how to create the role
  # by default it is not enabled
  # You can also set this to "default" to automatically create the role with the correct permissions
  monitoring_interval = "30"
  monitoring_role_name = "MyRDSMonitoringRole"
  create_monitoring_role = true


  tags = {
    Owner       = "user"
    Environment = "dev"
  }

  # DB subnet group
  subnet_ids = ["subnet-abcdef12", "subnet-abcdef34", "subnet-abcdef56"]

  # DB parameter group
  family = "mysql5.7"

  # DB option group
  major_engine_version = "5.7"

  # Database Deletion Protection
  deletion_protection = true

  parameters = [
    {
      name = "character_set_client"
      value = "utf8mb4"
    },
    {
      name = "character_set_server"
      value = "utf8mb4"
    }
  ]

  options = [
    {
      option_name = "MARIADB_AUDIT_PLUGIN"

      option_settings = [
        {
          name  = "SERVER_AUDIT_EVENTS"
          value = "CONNECT"
        },
        {
          name  = "SERVER_AUDIT_FILE_ROTATIONS"
          value = "37"
        },
      ]
    },
  ]

  ingress = [
    {
      cidr = "10.20.0.0/16"
    },
    {
      security_group_name = "test"
    }
  ]
}
```
## Non-Compliant Code Examples
```terraform
resource "aws_db_security_group" "positive1" {
  name = "rds_sg"

  ingress {
    cidr = "0.0.0.0/0"
  }
}
```

```terraform
module "rds" {
  source = "terraform-aws-modules/rds/aws"
  version = "~> 3.0"

  identifier = "demodb"

  engine            = "mysql"
  engine_version    = "5.7.19"
  instance_class    = "db.t2.large"
  allocated_storage = 5
  auto_minor_version_upgrade = true
  db_name  = "demodb"
  username = "user"
  password = "YourPwdShouldBeLongAndSecure!"
  port     = "3306"

  iam_database_authentication_enabled = true

  vpc_security_group_ids = ["sg-12345678"]

  maintenance_window = "Mon:00:00-Mon:03:00"
  backup_window      = "03:00-06:00"

  # Enhanced Monitoring - see example for details on how to create the role
  # by default it is not enabled
  # You can also set this to "default" to automatically create the role with the correct permissions
  monitoring_interval = "30"
  monitoring_role_name = "MyRDSMonitoringRole"
  create_monitoring_role = true


  tags = {
    Owner       = "user"
    Environment = "dev"
  }

  # DB subnet group
  subnet_ids = ["subnet-abcdef12", "subnet-abcdef34", "subnet-abcdef56"]

  # DB parameter group
  family = "mysql5.7"

  # DB option group
  major_engine_version = "5.7"

  # Database Deletion Protection
  deletion_protection = true

  parameters = [
    {
      name = "character_set_client"
      value = "utf8mb4"
    },
    {
      name = "character_set_server"
      value = "utf8mb4"
    }
  ]

  options = [
    {
      option_name = "MARIADB_AUDIT_PLUGIN"

      option_settings = [
        {
          name  = "SERVER_AUDIT_EVENTS"
          value = "CONNECT"
        },
        {
          name  = "SERVER_AUDIT_FILE_ROTATIONS"
          value = "37"
        },
      ]
    },
  ]

  ingress = [
    {
      cidr = "10.20.0.0/16"
    },
    {
      security_group_name = "test"
    },
    {
      cidr = "0.0.0.0/0"
    }
  ]
}
```