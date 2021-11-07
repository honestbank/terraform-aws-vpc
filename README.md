# Terraform AWS VPC

This script builds an AWS VPC meant for use by an EKS cluster (to be built by
a downstream repo/module).

This module is currently consumed by the [api-cloud-infrastructure repo](https://github.com/honestbank/api-cloud-infrastructure).

---

>
> To regenerate this section, delete everything under the horizontal divider below and run
> `terraform-docs markdown ./ >> README.md` in the repo root.
>

---
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 3.54.0, < 4.0.0 |

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_aws-vpc"></a> [aws-vpc](#module\_aws-vpc) | ./aws-vpc | n/a |
| <a name="module_vpc"></a> [vpc](#module\_vpc) | terraform-aws-modules/vpc/aws | n/a |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_azs"></a> [azs](#input\_azs) | az's (Availability Zones): Where, geographically to allocate the subnets, referenced using AWS's AZ codes - ie. ap-southeast-1a = Singapore AZ 'A' | `list(string)` | n/a | yes |
| <a name="input_cidr"></a> [cidr](#input\_cidr) | The CIDR allocation for the VPC. Largest is /16, smallest is /28. We use /16 to supply 8 x /19 subnets (6 active, 2 reserved for future expansion) | `string` | n/a | yes |
| <a name="input_enable_flow_log"></a> [enable\_flow\_log](#input\_enable\_flow\_log) | Enable VPC flow logs | `bool` | n/a | yes |
| <a name="input_flow_log_cloudwatch_log_group_retention_in_days"></a> [flow\_log\_cloudwatch\_log\_group\_retention\_in\_days](#input\_flow\_log\_cloudwatch\_log\_group\_retention\_in\_days) | The VPC flow log retention period in days | `number` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | Name allocated to the VPC. Used as the VPC name and as a prefix to other items, for example subnets | `string` | n/a | yes |
| <a name="input_private_subnets"></a> [private\_subnets](#input\_private\_subnets) | A list of strings specifying the private subnet cidr Ranges. For example - ['10.250.128.0/19', '10.250.160.0/19', '10.250.192.0/19'] | `list(string)` | n/a | yes |
| <a name="input_public_subnets"></a> [public\_subnets](#input\_public\_subnets) | A list of strings specifying the public subnet cidr Ranges. For example - ['10.250.0.0/19', '10.250.32.0/19', '10.250.64.0/19'] | `list(string)` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_private_subnets"></a> [private\_subnets](#output\_private\_subnets) | List of private subnet IDs returned from terraform-aws-vpc module (aws\_subnet.private.*.id) |
| <a name="output_public_subnets"></a> [public\_subnets](#output\_public\_subnets) | List of public subnet IDs, returned from terraform-aws-vpc module (aws\_subnet.public.*.id) |
| <a name="output_vpc_id"></a> [vpc\_id](#output\_vpc\_id) | AWS identifier for the VPC provisioned by this module |
