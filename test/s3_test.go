package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	//"github.com/stretchr/testify/assert"
)

func TestTerraformAwsS3Example(t *testing.T) {
	t.Parallel()

	// Give this S3 Bucket a unique ID for a name tag so we can distinguish it from any other Buckets provisioned
	// in your AWS account
	expectedName := fmt.Sprintf("my-first-bucket-with-terraform-%s", strings.ToLower(random.UniqueId()))

	//Put a tag in yout bucket
	expectedTagName := "site bucket"

	// Give this S3 Bucket an environment to operate as a part of for the purposes of resource tagging
	expectedTagEnvironment := "DEV"

	// AWS region to test in.
	awsRegion := "sa-east-1"



	
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"bucket_name":            expectedName,
			"tag_bucket_name":        expectedTagName,
			"tag_bucket_environment": expectedTagEnvironment,
		},

		// Environment variables to set when running Terraform
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},

		// Variables to pass to our Terraform code using -var-file= options
		VarFiles: []string{"params.tfvars"},
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of an output variable
	//bucketID := terraform.Output(t, terraformOptions, "bucket_id")
	bucketID := terraform.Output(t, terraformOptions, "bucket_name")

	//Juat print the bucket name
	fmt.Println("bkt ", bucketID)

	// Verify that our Bucket has versioning enabled
	//assert.Equal(t, expectedName, bucketID)
	aws.AssertS3BucketExists(t, awsRegion, bucketID)

}
