module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "~> 3.0"

  # General
  name            = var.name
  azs             = var.azs
  cidr            = var.cidr
  private_subnets = var.private_subnets
  public_subnets  = var.public_subnets

  # Gateways
  enable_nat_gateway = true
  enable_vpn_gateway = false

  # Flow Logs
  create_flow_log_cloudwatch_iam_role             = var.enable_flow_log
  create_flow_log_cloudwatch_log_group            = var.enable_flow_log
  enable_flow_log                                 = var.enable_flow_log
  flow_log_cloudwatch_log_group_retention_in_days = var.flow_log_cloudwatch_log_group_retention_in_days

  # Tags
  tags = {
    Terraform   = "true"
    Environment = var.name
  }
}
