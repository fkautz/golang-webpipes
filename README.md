Simple Go Webpipes Implementation
=================================

See src/webpipes/wpecho/webpipes_echo.go for the full example

Import Webpipes
---------------
```
  import(
    "webpipes"
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
