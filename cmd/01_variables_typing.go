var (
  name string
  age int
  location string
)

var (
  name, location string
  age int
)


var name string
var location string
var age int

//A var declaration can include initializers, one per variable.
var (
  name string = "Princess Yaruka"
  location string = "São Paulo"
  age int = 32
)


var (
  name = "Princess Yaruka"
  location = "São Paulo"
  age = 32
)

//You can also initialize variables on the same line:
var (
  name, location, age = "Princess Yaruka", "São Paulo", 32
)
