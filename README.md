# Terraform AWS VPC

This script builds an AWS VPC, with:

* Internet Gatway (Attached to VPC)
* NAT Gateways (Server access to the Internet for updates, time checks, etc)
* Public and Private Subnets
* Public and Private Route Tables

Once built, this environment can be used as a stepping stone to build compute infrastructure on top, for example,
a EKS/K8S cluster.

## Required Inputs

This code is presented as a module, and requires the following inputs:

* Name: Name suffix to be attached to the VPC and associated items (`test, dev, qa, prod`)
* CIDR: Networking Range for the subnet. Largest that AWS allows is a /16
* Availability Zones
* Public Subnets
* Private Subnets

Please see the test.tfvars file for an example.

## Tests

To run a simple test, review and run the `$ run.sh` script.
