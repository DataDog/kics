---
title: "S3 Bucket With All Permissions"
meta:
  name: "terraform/s3_bucket_with_all_permissions"
  id: "a4966c4f-9141-48b8-a564-ffe9959945bc"
  cloud_provider: "terraform"
  severity: "CRITICAL"
  category: "Access Control"
---
## Metadata
**Name:** `terraform/s3_bucket_with_all_permissions`
**Id:** `a4966c4f-9141-48b8-a564-ffe9959945bc`
**Cloud Provider:** terraform
**Severity:** Critical
**Category:** Access Control
## Description
When an S3 bucket policy allows all actions ('s3:*') to all principals ('*'), it creates a severe security vulnerability by exposing your data to unauthorized access, modification, and deletion by anyone on the internet. Instead of using overly permissive policies like 's3:*', implement the principle of least privilege by granting only specific permissions (e.g., 's3:putObject') that are required for legitimate operations. For example, replace insecure configurations like 'Action': ['s3:*'] with more restrictive ones such as 'Action': ['s3:putObject'] to ensure your S3 buckets remain protected while still enabling necessary functionality.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket)

## Non-Compliant Code Examples
```aws
resource "aws_s3_bucket" "positive1" {
  bucket = "S3B_181355"
  acl    = "private"

  policy = <<EOF
	{
	  "Id": "id113",
	  "Version": "2012-10-17",
	  "Statement": [
		{
		  "Action": [
			"s3:*"
		  ],
		  "Effect": "Allow",
		  "Resource": "arn:aws:s3:::S3B_181355/*",
		  "Principal": "*"
		}
	  ]
	}
  EOF
}

```

```aws
module "s3_bucket" {
  source = "terraform-aws-modules/s3-bucket/aws"
  version = "3.7.0"

  bucket = "my-s3-bucket"
  acl    = "private"

  versioning = {
    enabled = true
  }

  policy = <<EOF
	{
	  "Id": "id113",
	  "Version": "2012-10-17",
	  "Statement": [
		{
		  "Action": [
			"s3:*"
		  ],
		  "Effect": "Allow",
		  "Resource": "arn:aws:s3:::S3B_181355/*",
		  "Principal": "*"
		}
	  ]
	}
  EOF
}

```

## Compliant Code Examples
```aws
resource "aws_s3_bucket" "negative1" {
  bucket = "S3B_181355"
  acl    = "private"

  policy = <<EOF
	{
	  "Id": "id113",
	  "Version": "2012-10-17",
	  "Statement": [
		{
		  "Action": [
			"s3:putObject"
		  ],
		  "Effect": "Allow",
		  "Resource": "arn:aws:s3:::S3B_181355/*",
		  "Principal": "*"
		}
	  ]
	}
  EOF
}

```

```aws
module "s3_bucket" {
  source = "terraform-aws-modules/s3-bucket/aws"
  version = "3.7.0"

  bucket = "my-s3-bucket"
  acl    = "private"

  versioning = {
    enabled = true
  }

  policy = <<EOF
	{
	  "Id": "id113",
	  "Version": "2012-10-17",
	  "Statement": [
		{
		  "Action": [
			"s3:putObject"
		  ],
		  "Effect": "Allow",
		  "Resource": "arn:aws:s3:::S3B_181355/*",
		  "Principal": "*"
		}
	  ]
	}
  EOF
}

```