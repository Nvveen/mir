This is the Mir, the web indexing and search engine written in Go.
To compile the sprinter, use the go building tool like such (in `cmd/sprinter`)
`go build .`.

To benchmarks, the sprinter program accepts a `-bench` flag, but it needs
a MongoDB testing backend in the specific location of `../storage/mongo_test`.
This `mongo_test` folder can be found in the `storage` folder, and because these
benchmarks use functions that exist in the `TestDB` source, this location is
hardcoded. When not using this flag, this error won't be encountered.

To print the `sprinter` output, use the `-verbose` flag.

Currently, the `sprinter` command also shows the flags for the `testing` command,
because of importing this in the `main` package. The only two `sprinter` flags
are `-verbose` and `-bench`.
