# CGO examples - Shared object


1. Build the shared library with
```
gcc -shared -o libhello.so hello.co
```
There is already a compiled version ready to use in gohello/libhello

2. Build the Go binary, note the LDFLAGS in gohello.go, also the library name is `libhello` while the
passed flag is only `-lgello` 

```
go build main.go
```
at this point we can build the binary, but when we try to run it we will get an error.

3. tell the binary where to find the shared object
```
LD_LIBRARY_PATH=$PWD/gohello/libhello:$LD_LIBRARY_PATH ./main
```
or suing go run

```
LD_LIBRARY_PATH=$PWD/gohello/libhello:$LD_LIBRARY_PATH go run main.go
```