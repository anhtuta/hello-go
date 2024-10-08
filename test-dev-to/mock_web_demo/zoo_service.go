package mock_web_demo

import (
	"errors"
	"regexp"
)

var errFactNotFound = errors.New("fact not found")

var favSnackMatcher = regexp.MustCompile(
	"favorite snack is (.*)",
)

// getSlothsFavoriteSnack calls the zoo API to get facts about sloths, then searches and returns sloths' favorite snack.
//
// Nếu dùng concrete parameter type *ZooHTTPClient thì khi test sẽ không thể mock được:
// func getSlothsFavoriteSnack(c *ZooHTTPClient) (string, error) {
//
// Nếu dùng con trỏ của interface *ZooClient, thì khi gọi phải dereference con trỏ:
//
//	func getSlothsFavoriteSnack(c *ZooClient) {
//	    res, err := (*c).ListAnimalFacts(...)
//
// Note: you have to use (*c).ListAnimalFacts instead of c.ListAnimalFacts is because c is
// a pointer to an interface, not an interface itself. In Go, you cannot call methods on
// a pointer to an interface directly; you need to dereference the pointer first.
//
// Nếu chỉ dùng kiểu interface ZooClient, thì ta không thể modify field của biến `c`, nhưng do interface không có field
// gì (immutable), nên ta không cần modify gì cả, do đó không cần dùng con trỏ!
// Nhưng nếu có 1 kiểu struct implement interface đó dùng pointer receiver, thì ta vẫn có thể modify field của struct
// đó, mặc dù param c có kiểu interface, không phải con trỏ của interface/struct đó.
// But the underlying concrete type that the interface refers to can still be modified if it has pointer receivers.
func getSlothsFavoriteSnack(c ZooClient) (string, error) {
	// Since ListAnimalFacts talks to the server for the zoo's API, the code is nondeterministic.
	// So testing it in a repeatable way is a great use case for Testify Mock.
	res, err := c.ListAnimalFacts(AnimalFactsQuery{
		AnimalName: "sloth",
		PageToken:  "ALL", // just want to fetch all the facts for simplicity
	})
	if err != nil {
		return "", err
	}

	// check if any facts match the "favorite snack is" regex and if so, return the match
	for _, f := range res.Facts {
		match := favSnackMatcher.FindStringSubmatch(f)
		if len(match) < 2 {
			continue
		}
		return match[1], nil
	}

	// otherwise if the fact about sloths' favorite snack
	// isn't in the zoo API, return errFactNotFound.
	return "", errFactNotFound
}
