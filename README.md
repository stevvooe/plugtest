# plugtest

Minimal example of using go 1.8 plugins with init registration.

# Usage

First, build the plugin:

```console
$ go build -buildmode=plugin -o plugin1.so ./plugin1
```

Then, build `plugger`:

```console
$ go build ./cmd/plugger
```

These don't need to be done in a particular order. Now, we can run plugger:

```console
$ ./plugger plugin1
```

Let's build the other plugin:

```console
$ go build -buildmode=plugin -o plugin2.so ./plugin2
```

Now, we can run that one as well:

```console
$ ./plugger plugin2
```
