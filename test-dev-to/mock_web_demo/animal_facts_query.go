package mock_web_demo

// File này giống Model/DTO bên Java

// AnimalFactsQuery is a query to get animal facts
type AnimalFactsQuery struct {
	AnimalName string
	PageToken  string
}
