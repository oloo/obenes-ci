variable "access_key" {}
variable "secret_key" {}
variable "default_availability_zone" {
  default = "us-east-1d"
}

variable "vpc" {
  default = "vpc-79e3bd1d"
}

variable "subnet" {
  default = "subnet-25f72d7d"
}

provider "aws" {
  access_key = "${var.access_key}"
  secret_key = "${var.secret_key}"
  region = "us-east-1"
}

resource "aws_instance" "obene-concourse-ci-instance" {
  ami = "ami-b63769a1"
  instance_type = "t2.small"
  availability_zone = "${var.default_availability_zone}"
  tags {
    Name = "Obene Concourse CI"
    Environment = "production"
    Group = "Obene"
  }
}

resource "aws_elb" "obene-concourse-elb" {
  name = "obene-concourse-elb"
  internal = false
  availability_zones = ["${var.default_availability_zone}"]
  instances = ["${aws_instance.obene-concourse-ci-instance.id}"]

  listener {
    instance_port = 8080
    instance_protocol = "http"
    lb_port = 80
    lb_protocol = "http"
  }

  tags {
    Name = "Obene Concourse CI ELB"
    Environment = "production"
    Group = "Obene"
  }
}