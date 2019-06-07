package main

// High-level IR for the API spec being described.
type hir struct {
	topics []topic
}

// An API topic being described.
type topic struct {
	models []apiModel
}

type visibility int

const (
	visibilityPrivate visibility = iota + 1
	visibilityPublic
)

// A model used by the APIs.
type apiModel struct {
	ident  string
	doc    string
	vis    visibility
	fields []apiModelField

	// TODO: retain source order
	// map[languageTag][]snippet
	inlineCodeSections map[string][]string
}

type apiModelField struct {
	ident string
	doc   string
	typ   string
	vis   visibility
	tags  map[string]string
}
