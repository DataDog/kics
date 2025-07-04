---
title: "IMDSv1 Enabled"
group-id: "rules/terraform/aws"
meta:
  name: "aws/imdsv1_is_enabled"
  id: "f1g2h3i4-j5k6-7lmn-8opq-9012rstuvwxy"
  display_name: "IMDSv1 Enabled"
  cloud_provider: "aws"
  framework: "Terraform"
  severity: "HIGH"
  category: "Best Practices"
---
## Metadata

**Id:** `f1g2h3i4-j5k6-7lmn-8opq-9012rstuvwxy`

**Cloud Provider:** aws

**Framework:** Terraform

**Severity:** High

**Category:** Best Practices

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/instance#metadata-options)

### Description

 AWS Instance Metadata Service Version 1 (IMDSv1) is susceptible to Server-Side Request Forgery (SSRF) attacks, which can allow attackers to access instance metadata and potentially steal credentials or sensitive information from EC2 instances. IMDSv2 mitigates this risk by requiring session tokens for metadata requests, providing an additional layer of protection against SSRF vulnerabilities. To secure your infrastructure, set `http_tokens` to "required" in your AWS instance or launch template metadata options as shown in the following example:

```hcl
resource "aws_instance" "secure_example" {
  ami           = "ami-123456"
  instance_type = "t2.micro"

  metadata_options {
    http_tokens = "required"  // Secure configuration
  }
}
```


## Compliant Code Examples
```terraform
resource "aws_launch_template" "good_example" {
  name_prefix   = "example"
  image_id      = "ami-123456"
  instance_type = "t2.micro"

  metadata_options {
    http_tokens = "required" # ✅ Secure
  }
}

```

```terraform
resource "aws_instance" "good_example" {
  ami           = "ami-123456"
  instance_type = "t2.micro"

  metadata_options {
    http_tokens = "required" # ✅ Secure
  }
}

```
## Non-Compliant Code Examples
```terraform
resource "aws_instance" "bad_example" {
  ami           = "ami-123456"
  instance_type = "t2.micro"

  metadata_options {
    http_tokens = "optional" # ❌ Should be "required"
  }
}

resource "aws_launch_template" "bad_example" {
  name_prefix   = "example"
  image_id      = "ami-123456"
  instance_type = "t2.micro"

  metadata_options {
    http_tokens = "optional" # ❌ Should be "required"
  }
}

```