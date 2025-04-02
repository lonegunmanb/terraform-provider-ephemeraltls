# This example creates a ephemeral TLS private key.

ephemeral "ephemeraltls_private_key" "test" {
  algorithm = "RSA"
  rsa_bits  = 4096
}
