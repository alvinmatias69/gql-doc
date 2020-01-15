package test_data

var query = `
  register(name: String!): Response!
`

var mutationType = `
  type Response {
	  success: Boolean!
	  error: String
  }
`
