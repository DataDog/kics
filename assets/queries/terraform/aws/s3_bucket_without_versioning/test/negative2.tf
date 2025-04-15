module "s3_bucket" {

  restrict_public_buckets = true

  version = "3.7.0"

  bucket = "my-s3-bucket"
  acl    = "private"

  versioning = {
    enabled = true
  }

}
