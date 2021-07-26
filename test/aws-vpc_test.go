package test_api_storage_infrastructure

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

func TestTerraformAwsVpc(t *testing.T) {
	t.Parallel()

	azs := []string{"ap-southeast-1a", "ap-southeast-1b"}
	public_subnets := []string{"10.0.0.0/19", "10.0.32.0/19"}
	private_subnets := []string{"10.0.128.0/19", "10.0.160.0/19"}

	workingDir := test_structure.CopyTerraformFolderToTemp(t, "../.", ".")
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: workingDir,
		Vars: map[string]interface{}{
			"stage":                "labs",
			"environment":          "compute",
			"name":                 "labs-compute-vpc",
			"cidr":                 "10.0.0.0/16",
			"azs":                  azs,
			"public_subnets":       public_subnets,
			"private_subnets":      private_subnets
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)
}
