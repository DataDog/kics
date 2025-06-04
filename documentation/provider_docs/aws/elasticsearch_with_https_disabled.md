---
title: "Elasticsearch with HTTPS disabled"
meta:
  name: "aws/elasticsearch_with_https_disabled"
  id: "2e9e0729-66d5-4148-9d39-5e6fb4bf2a4e"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `aws/elasticsearch_with_https_disabled`

**Id:** `2e9e0729-66d5-4148-9d39-5e6fb4bf2a4e`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Networking and Firewall

## Description
Amazon Elasticsearch does not have encryption for its domains enabled. To prevent such a scenario, update the attribute 'EnforceHTTPS' to true.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elasticsearch_domain#enforce_https)

## Non-Compliant Code Examples
```terraform
provider "aws" {
  region = "us-west-2"
}

resource "aws_elasticsearch_domain" "example" {
  domain_name           = "my-elasticsearch-domain"
  elasticsearch_version = "7.10"

  cluster_config {
    instance_type = "t2.small.elasticsearch"
    instance_count = 1
    dedicated_master_enabled = false
  }

  ebs_options {
    ebs_enabled = true
    volume_type = "gp2"
    volume_size = 10
  }

  vpc_options {
    subnet_ids         = ["subnet-xxxxxxxx", "subnet-yyyyyyyy"]
    security_group_ids = ["sg-xxxxxxxx"]
  }

  domain_endpoint_options {
    enforce_https = false
  }

  tags = {
    Name        = "my-elasticsearch-domain"
    Environment = "production"
  }
}

```

## Compliant Code Examples
```terraform
provider "aws" {
  region = "us-west-2"
}

resource "aws_elasticsearch_domain" "example" {
  domain_name           = "my-elasticsearch-domain"
  elasticsearch_version = "7.10"

  cluster_config {
    instance_type = "t2.small.elasticsearch"
    instance_count = 1
    dedicated_master_enabled = false
  }

  ebs_options {
    ebs_enabled = true
    volume_type = "gp2"
    volume_size = 10
  }

  vpc_options {
    subnet_ids         = ["subnet-xxxxxxxx", "subnet-yyyyyyyy"]
    security_group_ids = ["sg-xxxxxxxx"]
  }

  domain_endpoint_options {
    enforce_https = true
  }

  tags = {
    Name        = "my-elasticsearch-domain"
    Environment = "production"
  }
}

```