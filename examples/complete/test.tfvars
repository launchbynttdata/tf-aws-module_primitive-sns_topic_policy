logical_product_family  = "launch"
logical_product_service = "sns"
class_env               = "dev"
instance_env            = 0
instance_resource       = 0

resource_names_map = {
  sns_topic = {
    name       = "sns"
    max_length = 256
  }
}

tags = {
  Environment = "test"
}
