package id

// ID represents a unique app identifier with a domain and name.
type ID struct {
	Domain string
	Name   string
}

// New creates a new ID instance with the given domain and name.
func New(domain, name string) *ID {
	return &ID{Domain: domain, Name: name}
}

// String returns the string representation of the ID in the format "domain.name".
func (id *ID) String() string {
	return id.Domain + "." + id.Name
}
