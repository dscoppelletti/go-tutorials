module example.com/hello

go 1.18

// I need the following only because I have detached the module from its
// original repository
replace example.com/example => ../example

require example.com/example v0.0.0

