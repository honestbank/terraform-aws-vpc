package test

import (
	"github.com/gruntwork-io/terratest/modules/logger"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

func TestTerraformAwsVpc(t *testing.T) {
	t.Parallel()

	awsRegion := "ap-southeast-1"
	azs := []string{"ap-southeast-1a", "ap-southeast-1b"}
	public_subnets := []string{"10.0.0.0/19"}
	private_subnets := []string{"10.0.128.0/19"}

	workingDir := test_structure.CopyTerraformFolderToTemp(t, "../aws-vpc", ".")
	logger.Logf(t, "path to test folder %s\n", workingDir)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: workingDir,
		Vars: map[string]interface{}{
			"name":            "labs-compute-vpc",
			"cidr":            "10.0.0.0/16",
			"azs":             azs,
			"public_subnets":  public_subnets,
			"private_subnets": private_subnets,
			"enable_flow_log" : true,
			"flow_log_cloudwatch_log_group_retention_in_days": 30,
		},
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
	})

	terraform.InitAndPlan(t, terraformOptions);

	//defer terraform.Destroy(t, terraformOptions)
	//
	//terraform.InitAndApply(t, terraformOptions)
	//
	//privateSubnetId := terraform.Output(t, terraformOptions, "private_subnets")
	//
	//vpcId := terraform.Output(t, terraformOptions, "vpc_id")
	//
	//subnets := aws.GetSubnetsForVpc(t, vpcId, awsRegion)
	//
	//require.Equal(t, 2, len(subnets))
	//
	//assert.False(t, aws.IsPublicSubnet(t, privateSubnetId, awsRegion))
}
