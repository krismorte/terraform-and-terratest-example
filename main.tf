

resource "aws_s3_bucket" "my-bucket" {
  bucket = "${var.bucket_name}"
  acl    = "private"

  tags = {
    Name        = "${var.tag_bucket_name}"
    Environment = "${var.tag_bucket_environment}"
  }
}
