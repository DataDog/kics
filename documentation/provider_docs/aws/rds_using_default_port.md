---
title: "RDS Using Default Port"
meta:
  name: "aws/rds_using_default_port"
  id: "bca7cc4d-b3a4-4345-9461-eb69c68fcd26"
  cloud_provider: "aws"
  severity: "LOW"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `aws/rds_using_default_port`

**Id:** `bca7cc4d-b3a4-4345-9461-eb69c68fcd26`

**Cloud Provider:** aws

**Severity:** Low

**Category:** Networking and Firewall

## Description
RDS should not use the default port (an attacker can easily guess the port). For engines related to Aurora, MariaDB or MySQL, the default port is 3306. PostgreSQL default port is 5432, Oracle default port is 1521 and SQL Server default port is 1433

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_instance#port)

## Non-Compliant Code Examples
```terraform
resource "aws_db_instance" "positive2" {
  allocated_storage    = 10
  engine               = "postgres"
  engine_version       = "5.7"
  instance_class       = "db.t3.micro"
  name                 = "mydb"
  username             = "foo"
  password             = "foobarbaz"
  skip_final_snapshot  = true
  port                 = 5432
}

```

```terraform
resource "aws_db_instance" "positive3" {
  allocated_storage    = 10
  engine               = "oracle-ee"
  engine_version       = "5.7"
  instance_class       = "db.t3.micro"
  name                 = "mydb"
  username             = "foo"
  password             = "foobarbaz"
  skip_final_snapshot  = true
  port                 = 1521
}

```

```terraform
resource "aws_db_instance" "positive4" {
  allocated_storage    = 10
  engine               = "sqlserver-ee"
  engine_version       = "5.7"
  instance_class       = "db.t3.micro"
  name                 = "mydb"
  username             = "foo"
  password             = "foobarbaz"
  skip_final_snapshot  = true
  port                 = 1433
}

```

```terraform
resource "aws_db_instance" "positive1" {
  allocated_storage    = 10
  engine               = "mysql"
  engine_version       = "5.7"
  instance_class       = "db.t3.micro"
  name                 = "mydb"
  username             = "foo"
  password             = "foobarbaz"
  parameter_group_name = aws_elasticache_parameter_group.default.id
  skip_final_snapshot  = true
  port                 = 3306
}

```

## Compliant Code Examples
```terraform
resource "aws_db_instance" "negative3" {
  allocated_storage    = 10
  engine               = "oracle-ee"
  engine_version       = "5.7"
  instance_class       = "db.t3.micro"
  name                 = "mydb"
  username             = "foo"
  password             = "foobarbaz"
  skip_final_snapshot  = true
  port                 = 1522
}

```

```terraform
resource "aws_db_instance" "negative2" {
  allocated_storage    = 10
  engine               = "postgres"
  engine_version       = "5.7"
  instance_class       = "db.t3.micro"
  name                 = "mydb"
  username             = "foo"
  password             = "foobarbaz"
  skip_final_snapshot  = true
  port                 = 5433
}

```

```terraform
resource "aws_db_instance" "negative1" {
  allocated_storage    = 10
  engine               = "mysql"
  engine_version       = "5.7"
  instance_class       = "db.t3.micro"
  name                 = "mydb"
  username             = "foo"
  password             = "foobarbaz"
  parameter_group_name = aws_elasticache_parameter_group.default.id
  skip_final_snapshot  = true
  port                 = 3307
}

```

```terraform
resource "aws_db_instance" "negative4" {
  allocated_storage    = 10
  engine               = "sqlserver-ee"
  engine_version       = "5.7"
  instance_class       = "db.t3.micro"
  name                 = "mydb"
  username             = "foo"
  password             = "foobarbaz"
  skip_final_snapshot  = true
  port                 = 1434
}

```