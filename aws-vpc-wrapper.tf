terraform {
  backend "remote" {
    hostname     = "app.terraform.io"
    organization = "honestbank"

    workspaces {
      name = "terraform-aws-vpc"
    }
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 3.54.0, < 4.0.0"
    }
  }

  required_version = "~> 1.0"
}

resource "random_id" "vpc_random_suffix" {
  byte_length = 4
}

module "aws-vpc" {
  source = "./aws-vpc"

  name            = "${var.name}-${random_id.vpc_random_suffix.hex}"
  azs             = var.azs
  cidr            = var.cidr
  private_subnets = var.private_subnets
  public_subnets  = var.public_subnets

  enable_flow_log                                 = var.enable_flow_log
  flow_log_cloudwatch_log_group_retention_in_days = var.flow_log_cloudwatch_log_group_retention_in_days
}
