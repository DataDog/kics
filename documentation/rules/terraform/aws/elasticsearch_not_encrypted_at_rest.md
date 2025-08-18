---
title: "Elasticsearch not encrypted at rest"
group_id: "rules/terraform/aws"
meta:
  name: "aws/elasticsearch_not_encrypted_at_rest"
  id: "24e16922-4330-4e9d-be8a-caa90299466a"
  display_name: "Elasticsearch not encrypted at rest"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "HIGH"
  category: "Encryption"
---
## Metadata

**Id:** `24e16922-4330-4e9d-be8a-caa90299466a`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** High

**Category:** Encryption

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elasticsearch_domain)

### Description

 This check ensures that AWS Elasticsearch domains have encryption at rest enabled, which protects sensitive data stored in Elasticsearch indices from unauthorized access if the underlying physical storage is compromised. Without encryption at rest, data stored in Elasticsearch is vulnerable to exposure if someone gains access to the underlying storage media, potentially leading to data breaches and compliance violations.

To properly secure Elasticsearch, you must explicitly configure the encrypt_at_rest block with enabled set to true:

```
resource "aws_elasticsearch_domain" "example" {
  domain_name           = "example"
  elasticsearch_version = "1.5"

  encrypt_at_rest {
      enabled = true
  }
}
```

Insecure configurations either omit the `encrypt_at_rest` block entirely or explicitly set `enabled` to `false`.


## Compliant Code Examples
```terraform
resource "aws_elasticsearch_domain" "negative1" {
  domain_name           = "example"
  elasticsearch_version = "1.5"

  encrypt_at_rest {
      enabled = true
  }
}
```

```terraform
module "elasticsearch" {
  source = "terraform-aws-modules/elasticsearch/aws"

  domain_name           = "test-domain"
  elasticsearch_version = "2.3"

  cluster_config = {
    instance_type = "m4.large.elasticsearch"
  }

  ebs_options = {
    ebs_enabled = true
    volume_size = 10
  }

  encrypt_at_rest {
    enabled = true
  }

  access_policies = <<CONFIG
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "es:*",
      "Principal": "*",
      "Effect": "Allow",
      "Resource": "arn:aws:es:eu-west-1:123456789012:domain/test-domain/*"
    }
  ]
}
CONFIG

  snapshot_options = {
    automated_snapshot_start_hour = 23
  }

  tags = {
    Domain = "TestDomain"
  }
}
```
## Non-Compliant Code Examples
```terraform
module "elasticsearch" {
  source = "terraform-aws-modules/elasticsearch/aws"

  domain_name           = "test-domain"
  elasticsearch_version = "2.3"

  cluster_config = {
    instance_type = "m4.large.elasticsearch"
  }

  ebs_options = {
    ebs_enabled = true
    volume_size = 10
  }

  access_policies = <<CONFIG
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "es:*",
      "Principal": "*",
      "Effect": "Allow",
      "Resource": "arn:aws:es:eu-west-1:123456789012:domain/test-domain/*"
    }
  ]
}
CONFIG

  snapshot_options = {
    automated_snapshot_start_hour = 23
  }

  tags = {
    Domain = "TestDomain"
  }
}
```

```terraform
resource "aws_elasticsearch_domain" "positive1" {
  domain_name           = "example"
  elasticsearch_version = "1.5"
}

resource "aws_elasticsearch_domain" "positive2" {
  domain_name           = "example"
  elasticsearch_version = "1.5"

  encrypt_at_rest {
      enabled = false
  }
}
```