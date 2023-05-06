This example was created for my blog post ["Go interfaces bypass cyclic dependencies"](https://www.ozoniuss.com/me/blog/go-interfaces-bypass-cyclic-dependencies/).

Run all code using 

```
go run main.go
```

which will print "hello client" in various formats.

The idea is that inside client you can define a `helloPrinter` interface, which can both be implemented inside the client package (see `ReversePrinter`), but also by the external `color` package. Then, you could create functions in the `client` package that make use of the `color` implementation of `PrintHello`, such as `PrintHelloRed`. This is possible because in go interfaces are implemented implicitly, meaning there is no need for `color` to have a dependency of `client`, which in turn would lead to a cyclic dependency since `client` also uses `color`.

This is more of a fun fact than a real world example. Check out my blog post for more details.