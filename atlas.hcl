// https://atlasgo.io/atlas-schema/projects
variable "user" {
  type    = string
  default = "root"
}
variable "pass" {
  type    = string
  default = "root"
}
variable "addr" {
  type    = string
  default = "127.0.0.1:3311"
}

locals {
  base_url = "mysql://${var.user}:${var.pass}@${var.addr}"
  example_database  = "${local.base_url}/example_database"
}

env "local" {
  // Declare where the schema definition resides.
  // Also supported: ["file://multi.hcl", "file://schema.hcl"].
  src = fileset("schema/*.hcl")

  // Define the URL of the database which is managed
  // in this environment.
  url = local.example_database

  // Define the URL of the Dev Database for this environment
  // See: https://atlasgo.io/concepts/dev-database
  // dev = "docker://mysql/8/dev"
}

env "dev" {
  src = ileset("schema/*.hcl")
  url = getenv("DATABASE_URL")
}

