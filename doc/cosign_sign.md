## cosign sign

Sign the supplied container image.

### Synopsis

Sign the supplied container image.

```
cosign sign [flags]
```

### Examples

```
  cosign sign --key <key path>|<kms uri> [--payload <path>] [-a key=value] [--upload=true|false] [-f] [-r] <image uri>

  # sign a container image with Google sign-in (experimental)
  COSIGN_EXPERIMENTAL=1 cosign sign <IMAGE>

  # sign a container image with a local key pair file
  cosign sign --key cosign.key <IMAGE>

  # sign a multi-arch container image AND all referenced, discrete images
  cosign sign --key cosign.key --r <MULTI-ARCH IMAGE>

  # sign a container image and add annotations
  cosign sign --key cosign.key -a key1=value1 -a key2=value2 <IMAGE>

  # sign a container image with a key pair stored in Azure Key Vault
  cosign sign --key azurekms://[VAULT_NAME][VAULT_URI]/[KEY] <IMAGE>

  # sign a container image with a key pair stored in AWS KMS
  cosign sign --key awskms://[ENDPOINT]/[ID/ALIAS/ARN] <IMAGE>

  # sign a container image with a key pair stored in Google Cloud KMS
  cosign sign --key gcpkms://projects/[PROJECT]/locations/global/keyRings/[KEYRING]/cryptoKeys/[KEY]/versions/[VERSION] <IMAGE>

  # sign a container image with a key pair stored in Hashicorp Vault
  cosign sign --key hashivault://[KEY] <IMAGE>

  # sign a container image with a key pair stored in a Kubernetes secret
  cosign sign --key k8s://[NAMESPACE]/[KEY] <IMAGE>

  # sign a container in a registry which does not fully support OCI media types
  COSIGN_DOCKER_MEDIA_TYPES=1 cosign sign --key cosign.key legacy-registry.example.com/my/image
```

### Options

```
      --allow-insecure-registry     whether to allow insecure connections to registries. Don't use this for anything but testing
  -a, --annotations strings         extra key=value pairs to sign
      --attachment string           related image attachment to sign (sbom), default none
      --cert string                 path to the x509 certificate to include in the Signature
  -f, --force                       skip warnings and confirmations
      --fulcio-url string           [EXPERIMENTAL] address of sigstore PKI server (default "https://fulcio.sigstore.dev")
  -h, --help                        help for sign
      --identity-token string       [EXPERIMENTAL] identity token to use for certificate from fulcio
      --key string                  path to the private key file, KMS URI or Kubernetes Secret
      --oidc-client-id string       [EXPERIMENTAL] OIDC client ID for application (default "sigstore")
      --oidc-client-secret string   [EXPERIMENTAL] OIDC client secret for application
      --oidc-issuer string          [EXPERIMENTAL] OIDC provider to be used to issue ID token (default "https://oauth2.sigstore.dev/auth")
      --payload string              path to a payload file to use rather than generating one
  -r, --recursive                   if a multi-arch image is specified, additionally sign each discrete image
      --rekor-url string            [EXPERIMENTAL] address of rekor STL server (default "https://rekor.sigstore.dev")
      --sk                          whether to use a hardware security key
      --slot string                 security key slot to use for generated key (default: signature) (authentication|signature|card-authentication|key-management)
      --upload                      whether to upload the signature (default true)
```

### Options inherited from parent commands

```
      --output-file string   log output to a file
  -d, --verbose              log debug output
```

### SEE ALSO

* [cosign](cosign.md)	 - 

