# Complete Example - SNS Topic Policy

This example creates an SNS topic and attaches a policy to it using the primitive module.

## Usage

```hcl
data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

module "resource_names" {
  source   = "terraform.registry.launch.nttdata.com/module_library/resource_name/launch"
  version  = "~> 2.0"

  for_each = var.resource_names_map

  logical_product_family  = var.logical_product_family
  logical_product_service = var.logical_product_service
  class_env               = var.class_env
  instance_env            = var.instance_env
  instance_resource       = var.instance_resource
  cloud_resource_type     = each.value.name
  maximum_length          = each.value.max_length

  region = join("", split("-", data.aws_region.current.name))
}

resource "aws_sns_topic" "topic" {
  name = module.resource_names["sns_topic"].standard
  tags = var.tags
}

data "aws_iam_policy_document" "sns_topic_policy" {
  statement {
    sid    = "AllowCurrentAccountPublish"
    effect = "Allow"

    principals {
      type        = "AWS"
      identifiers = ["arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"]
    }

    actions = [
      "sns:GetTopicAttributes",
      "sns:SetTopicAttributes",
      "sns:AddPermission",
      "sns:RemovePermission",
      "sns:DeleteTopic",
      "sns:Subscribe",
      "sns:ListSubscriptionsByTopic",
      "sns:Publish"
    ]

    resources = [aws_sns_topic.topic.arn]
  }
}

module "sns_topic_policy" {
  source = "../.."

  arn    = aws_sns_topic.topic.arn
  policy = data.aws_iam_policy_document.sns_topic_policy.json
}
```

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_class_env"></a> [class_env](#input_class_env) | Class environment for resource naming (e.g., dev, prod). | `string` | n/a | yes |
| <a name="input_instance_env"></a> [instance_env](#input_instance_env) | Instance environment number for resource naming. | `number` | n/a | yes |
| <a name="input_instance_resource"></a> [instance_resource](#input_instance_resource) | Instance resource number for resource naming. | `number` | n/a | yes |
| <a name="input_logical_product_family"></a> [logical_product_family](#input_logical_product_family) | Logical product family for resource naming. | `string` | n/a | yes |
| <a name="input_logical_product_service"></a> [logical_product_service](#input_logical_product_service) | Logical product service for resource naming. | `string` | n/a | yes |
| <a name="input_resource_names_map"></a> [resource_names_map](#input_resource_names_map) | Map of resource types to naming configuration for the resource_name module. | `map(object({ name = string, max_length = optional(number, 60) }))` | n/a | yes |
| <a name="input_tags"></a> [tags](#input_tags) | Map of tags to assign to the SNS topic. | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_arn"></a> [arn](#output_arn) | The ARN of the SNS topic. |
| <a name="output_id"></a> [id](#output_id) | The ID of the SNS topic policy (same as the topic ARN). |
| <a name="output_owner"></a> [owner](#output_owner) | The AWS Account ID of the SNS topic owner. |

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.5 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.14 |

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_resource_names"></a> [resource\_names](#module\_resource\_names) | terraform.registry.launch.nttdata.com/module_library/resource_name/launch | ~> 2.0 |
| <a name="module_sns_topic_policy"></a> [sns\_topic\_policy](#module\_sns\_topic\_policy) | ../.. | n/a |

## Resources

| Name | Type |
|------|------|
| [aws_kms_key.topic](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/kms_key) | resource |
| [aws_sns_topic.topic](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/sns_topic) | resource |
| [aws_caller_identity.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity) | data source |
| [aws_iam_policy_document.sns_topic_policy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document) | data source |
| [aws_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_arn"></a> [arn](#input\_arn) | The ARN of the SNS topic to attach the policy to. Defaults to the example topic ARN. | `string` | `null` | no |
| <a name="input_class_env"></a> [class\_env](#input\_class\_env) | Class environment for resource naming (e.g., dev, prod). | `string` | n/a | yes |
| <a name="input_instance_env"></a> [instance\_env](#input\_instance\_env) | Instance environment number for resource naming. | `number` | n/a | yes |
| <a name="input_instance_resource"></a> [instance\_resource](#input\_instance\_resource) | Instance resource number for resource naming. | `number` | n/a | yes |
| <a name="input_logical_product_family"></a> [logical\_product\_family](#input\_logical\_product\_family) | Logical product family for resource naming. | `string` | n/a | yes |
| <a name="input_logical_product_service"></a> [logical\_product\_service](#input\_logical\_product\_service) | Logical product service for resource naming. | `string` | n/a | yes |
| <a name="input_policy"></a> [policy](#input\_policy) | The fully-formed AWS policy as JSON. Defaults to the example least-privilege policy document. | `string` | `null` | no |
| <a name="input_resource_names_map"></a> [resource\_names\_map](#input\_resource\_names\_map) | Map of resource types to naming configuration for the resource\_name module. | <pre>map(object({<br/>    name       = string<br/>    max_length = optional(number, 60)<br/>  }))</pre> | n/a | yes |
| <a name="input_tags"></a> [tags](#input\_tags) | Map of tags to assign to the SNS topic. | `map(string)` | `{}` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_arn"></a> [arn](#output\_arn) | The ARN of the SNS topic. |
| <a name="output_id"></a> [id](#output\_id) | The ID of the SNS topic policy (same as the topic ARN). |
| <a name="output_owner"></a> [owner](#output\_owner) | The AWS Account ID of the SNS topic owner. |
<!-- END_TF_DOCS -->
