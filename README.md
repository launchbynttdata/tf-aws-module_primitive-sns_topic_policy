# Terraform AWS Module: SNS Topic Policy

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![License: CC BY-NC-ND 4.0](https://img.shields.io/badge/License-CC_BY--NC--ND_4.0-lightgrey.svg)](https://creativecommons.org/licenses/by-nc-nd/4.0/)

## Overview

This Terraform module wraps the [aws_sns_topic_policy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/sns_topic_policy) resource to attach an IAM policy to an SNS topic. It exposes all documented arguments and attributes.

## Usage

```hcl
module "sns_topic_policy" {
  source  = "terraform.registry.launch.nttdata.com/module_primitive/sns_topic_policy/aws"
  version = "~> 1.0"

  arn    = aws_sns_topic.example.arn
  policy = data.aws_iam_policy_document.sns_topic_policy.json
}
```

The policy is typically constructed using the [aws_iam_policy_document](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_policy_document) data source.

## Pre-Commit Hooks

The `.pre-commit-config.yaml` file defines hooks for Terraform, Go, and common linting. The `commitlint` hook enforces conventional commit format. The `detect-secrets-hook` prevents new secrets from being introduced into the baseline. See the [pre-commit documentation](https://pre-commit.com/) for installation. Install the commit-msg hook manually:

```shell
pre-commit install --hook-type commit-msg
```

## Testing Locally

1. Run `make configure` to install dependencies.
2. For AWS: ensure AWS credentials are configured (e.g., `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, or `AWS_PROFILE`). Run `make env` if your Makefile provides AWS environment setup.
3. Create `examples/complete/provider.tf` with your AWS provider configuration if not delivered by the Makefile.
4. Run `make check` to run lint, validate, plan, and tests.

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.5 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.14 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_sns_topic_policy.topic_policy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/sns_topic_policy) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_arn"></a> [arn](#input\_arn) | The ARN of the SNS topic to attach the policy to. | `string` | n/a | yes |
| <a name="input_policy"></a> [policy](#input\_policy) | The fully-formed AWS policy as JSON. For more information, see the AWS IAM Policy Document Guide. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_arn"></a> [arn](#output\_arn) | The ARN of the SNS topic. |
| <a name="output_id"></a> [id](#output\_id) | The ID of the resource (same as the topic ARN). |
| <a name="output_owner"></a> [owner](#output\_owner) | The AWS Account ID of the SNS topic owner. |
<!-- END_TF_DOCS -->
