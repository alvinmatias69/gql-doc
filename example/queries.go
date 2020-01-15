package gqlDocumentation

var query = `
	getImages(id:[Int!]!): Images!
`

var queryType = `
# Query to get Affiliate Explore Page
type Images {
	data: [String]
}

`
