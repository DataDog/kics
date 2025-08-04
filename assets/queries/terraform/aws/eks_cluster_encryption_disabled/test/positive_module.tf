module "eks" {
  source          = "terraform-aws-modules/eks/aws"
  cluster_name    = "my-cluster"
  cluster_version = "1.10"
  subnets         = ["subnet-12345678", "subnet-87654321"]

  vpc_id = "vpc-1234556abcdef"

  worker_groups_launch_template = [
    {
      instance_type = "m5.large"
    }
  ]
}