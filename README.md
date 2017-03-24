# Go Performance

Compare the performance of [Node](https://nodejs.org), [Go](https://golang.org) and [Python](https://www.python.org/)

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

## I/O

See the [io source code](./io)

```
$ time node io/main.js && time go run io/main.go && time python2.7 io/main.py

real    0m1.872s
user    0m1.595s
sys     0m0.233s

real    0m1.667s
user    0m0.490s
sys     0m1.093s

real    0m0.636s
user    0m0.580s
sys     0m0.040s
```

## Bubble Sort

See the [bubble source code](./bubble)

```
$ time node bubble/main.js && time go run bubble/main.go && time python2.7 bubble/main.py

real    0m2.133s
user    0m2.061s
sys     0m0.022s

real    0m1.513s
user    0m1.163s
sys     0m0.186s

real    0m14.118s
user    0m13.765s
sys     0m0.203s
```
