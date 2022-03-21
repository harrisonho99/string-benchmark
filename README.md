# *Benchmark String concatenation*

  ## Normal concat
  ```go
      var s String
      s +="hello" 
  ```
  Loop count: **10000** times

  Allocation: **3305** times

  Took: **26.481084ms**


  ## Use ```strings.Builder``` from standard library
  ```go
      var s strings.Builder
      s.WriteString("hello") 
  ```
 Loop count: **10000** times

  Allocation: **20** times

  Took: **432.959µs**
  
  Compare: **61.2** times faster than **normal concat**

  ## Use ```bytes.Buffer``` from standard library
  ```go
      var s bytes.Buffer
      s.WriteString("hello") 
  ```
Loop count: **10000** times

  Allocation: **11** times

  Took: **460.25µs**

  Compare: **0.94** times slower than **strings.Builder**


  **Test machine** :`` Go: 1.18, Macbook pro 14, Arch: ARM-64, MacOS: 12.3, ``