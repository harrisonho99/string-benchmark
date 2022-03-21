# *Benchmark String concatenation*

  ## Normal concat
  ```go
      var s String
      s +="hello" 
  ```
  Count: **1000** times

  Took: **26.481084ms**


  ## Use ```strings.Builder``` from standard library
  ```go
      var s strings.Builder
      s.WriteString("hello") 
  ```
  Count: **1000** times

  Took: **432.959µs**
  
  Compare: **61.2** times faster than **normal concat**

  ## Use ```bytes.Buffer``` from standard library
  ```go
      var s bytes.Buffer
      s.WriteString("hello") 
  ```
  Count: **1000** times

  Took: **460.25µs**

  Compare: **0.94** times slower than **strings.Builder**