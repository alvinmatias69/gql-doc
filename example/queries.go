package gqlDocumentation

var query = `
	getImages(id:[Int!]!): Images!
`

var queryType = `
# Query to get user images by ID
type Images {
	data: [String]
}

`
