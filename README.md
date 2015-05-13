This is Mir, the web indexing and search engine written in Go.
To compile the sprinter, use the go building tool like such (in `cmd/sprinter`)
`go build .`, or use the Makefile by issuing `make sprinter` in the main
directory. To build this you obviously need the Go compiler. Release
packages can be found at [Github](https://github.com/golang/go/release).
Once the Go compiler is installed and the Go path set, you can install this package.

To benchmarks, the sprinter program accepts a `-bench` flag, but it needs
a MongoDB testing backend in the specific location of `../storage/mongo_test`.
This `mongo_test` folder can be found in the `storage` folder, and because these
benchmarks use functions that exist in the `TestDB` source, this location is
hardcoded. When not using this flag, this error won't be encountered. This
means that two extra programs have to be installed, MongoDB and supervisor,
to start the benchmarking.

To print the `sprinter` output, use the `-verbose` flag.

Currently, the `sprinter` command also shows the flags for the `testing` command,
because of importing this in the `main` package. The only two `sprinter` flags
are `-verbose` and `-bench`.

A target in the main Makefile can be found, named `make archive`, that builds
the program and creates a directory structure so that the benchmarking can be run.
This folder is called `sprinter` and resides in `cmd/sprinter`, or you can
unpack it from `./sprinter.tar.gz`.

