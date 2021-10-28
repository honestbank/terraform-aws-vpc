package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

func TestTerraformAwsVpc(t *testing.T) {
	t.Parallel()

	awsRegion := "ap-southeast-1"
	azs := []string{"ap-southeast-1a", "ap-southeast-1b"}
	public_subnets := []string{"10.0.0.0/19"}
	private_subnets := []string{"10.0.128.0/19"}

	workingDir := test_structure.CopyTerraformFolderToTemp(t, "..", "aws-vpc")
	logger.Logf(t, "path to test folder %s\n", workingDir)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: workingDir,
		Vars: map[string]interface{}{
			"name":            "labs-compute-vpc",
			"cidr":            "10.0.0.0/16",
			"azs":             azs,
			"public_subnets":  public_subnets,
			"private_subnets": private_subnets,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	privateSubnetId := terraform.Output(t, terraformOptions, "private_subnets")

	vpcId := terraform.Output(t, terraformOptions, "vpc_id")

	subnets := aws.GetSubnetsForVpc(t, vpcId, awsRegion)

	require.Equal(t, 2, len(subnets))

	assert.False(t, aws.IsPublicSubnet(t, privateSubnetId, awsRegion))
}
