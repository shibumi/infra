package internal

// BoolPtr needs a bool and returns a pointer to the bool.
// This function is needed for pulumi's cloud-config.
// Pulumi's cloud-config does not seem to support pulumi.Bool or pulumi.BoolPtr :(
func BoolPtr(b bool) *bool {
	return &b
}

// StringPtr needs a string and returns a pointer to the string.
// This function is needed for pulumi's cloud-config.
// Pulumi's cloud-config does not seem to support pulumi.String or pulumi.StringPtr :(
func StringPtr(s string) *string {
	return &s
}
