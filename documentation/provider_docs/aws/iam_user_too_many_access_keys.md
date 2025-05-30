---
title: "IAM User Has Too Many Access Keys"
meta:
  name: "aws/iam_user_too_many_access_keys"
  id: "3561130e-9c5f-485b-9e16-2764c82763e5"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Insecure Configurations"
---

## Metadata
**Name:** `aws/iam_user_too_many_access_keys`

**Id:** `3561130e-9c5f-485b-9e16-2764c82763e5`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Insecure Configurations

## Description
Any IAM User should not have more than one access key since it increases the risk of unauthorized access and compromise credentials

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_access_key#user)

## Non-Compliant Code Examples
```terraform
resource "aws_iam_access_key" "positive1" {
  user    = aws_iam_user.lb.name
  pgp_key = "keybase:some_person_that_exists"
}

resource "aws_iam_access_key" "positive2" {
  user    = aws_iam_user.lb.name
  pgp_key = "keybase:some_person_that_exists"
}


resource "aws_iam_user" "lb" {
  name = "loadbalancer"
  path = "/system/"
}

```

## Compliant Code Examples
```terraform
resource "aws_iam_user" "userExample" {
  name = "loadbalancer"
  path = "/system/"

  tags = {
    tag-key = "tag-value"
  }
}

resource "aws_iam_access_key" "negative1" {
  user    = aws_iam_user.userExample.name
  pgp_key = "keybase:some_person_that_exists"
}


```