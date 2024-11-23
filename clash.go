package goclash

// New creates a new clash client, using the provided credentials.
func New(creds Credentials) (*Client, error) {
	return newClient(creds)
}

func NewWithKeys(keys ...string) (*Client, error) {
	return newKeysClient(keys)
}
