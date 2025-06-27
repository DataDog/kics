resource "aws_s3_bucket" "bucket" {
  bucket = var.bucket_name

  versioning {
    enabled = var.turn_on_versioning
  }

  tags = {
    "terraform.managed" = "true"
  }
}
