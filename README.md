# Terraform AWS VPC

This script builds an AWS VPC, with:

* Internet Gatway (Attached to VPC)
* NAT Gateways (Server access to the Internet for updates, time checks, etc)
* Public and Private Subnets
* Public and Private Route Tables

Once built, this environment can be used as a stepping stone to build compute infrastructure on top, for example,
a EKS/K8S cluster.

## Inputs

>
> 💾 See the [test.tfvars file](./test.tfvars) for an example.
>

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_azs"></a> [azs](#input\_azs) | az's (Availability Zones): Where, geographically to allocate the subnets, referenced using AWS's AZ codes - ie. ap-southeast-1a = Singapore AZ 'A' | `list(string)` | n/a | yes |
| <a name="input_cidr"></a> [cidr](#input\_cidr) | The CIDR allocation for the VPC. Largest is /16, smallest is /28. We use /16 to supply 8 x /19 subnets (6 active, 2 reserved for future expansion) | `any` | n/a | yes |
| <a name="input_name"></a> [name](#input\_name) | Name allocated to the VPC. Used as the VPC name and as a prefix to other items, for example subnets | `any` | n/a | yes |
| <a name="input_private_subnets"></a> [private\_subnets](#input\_private\_subnets) | A list of strings specifying the private subnet cidr Ranges. For example - ['10.250.128.0/19', '10.250.160.0/19', '10.250.192.0/19'] | `list(string)` | n/a | yes |
| <a name="input_public_subnets"></a> [public\_subnets](#input\_public\_subnets) | A list of strings specifying the public subnet cidr Ranges. For example - ['10.250.0.0/19', '10.250.32.0/19', '10.250.64.0/19'] | `list(string)` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_private_subnets"></a> [private\_subnets](#output\_private\_subnets) | List of private subnet IDs returned from terraform-aws-vpc module (aws\_subnet.private.*.id) |
| <a name="output_public_subnets"></a> [public\_subnets](#output\_public\_subnets) | List of public subnet IDs, returned from terraform-aws-vpc module (aws\_subnet.public.*.id) |
| <a name="output_vpc_id"></a> [vpc\_id](#output\_vpc\_id) | AWS identifier for the VPC provisioned by this module |
