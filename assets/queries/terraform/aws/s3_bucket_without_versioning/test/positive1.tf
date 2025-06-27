module "mybucket" {
  source = "../../../../../../modules/aws-simple-bucket"

  bucket_name = "dd-mybucket-eric-aws-simple-test-001"
  turn_on_versioning = false
}

provider "aws" {
  region = "us-west-2"
}
