module hello

go 1.15

// For production use, you’d publish your modules on a server, either inside
// your company or on the internet, and the Go command will download them from
// there. For now, you need to adapt the caller's module so it can find the
// greetings code on your local file system.

// Tells Go to replace the module path (the URL example.com/greetings) with a
// path you specify. In this case, that's a greetings directory next to the
// hello directory.
replace example.com/greetings => ../greetings

// The command "go build" locates the module greetings and adds it as a
// dependency. The command also creates the executable file "hello".
require example.com/greetings v0.0.0-00010101000000-000000000000
