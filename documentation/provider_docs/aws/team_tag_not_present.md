---
title: "Team Tag Missing"
meta:
  name: "aws/team_tag_not_present"
  id: "a2b3c4d5-e6f7-8901-gh23-ijkl456m7890"
  cloud_provider: "aws"
  severity: "INFO"
  category: "Best Practices"
---

## Metadata
**Name:** `aws/team_tag_not_present`

**Id:** `a2b3c4d5-e6f7-8901-gh23-ijkl456m7890`

**Cloud Provider:** aws

**Severity:** Info

**Category:** Best Practices

## Description
Ensures that every cloud resource has a 'Team' tag for ownership tracking.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/guides/resource-tagging)

## Non-Compliant Code Examples
```terraform
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
```terraform
# ✅ "team" tag is not a valid attribute for this resource type
resource "aws_acm_certificate_validation" "example" {
  certificate_arn         = aws_acm_certificate.example.arn
  validation_record_fqdns = [for record in aws_route53_record.example : record.fqdn]
}

```

```terraform
resource "aws_s3_bucket" "good_example" {
  bucket = "my-bucket"

  tags = {
    team = "Security" # ✅ "team" tag is present
  }
}

```

```terraform
resource "aws_s3_bucket" "good_example" {
  bucket = "my-bucket"

  tags = {
    Team = "Security" # ✅ "Team" tag is present
  }
}

```

```terraform
resource "aws_instance" "web_subnet2" {
  ami           = "ami-123456"
  instance_type = "t2.micro"

  tags = merge({
    Name = "${local.resource_prefix.value}-subnet2"
    }, {
    git_last_modified_by = "email@email.com"
    git_modifiers        = "foo.bar"
    git_org              = "checkmarx"
    team                 = "team"
  })
}

```

```terraform
resource "aws_instance" "good_example" {
  ami           = "ami-123456"
  instance_type = "t2.micro"

  tags = {
    Team        = "DevOps" # ✅ "Team" tag is present
    Environment = "Production"
  }
}

```