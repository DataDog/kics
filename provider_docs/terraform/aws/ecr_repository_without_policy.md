---
title: "ECR Repository Without Policy"
meta:
  name: "terraform/ecr_repository_without_policy"
  id: "69e7c320-b65d-41bb-be02-d63ecc0bcc9d"
  cloud_provider: "terraform"
  severity: "LOW"
  category: "Best Practices"
---
## Metadata
**Name:** `terraform/ecr_repository_without_policy`
**Id:** `69e7c320-b65d-41bb-be02-d63ecc0bcc9d`
**Cloud Provider:** terraform
**Severity:** Low
**Category:** Best Practices
## Description
Amazon Elastic Container Registry (ECR) repositories should have resource policies attached to them to enforce specific access controls. Without a repository policy, the repository may rely only on default AWS account permissions, which are not granular and may inadvertently allow unauthorized users or services to access, modify, or delete container images. This misconfiguration exposes the repository contents to security risks such as privilege escalation or data theft. To mitigate this, it is important to define an `aws_ecr_repository_policy` with the appropriate access permissions for each ECR repository.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecr_repository_policy)

## Non-Compliant Code Examples
```aws
resource "aws_ecr_repository" "foo" {
  name = "bar"
}


```

```aws
resource "aws_ecr_repository" "foo2" {
  name = "bar"
}


resource "aws_ecr_repository_policy" "foopolicy" {
  repository = aws_ecr_repository.foo.name

  policy = <<EOF
{
    "Version": "2008-10-17",
    "Statement": [
        {
            "Sid": "new policy",
            "Effect": "Allow",
            "Principal": "*",
            "Action": [
                "ecr:GetDownloadUrlForLayer",
                "ecr:BatchGetImage",
                "ecr:BatchCheckLayerAvailability",
                "ecr:PutImage",
                "ecr:InitiateLayerUpload",
                "ecr:UploadLayerPart",
                "ecr:CompleteLayerUpload",
                "ecr:DescribeRepositories",
                "ecr:GetRepositoryPolicy",
                "ecr:ListImages",
                "ecr:DeleteRepository",
                "ecr:BatchDeleteImage",
                "ecr:SetRepositoryPolicy",
                "ecr:DeleteRepositoryPolicy"
            ]
        }
    ]
}
EOF
}

```

## Compliant Code Examples
```aws
resource "aws_ecr_repository" "foo" {
  name = "bar"
}

resource "aws_ecr_repository_policy" "foopolicy" {
  repository = aws_ecr_repository.foo.name

  policy = <<EOF
{
    "Version": "2008-10-17",
    "Statement": [
        {
            "Sid": "new policy",
            "Effect": "Allow",
            "Principal": "*",
            "Action": [
                "ecr:GetDownloadUrlForLayer",
                "ecr:BatchGetImage",
                "ecr:BatchCheckLayerAvailability",
                "ecr:PutImage",
                "ecr:InitiateLayerUpload",
                "ecr:UploadLayerPart",
                "ecr:CompleteLayerUpload",
                "ecr:DescribeRepositories",
                "ecr:GetRepositoryPolicy",
                "ecr:ListImages",
                "ecr:DeleteRepository",
                "ecr:BatchDeleteImage",
                "ecr:SetRepositoryPolicy",
                "ecr:DeleteRepositoryPolicy"
            ]
        }
    ]
}
EOF
}

```