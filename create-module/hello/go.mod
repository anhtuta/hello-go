module hello-go/hello

go 1.21.6

require (
	hello-go/greetings v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

// For production use, you’d publish the hello-go/greetings module from its repository, where Go
// tools could find it to download it. For now, because you haven't published the module yet,
// you need to adapt the hello-go/hello module so it can find the hello-go/greetings code on
// your local file system. To do that, use the go mod edit command:
// go mod edit -replace hello-go/greetings=../greetings
// Then it will generate this line for you:
replace hello-go/greetings => ../greetings
