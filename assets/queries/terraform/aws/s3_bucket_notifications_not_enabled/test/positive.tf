resource "aws_s3_bucket" "bad_example" {
  bucket = "example-bucket-without-notification"
  acl    = "private"
}
