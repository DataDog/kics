---
title: "Elasticsearch Uses Default Security Group"
group-id: "rules/terraform/aws"
meta:
  name: "aws/elasticsearch_using_default_security_group"
  id: "d3e1f5a9-bb45-4c89-b97c-12d34ef56789"
  display_name: "Elasticsearch Uses Default Security Group"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "HIGH"
  category: "Networking and Firewall"
---
## Metadata

**Id:** `d3e1f5a9-bb45-4c89-b97c-12d34ef56789`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** High

**Category:** Networking and Firewall

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/elasticsearch_domain#vpc_options)

### Description

 AWS Elasticsearch and OpenSearch domains should have explicitly assigned security groups rather than using the default security group. When no security group is specified or an empty list is provided, the default security group is automatically assigned, which typically allows broad inbound/outbound traffic within the VPC and potentially exposes the service to unintended access from other resources. This vulnerability could lead to unauthorized access to sensitive data, potential data breaches, or service disruption.

To remediate this issue, always specify at least one security group ID in the security_group_ids list:

```
resource "aws_elasticsearch_domain" "good_example" {
  domain_name = "example"

  vpc_options {
    security_group_ids = ["sg-12345678"] // Explicit security group
  }
}
```

Avoid empty security group lists or omitting the security_group_ids attribute.


## Compliant Code Examples
```terraform
resource "aws_elasticsearch_domain" "good_example" {
  domain_name = "example"

  vpc_options {

  }
}

```

```terraform
resource "aws_opensearch_domain" "good_example" {
  domain_name = "example"

  vpc_options {
    security_group_ids = ["sg-87654321"] # ✅ Explicit security group assigned
  }
}

```

```terraform
resource "aws_elasticsearch_domain" "good_example" {
  domain_name = "example"

  vpc_options {
    security_group_ids = ["sg-12345678"] # ✅ Explicit security group assigned
  }
}

```
## Non-Compliant Code Examples
```terraform
resource "aws_elasticsearch_domain" "bad_example" {
  domain_name = "example"

  vpc_options {
    security_group_ids = []
  }
}

resource "aws_opensearch_domain" "bad_example" {
  domain_name = "example"

  vpc_options {
    security_group_ids = []
  }
}

```