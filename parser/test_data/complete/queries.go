package test_complete

var queries = `
# get profile by name
getProfile(name: String!): Profile
`

var queryType = `
# response data
type Response {
	name: String
	age: Int!
	groupID: [Int]
}

# input data
input Input {
	name: String!
}

# response type variant
enum ResponseType {
	SUCCESS
	FAILURE
}

# some interface
interface SomeItf {
	name: String!
}

# XYZ dragon union
union Dragon = X | Y | Z
`
