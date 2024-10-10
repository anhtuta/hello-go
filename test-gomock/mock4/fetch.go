package mock4

type Fetcher interface {
	FetchData() ([]byte, error)
}

// MyFunc is the function under test that uses the Fetcher interface
// MyFunc là hàm sẽ gọi tới method FetchData của Fetcher interface để lấy dữ liệu từ bên ngoài,
// ta sẽ viết UT cho hàm này và stub hàm FetchData.
func MyFunc(fetcher Fetcher) ([]byte, error) {
	return fetcher.FetchData()
}
