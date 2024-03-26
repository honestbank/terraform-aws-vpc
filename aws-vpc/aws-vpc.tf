terraform {
  required_version = "~> 1.0"

  # see https://www.terraform.io/language/modules/develop/providers#passing-providers-explicitly for more information
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.0"
    }
  }
}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = ">= 5.0"

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

  # DNS
  enable_dns_support   = true
  enable_dns_hostnames = var.enable_dns_hostnames

  # Tags
  tags = {
    Terraform   = "true"
    Environment = var.name
  }
}
