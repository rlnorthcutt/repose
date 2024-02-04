
- Goal and plan
- notes on OOP structure and organizational approach (inlcuding "wireframes")
- notes on comments and full variable names, and public vs private organization

1) setup
2) create placeholder objects with test output
3) create main.go with test for each object
4) create testgo shell script
5) test main and then build and test

6) Create logger.go for nice output - we will use this as a standard way to interact with the user
7) Create command.go with placeholder methods for each command
8) Create main.go controller with commands
9) test and build/test

10) create filesystem.go for create/delete (including check path and private methods)
11) create a test flie for filesystem & run tests (we will only o this for more complex objects)
12) update "new command" to use filesystem



Config system improved using this guide: https://dev.to/koddr/let-s-write-config-for-your-golang-web-app-on-right-way-yaml-5ggp

Using custom yml parser instead of yaml library
 - cut 0.8M (3.3 -> 2.5)

Build with thee flags to reduce size:https://gophercoding.com/reduce-go-binary-size/
also https://github.com/xaionaro/documentation/blob/master/golang/reduce-binary-size.md
go build -ldflags="-s -w"
- cut 0.8M (2.5 -> 1.7)