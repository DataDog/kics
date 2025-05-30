---
title: "EKS node group remote access disabled"
meta:
  name: "aws/eks_node_group_remote_access_disabled"
  id: "ba40ace1-a047-483c-8a8d-bc2d3a67a82d"
  cloud_provider: "aws"
  severity: "MEDIUM"
  category: "Networking and Firewall"
---

## Metadata
**Name:** `aws/eks_node_group_remote_access_disabled`

**Id:** `ba40ace1-a047-483c-8a8d-bc2d3a67a82d`

**Cloud Provider:** aws

**Severity:** Medium

**Category:** Networking and Firewall

## Description
EKS node group remote access is disabled when 'SourceSecurityGroups' is missing

#### Learn More

 - [Provider Reference](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/eks_node_group#remote_access)

## Non-Compliant Code Examples
```terraform
resource "aws_eks_node_group" "positive" {
  cluster_name    = aws_eks_cluster.example.name
  node_group_name = "example"
  node_role_arn   = aws_iam_role.example.arn
  subnet_ids      = aws_subnet.example[*].id

  scaling_config {
    desired_size = 1
    max_size     = 1
    min_size     = 1
  }

  remote_access {
    ec2_ssh_key = "my-rsa-key"
  }

  # Ensure that IAM Role permissions are created before and deleted after EKS Node Group handling.
  # Otherwise, EKS will not be able to properly delete EC2 Instances and Elastic Network Interfaces.
  depends_on = [
    aws_iam_role_policy_attachment.example-AmazonEKSWorkerNodePolicy,
    aws_iam_role_policy_attachment.example-AmazonEKS_CNI_Policy,
    aws_iam_role_policy_attachment.example-AmazonEC2ContainerRegistryReadOnly,
  ]
}

```

## Compliant Code Examples
```terraform
resource "aws_eks_node_group" "negative" {
  cluster_name    = aws_eks_cluster.example.name
  node_group_name = "example"
  node_role_arn   = aws_iam_role.example.arn
  subnet_ids      = aws_subnet.example[*].id

  scaling_config {
    desired_size = 1
    max_size     = 1
    min_size     = 1
  }

  remote_access {
    ec2_ssh_key = "my-rsa-key"
    source_security_groups_ids = "sg-213120ASNE"
  }

  # Ensure that IAM Role permissions are created before and deleted after EKS Node Group handling.
  # Otherwise, EKS will not be able to properly delete EC2 Instances and Elastic Network Interfaces.
  depends_on = [
    aws_iam_role_policy_attachment.example-AmazonEKSWorkerNodePolicy,
    aws_iam_role_policy_attachment.example-AmazonEKS_CNI_Policy,
    aws_iam_role_policy_attachment.example-AmazonEC2ContainerRegistryReadOnly,
  ]
}

```