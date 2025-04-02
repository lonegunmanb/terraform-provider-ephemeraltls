# ECDSA key with P384 elliptic curve
ephemeral "ephemeraltls_private_key" "ecdsa-p384-example" {
  algorithm   = "ECDSA"
  ecdsa_curve = "P384"
}

# RSA key of size 4096 bits
ephemeral "ephemeraltls_private_key" "rsa-4096-example" {
  algorithm = "RSA"
  rsa_bits  = 4096
}

# ED25519 key
ephemeral "ephemeraltls_private_key" "ed25519-example" {
  algorithm = "ED25519"
}
