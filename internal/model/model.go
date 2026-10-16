package model

// NormalizedSpec represents the simplified OpenAPI document for comparison.
type NormalizedSpec struct {
	Version   string
	Endpoints map[string]*Endpoint // Keyed by "METHOD /path" e.g., "GET /v1/users"
}

// Endpoint represents a single API operation.
type Endpoint struct {
	Method      string
	Path        string
	Parameters  map[string]*Parameter // Keyed by "in:name" e.g., "query:id"
	RequestBody *Schema
	Responses   map[string]*Response // Keyed by status code e.g., "200"
}

// Parameter represents a request parameter.
type Parameter struct {
	Name     string
	In       string // query, header, path, cookie
	Required bool
	Schema   *Schema
}

// Response represents an API response.
type Response struct {
	StatusCode string
	Content    *Schema
}

// Schema is a simplified representation of an OpenAPI Schema.
type Schema struct {
	Type       string
	Properties map[string]*Schema
	Required   []string
	Enums      []string
	// NOTE: MVP তে আমরা complex polymorphism (oneOf, anyOf) বা unresolved external refs
	// পেলে সেগুলোকে 'undetermined' হিসেবে ফ্ল্যাগ করব, তাই এখানে সেগুলো রাখছি না।
}
