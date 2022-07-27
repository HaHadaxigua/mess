package enum

//go:generate go run github.com/abice/go-enum -f=$GOFILE --marshal

// Directory
// ENUM(
//	ASC,
//      DESC,
//)
type Directory int
