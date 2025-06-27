variable "bucket_name" {
  description = "The name of the bucket"
  type        = string
}

variable "turn_on_versioning" {
  description = "Whether we activate versioning on this bucket"
  type        = bool
}
