module "mybucket" {
  source = "../../../../../../modules/aws-simple-bucket"

  bucket_name = "dd-mybucket-eric-aws-simple-test-001"
  versioning_config = {
    enabled = true
  }
}

provider "aws" {
  region = "us-west-2"
}
