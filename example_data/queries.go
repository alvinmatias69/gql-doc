package gqlDocumentation

var query = `
	getImages(id:[Int]!): Images!
`

var queryType = `
# Query to get user images by ID
type Images {
	data: [String]
}

# Data of given image
input Data {
	# Image of given data
	image: String!
}

# Compression enumeration
enum Compression {
	# Not cold
	HOT
	# Not hot
	COLD
}

# ImageQuery type
union ImageQuery = Min | Meta | Complete
`
