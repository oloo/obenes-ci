variable "access_key" {}
variable "secret_key" {}

provider "aws" {
  access_key = "${var.access_key}"
  secret_key = "${var.secret_key}"
  region = "us-east-1"
}

resource "aws_instance" "obene-concourse-ci" {
  ami = "ami-b63769a1"
  instance_type = "t2.small"
}