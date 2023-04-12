terraform {
  required_version = ">= 1.4.0, < 2.0.0"
}

resource "terraform_data" "example" {
  input = "example"
}

resource "terraform_data" "example_for_each" {
  for_each = var.example_for_each
  input    = each.value
}
