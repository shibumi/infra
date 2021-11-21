package internal

// SSHPublicKey extends the SSH public key with its ID (comment field)
// This makes handling easier. We just get the key from the pulumi configuration.
// An alternative is parsing the key and reading the comment field.
type SSHPublicKey struct {
	ID        string
	PublicKey string
}
