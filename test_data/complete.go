package test_data

var query = `
  #
  # register new user
  #
  register(name: String!): Response!
`

var mutationType = `
  # Response given after register new user
  type Response {
          # success indicate register result
	  success: Boolean!

          # error defines error message on failed register
	  error: String
  }
`
