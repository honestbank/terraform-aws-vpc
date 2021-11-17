output "private_subnets" {
  description = "List of private subnet IDs returned from terraform-aws-vpc module (aws_subnet.private.*.id)"
  value       = module.vpc.private_subnets
}

output "public_subnets" {
  description = "List of public subnet IDs, returned from terraform-aws-vpc module (aws_subnet.public.*.id)"
  value       = module.vpc.public_subnets
}

output "private_route_table_ids" {
  description = "List of IDs of private route tables."
  value       = module.vpc.private_route_table_ids
}

output "public_route_table_ids" {
  description = "List of IDs of public route tables."
  value       = module.vpc.public_route_table_ids
}

output "vpc_id" {
  description = "AWS identifier for the VPC provisioned by this module"
  value       = module.vpc.vpc_id
}
