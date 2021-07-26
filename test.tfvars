name = "labs-test-vpc-deployment"
cidr = "10.250.0.0/16"

azs             = ["ap-southeast-1a", "ap-southeast-1b", "ap-southeast-1c"]
public_subnets  = ["10.250.0.0/19", "10.250.32.0/19", "10.250.64.0/19"]
private_subnets = ["10.250.128.0/19", "10.250.160.0/19", "10.250.192.0/19"]
