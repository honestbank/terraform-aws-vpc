output "private_subnets" {
  description = "List of private subnet IDs returned from terraform-aws-vpc module (aws_subnet.private.*.id)"
  value       = module.aws-vpc.private_subnets
}

output "public_subnets" {
  description = "List of public subnet IDs, returned from terraform-aws-vpc module (aws_subnet.public.*.id)"
  value       = module.aws-vpc.public_subnets
}

output "vpc_id" {
  description = "AWS identifier for the VPC provisioned by this module"
  value       = module.aws-vpc.vpc_id
}
