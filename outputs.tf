output "bucket_name" {
  value = "${aws_s3_bucket.my-bucket.bucket}"
}

# output "bucket_arn" {
#   value = "${aws_s3_bucket.test_bucket.arn}"
# }