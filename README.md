# terraform-and-terratest-example

This project just create and destroy a simples S3 bucket using terraform and using terratest to test the code.

Prerequisites
--------

Create your account in AWS and user(IAM) and get the secret and token keys

[IAM Credentials](https://docs.aws.amazon.com/cli/latest/reference/iam/create-access-key.html)

## Install terraform, GO and Terratest

[TERRAFORM](https://learn.hashicorp.com/terraform/getting-started/install.html)

[GO](https://golang.org/doc/install)

[TERRATEST](https://github.com/gruntwork-io/terratest)


To using this project with terraform you will need create a file `params.tfvars` as below and put your AWS keys and region.
If you like change the bucket values too.

```terraform
#Providers Values
aws_region= "SOME_AWS_REGION"
aws_access_key = "IAM_ACCESS_KEY"
aws_secret_key = "IAM_SECRET_KEY"

#Bucket Values
bucket_name="my-first-bucket-with-terraform"
tag_bucket_name ="terraform-bucket"
tag_bucket_environment="test"

```

RUN TERRAFORM
--------

To validade this code with terraform run a PLAN command in your terminal

`terraform plan --var-file=params.tfvars`

To create the resources run a terraform APPLY

`terraform apply --var-file=params.tfvars`

*this command will create the resources at your account and can result in charges on your card

To remove the resources run a terraform DESTROY

`terraform destroy --var-file=params.tfvars`

RUN TERRATEST
--------

Enter in the test folder and call test from GO
```GO
cd test/

go test
```
