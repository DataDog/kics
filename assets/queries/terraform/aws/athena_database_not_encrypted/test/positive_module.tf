module "athena_database" {
  source  = "terraform-aws-modules/athena/aws"
  version = "~> 2.0"

  name   = "my_athena_db"
  bucket = "my_athena_bucket"

  force_destroy = true
}