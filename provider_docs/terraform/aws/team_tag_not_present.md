---
title: "Team Tag Missing"
meta:
  name: "terraform/team_tag_not_present"
  id: "a2b3c4d5-e6f7-8901-gh23-ijkl456m7890"
  cloud_provider: "terraform"
  severity: "INFO"
  category: "Best Practices"
---
## Metadata
**Name:** `terraform/team_tag_not_present`
**Id:** `a2b3c4d5-e6f7-8901-gh23-ijkl456m7890`
**Cloud Provider:** terraform
**Severity:** Info
**Category:** Best Practices
## Description
This check ensures that every cloud resource defined in Terraform includes a "Team" tag within the `tags` attribute, which is critical for tracking resource ownership and accountability. Without a "Team" tag (e.g., `tags = { Environment = "Production" }`), it becomes difficult to determine resource owners, leading to challenges in cost allocation, incident response, and lifecycle management. If left unaddressed, the absence of this tag can result in unmanaged resources remaining active, increasing the risk of security vulnerabilities and unnecessary expenses.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/guides/resource-tagging)

## Non-Compliant Code Examples
```aws
resource "aws_instance" "bad_example" {
  ami           = "ami-123456"
  instance_type = "t2.micro"

  tags = {
    Environment = "Production" # ❌ Missing "Team" tag
  }
}

resource "aws_s3_bucket" "bad_example" {
  bucket = "my-bucket"

  # ❌ No tags at all
}

```

## Compliant Code Examples
```aws
resource "aws_s3_bucket" "good_example" {
  bucket = "my-bucket"

  tags = {
    Team = "Security" # ✅ "Team" tag is present
  }
}

```

```aws
# ✅ "team" tag is not a valid attribute for this resource type
resource "aws_acm_certificate_validation" "example" {
  certificate_arn         = aws_acm_certificate.example.arn
  validation_record_fqdns = [for record in aws_route53_record.example : record.fqdn]
}

```

```aws
resource "aws_instance" "good_example" {
  ami           = "ami-123456"
  instance_type = "t2.micro"

  tags = {
    Team        = "DevOps" # ✅ "Team" tag is present
    Environment = "Production"
  }
}

```