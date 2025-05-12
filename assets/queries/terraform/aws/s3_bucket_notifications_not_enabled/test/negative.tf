resource "aws_s3_bucket" "good_example" {
  bucket = "example-bucket-with-notification"
}

resource "aws_s3_bucket_notification" "good_notification" {
  bucket = aws_s3_bucket.good_example.bucket

  topic {
    topic_arn = "arn:aws:sns:us-east-1:123456789012:example-topic"
    events    = ["s3:ObjectCreated:*"]
  }
}
