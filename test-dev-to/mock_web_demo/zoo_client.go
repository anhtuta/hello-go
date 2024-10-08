package mock_web_demo

// File này giống như interface bên Java vậy

// 1. Wrapping our client in an interface
// Do function ListAnimalFacts gọi tới external service --> nó là nondeterministic code --> cần mock.
// Muốn mock nó thì phải wrap nó trong 1 interface, bên code sẽ implement nó để gọi tới external service,
// bên test sẽ implement interface đó để mock.
type ZooClient interface {
	ListAnimalFacts(q AnimalFactsQuery) (*AnimalFactsResponse, error)
}
