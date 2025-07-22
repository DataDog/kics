module "bucket" {
    source = "./module"
    bucket_name = "bucket"
    enabled_versioning = false
}
