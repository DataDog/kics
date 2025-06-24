resource "aws_s3_bucket" "bucket" {
  bucket = var.bucket_name
  tags = {
    "terraform.managed" = "true"
  }
}
