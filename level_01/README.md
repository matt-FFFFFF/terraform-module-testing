# level_01

## Description

These exercises are designed to get you familiar with the fluent Terraform testing module.
It will require some skills in Go, but the exercises are designed to be easy to follow.

In these early examples the boilerplate code has been provided for you.

You will need to edit the [module_test.go](./test/module_test.go) file.

## Task 1 - TestSimple

Modify the TestSimple function to check the input attribute of the terraform_data.example resource is equal to "example".

## Task 2 - TestForEach

Modify the TestForEach function to check the input attribute of the terraform_data.example_for_each resources is equal to the data in the map.

## Task 3 - TestCondition

Add checks that the conditional resource is present and that the attribute is correct.
Also add checks that the resource is not present when the condition is false.
