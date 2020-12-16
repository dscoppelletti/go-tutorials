// Initialized by:
// go mod init hello
module hello

go 1.15

// When your code imports packages from another module, a go.mod file lists the
// specific modules and versions providing those packages. That file stays with
// your code, including in your source code repository.
// But before it ran the code, go run located and downloaded the rsc.io/quote
// module that contains the package you imported. By default, it downloaded the
// latest version -- v1.5.2. Go build commands are designed to locate the
// modules required for packages you import.
require rsc.io/quote v1.5.2 // indirect

// In addition to go.mod, the go command maintains a file named go.sum
// containing the expected cryptographic hashes of the content of specific
// module versions.
// The go command uses the go.sum file to ensure that future downloads of these
// modules retrieve the same bits as the first download, to ensure the modules
// your project depends on do not change unexpectedly, whether for malicious,
// accidental, or other reasons. Both go.mod and go.sum should be checked into
// version control.
