module "bucket" {
    source = "./nested_bucket"
    bucket_name = var.bucket_name
    enabled_versioning = var.enabled_versioning
}
