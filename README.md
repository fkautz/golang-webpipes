Simple Go Webpipes Implementation
=================================

See https://github.com/fkautz/golang-webpipes-echo for a full example

Import Webpipes
---------------
```
  import(
    "github.com/fkautz/golang-webpipes"
  )
```

Define your function
--------------------

```
  func EchoPipe(inputs map[string]string) (map[string]string, error) {
    return inputs, nil
  }
```

Create a block definition
-------------------------
```
  block := webpipes.Block{
    "echo",
    "/echo",
    "Echo Service",
    []webpipes.InputParameter{
      webpipes.InputParameter{
        "input",
        "string",
        "input stream to echo",
        false,
        "",
      },
    },
    []webpipes.OutputParameter{
      webpipes.OutputParameter{
        "output",
        "string",
        "echoed string",
      },
    },
  }
```

Define a new pipe and run it
----------------------------
```
  pipe := webpipes.GoWebPipe{block, EchoPipe}
  http.Handle("/echo", pipe)
  err := http.ListenAndServe(":8080", nil)
```
