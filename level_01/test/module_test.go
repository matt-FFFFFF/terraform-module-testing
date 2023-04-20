package test

import (
	"testing"

	"github.com/Azure/terratest-terraform-fluent/check"
	"github.com/Azure/terratest-terraform-fluent/setuptest"
	"github.com/stretchr/testify/require"
)

const (
	moduleDir = "../"
)

func TestSimple(t *testing.T) {
	t.Parallel()

	// Setup the test
	test, err := setuptest.Dirs(moduleDir, "").WithVars(nil).InitPlanShow(t)
	require.NoError(t, err)
	defer test.Cleanup()

	// Check the results
	check.InPlan(test.Plan).NumberOfResourcesEquals(1).ErrorIsNil(t)
	// How can we check that the input attribute of the terraform_data.example resource is equal to "example"?
	// hint: see the github.com/Azure/terratest-terraform-fluent README...

	t.Fail() // remove this once you have done!
}

func TestCondition(t *testing.T) {
	t.Parallel()

	// Define module variables
	// See https://go.dev/tour/methods/14 for a description of the any value (aka empty interface).
	// This is a good article: https://bitfieldconsulting.com/golang/map-string-interface.

	// First define a map of strings to any value, these will contain the terraform variables passed to the module
	// The map key (string) is the name of the variable, the value is the value of the variable.
	// The value can be any type.
	tfvars := map[string]any{
		"example_condition": true,
	}

	// We are setting up two tests, one where the condition is true and one where it is false.
	// We do this using subtests, created by the t.Run() function.
	// This function takes two values: a name for the test and a function to run the test.

	// The first test, where the condition is true.
	t.Run("present", func(t *testing.T) {
		// Setup the test using the Terraform variables defined above
		test, err := setuptest.Dirs(moduleDir, "").WithVars(tfvars).InitPlanShow(t)
		require.NoError(t, err)
		defer test.Cleanup()

		// Check that the resource is present in the present test and that the input attribute has the value "example_condition"...

		t.Fail() // remove this once you have done!
	})

	// This is the second test, where no variables are supplied and the condition is false due to the default value of the variable.
	t.Run("notPresent", func(t *testing.T) {
		// Setup the test using the Terraform variables defined above
		test, err := setuptest.Dirs(moduleDir, "").WithVars(nil).InitPlanShow(t)
		require.NoError(t, err)
		defer test.Cleanup()

		// Check that the resource is NOT present in the notPresent test...

		t.Fail() // remove this once you have done!
	})
}

func TestForEach(t *testing.T) {
	t.Parallel()

	// Define module variables
	// See https://go.dev/tour/methods/14 for a description of the any value (aka empty interface).
	// This is a good article: https://bitfieldconsulting.com/golang/map-string-interface.
	vars := map[string]any{
		"example_for_each": map[string]string{
			"instance_1": "data_1",
			"instance_2": "data_2",
		},
	}

	// Setup the test
	test, err := setuptest.Dirs(moduleDir, "").WithVars(vars).InitPlanShow(t)
	require.NoError(t, err)
	defer test.Cleanup()

	// Check the results
	check.InPlan(test.Plan).NumberOfResourcesEquals(3).ErrorIsNil(t)

	// How can we check the values of the terraform_data.example_for_each resources?
	// hint: use a loop...

	t.Fail() // remove this once you have done!
}
