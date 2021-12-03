output "private_route_table_ids" {
  description = "List of IDs of private route tables."
  value       = module.aws-vpc.private_route_table_ids
}

output "private_subnets" {
  description = "List of private subnet IDs returned from terraform-aws-vpc module (aws_subnet.private.*.id)"
  value       = module.aws-vpc.private_subnets
}

output "public_route_table_ids" {
  description = "List of IDs of public route tables."
  value       = module.aws-vpc.public_route_table_ids
}

output "public_subnets" {
  description = "List of public subnet IDs, returned from terraform-aws-vpc module (aws_subnet.public.*.id)"
  value       = module.aws-vpc.public_subnets
}

output "vpc_id" {
  description = "AWS identifier for the VPC provisioned by this module"
  value       = module.aws-vpc.vpc_id
}

output "igw_id" {
  description = "AWS identifier for the Internet gateway at VPC level"
  value       = module.vpc.igw_id
}
