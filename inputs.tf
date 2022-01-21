variable "name" {
  description = "Name allocated to the VPC. Used as the VPC name and as a prefix to other items, for example subnets"
  type        = string
}

variable "cidr" {
  description = "The CIDR allocation for the VPC. Largest is /16, smallest is /28. We use /16 to supply 8 x /19 subnets (6 active, 2 reserved for future expansion)"
  type        = string
}

variable "enable_flow_log" {
  description = "Enable VPC flow logs"
  type        = bool
}

variable "flow_log_cloudwatch_log_group_retention_in_days" {
  description = "The VPC flow log retention period in days"
  type        = number
}

variable "public_subnets" {
  description = "A list of strings specifying the public subnet cidr Ranges. For example - ['10.250.0.0/19', '10.250.32.0/19', '10.250.64.0/19']"
  type        = list(string)
}

variable "private_subnets" {
  description = "A list of strings specifying the private subnet cidr Ranges. For example - ['10.250.128.0/19', '10.250.160.0/19', '10.250.192.0/19']"
  type        = list(string)
}

variable "azs" {
  description = "az's (Availability Zones): Where, geographically to allocate the subnets, referenced using AWS's AZ codes - ie. ap-southeast-1a = Singapore AZ 'A'"
  type        = list(string)
}

variable "enable_dns_hostnames" {
  description = "Wether to enable DNS hostnames in the VPC"
  default     = false
}
