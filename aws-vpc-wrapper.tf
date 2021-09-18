terraform {
  backend "remote" {
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

module "aws-vpc" {
  source = "./aws-vpc"

  name = var.name

  cidr            = var.cidr
  azs             = var.azs
  private_subnets = var.private_subnets
  public_subnets  = var.public_subnets
}
