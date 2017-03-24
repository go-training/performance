# Go Performance

Compare the performance of Node, Go and Python

* Node version: **7.7.4**
* Python version: **2.7.0**
* Go version: **1.8.0**

## Loop/Arithmetic

See the [loop source code](./loop)

```
$ time node loop/main.js && time go run loop/main.go && time python2.7 loop/main.py

real    0m1.081s
user    0m1.066s
sys     0m0.012s

real    0m0.476s
user    0m0.426s
sys     0m0.055s

real    0m8.822s
user    0m8.551s
sys     0m0.045s
```
