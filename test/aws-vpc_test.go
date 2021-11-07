package test

import (
	"io"
	"log"
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTerraformAwsVpc(t *testing.T) {
	t.Parallel()

	awsRegion := "ap-southeast-1"
	azs := []string{"ap-southeast-1a", "ap-southeast-1b"}
	public_subnets := []string{"10.0.0.0/19"}
	private_subnets := []string{"10.0.128.0/19"}

	moduleFolder := "aws-vpc"
	workingDir := test_structure.CopyTerraformFolderToTemp(t, "..", moduleFolder)

	// Copy a dummy AWS provider to suppress 'region' warnings
	src := "aws-provider.tf"
	log.Printf("stat-ing source file: %s", src)
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		log.Printf("error accessing file: %s, failing test NOW", src)
		t.FailNow()
	}

	log.Printf("checking if %s is regular file", src)
	if !sourceFileStat.Mode().IsRegular() {
		log.Printf("%s is not a regular file, failing test NOW", src)
		t.FailNow()
	}

	log.Printf("opening source file: %s", src)
	source, err := os.Open(src)
	if err != nil {
		log.Printf("error opening file: %s, failing test NOW", src)
		t.FailNow()
	}
	defer source.Close()

	destFile := workingDir + "/" + src
	log.Printf("creating destination file: %s", destFile)
	destination, err := os.Create(destFile)
	if err != nil {
		log.Printf("error creating destination: %s, failing test NOW", destFile)
		t.FailNow()
	}
	defer destination.Close()

	log.Printf("copying file from source: %s to destination: %s", src, destFile)
	_, err = io.Copy(destination, source)
	if err != nil {
		log.Printf("error copying source: %s to destination: %s, failing test NOW", src, destFile)
		t.FailNow()
	}

	logger.Logf(t, "path to test folder %s\n", workingDir)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: workingDir,
		Vars: map[string]interface{}{
			"name":            "labs-compute-vpc",
			"cidr":            "10.0.0.0/16",
			"azs":             azs,
			"public_subnets":  public_subnets,
			"private_subnets": private_subnets,
			"enable_flow_log": false,
			"flow_log_cloudwatch_log_group_retention_in_days": 0,
		},
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	privateSubnetId := terraform.Output(t, terraformOptions, "private_subnets")
	vpcId := terraform.Output(t, terraformOptions, "vpc_id")
	vpc := aws.GetVpcById(t, vpcId, awsRegion)
	subnets := vpc.Subnets

	require.Equal(t, 2, len(subnets))
	assert.False(t, aws.IsPublicSubnet(t, privateSubnetId, awsRegion))
}
