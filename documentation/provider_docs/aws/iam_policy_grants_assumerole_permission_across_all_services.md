---
title: "IAM Policy Grants 'AssumeRole' Permission Across All Services"
meta:
  name: "aws/iam_policy_grants_assumerole_permission_across_all_services"
  id: "bcdcbdc6-a350-4855-ae7c-d1e6436f7c97"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Access Control"
---

## Metadata
**Name:** `aws/iam_policy_grants_assumerole_permission_across_all_services`

**Id:** `bcdcbdc6-a350-4855-ae7c-d1e6436f7c97`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Access Control

## Description
IAM Policy should not grant 'AssumeRole' permission across all services.

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role)

## Non-Compliant Code Examples
```terraform
//  Create a role which OpenShift instances will assume.
//  This role has a policy saying it can be assumed by ec2
//  instances.
resource "aws_iam_role" "positive1" {
  name = "${var.name_tag_prefix}-openshift-instance-role"

  assume_role_policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": "sts:AssumeRole",
            "Principal": {
                "Service": "ec2.amazonaws.com",
                "AWS": "*"
            },
            "Effect": "Allow",
            "Sid": ""
        }
    ]
}
EOF
}

//  This policy allows an instance to forward logs to CloudWatch, and
//  create the Log Stream or Log Group if it doesn't exist.
resource "aws_iam_policy" "positive3" {
  name        = "${var.name_tag_prefix}-openshift-instance-forward-logs"
  path        = "/"
  description = "Allows an instance to forward logs to CloudWatch"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents",
        "logs:DescribeLogStreams"
    ],
      "Resource": [
        "arn:aws:logs:*:*:*"
    ]
  }
 ]
}
EOF
}


//  Attach the policies to the role.
resource "aws_iam_policy_attachment" "positive4" {
  name       = "${var.name_tag_prefix}-openshift-attachment-forward-logs"
  roles      = ["${aws_iam_role.openshift-instance-role.name}"]
  policy_arn = "aws_iam_policy.openshift-policy-forward-logs.arn"
}

//  Create a instance profile for the role.
resource "aws_iam_instance_profile" "positive5" {
  name  = "${var.name_tag_prefix}-openshift-instance-profile"
  role = "aws_iam_role.openshift-instance-role.name"
}

resource "aws_iam_role" "positive2" {
  name = "${var.name_tag_prefix}-openshift-instance-role"

  assume_role_policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": "sts:AssumeRole",
            "Principal": {
                "Service": "ec2.amazonaws.com",
                "AWS": "*"
            },
            "Sid": ""
        }
    ]
}
EOF
}

```

## Compliant Code Examples
```terraform
//  Create a role which OpenShift instances will assume.
//  This role has a policy saying it can be assumed by ec2
//  instances.
resource "aws_iam_role" "negative1" {
  name = "${var.name_tag_prefix}-openshift-instance-role"

  assume_role_policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": "sts:AssumeRole",
            "Principal": {
                "Service": "ec2.amazonaws.com"
            },
            "Effect": "Allow",
            "Sid": ""
        }
    ]
}
EOF
}

//  This policy allows an instance to forward logs to CloudWatch, and
//  create the Log Stream or Log Group if it doesn't exist.
resource "aws_iam_policy" "negative2" {
  name        = "${var.name_tag_prefix}-openshift-instance-forward-logs"
  path        = "/"
  description = "Allows an instance to forward logs to CloudWatch"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents",
        "logs:DescribeLogStreams"
    ],
      "Resource": [
        "arn:aws:logs:*:*:*"
    ]
  }
 ]
}
EOF
}


//  Attach the policies to the role.
resource "aws_iam_policy_attachment" "negative3" {
  name       = "${var.name_tag_prefix}-openshift-attachment-forward-logs"
  roles      = ["${aws_iam_role.openshift-instance-role.name}"]
  policy_arn = "aws_iam_policy.openshift-policy-forward-logs.arn"
}

//  Create a instance profile for the role.
resource "aws_iam_instance_profile" "negative4" {
  name  = "${var.name_tag_prefix}-openshift-instance-profile"
  role = "aws_iam_role.openshift-instance-role.name"
}

```