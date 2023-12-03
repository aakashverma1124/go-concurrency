# go-concurrency

### Synchronous Vs Asynchronous in Single Core

![1.png](https://learnwithinnoskrit.s3.amazonaws.com/go-concurrency/1.png)

### Synchronous Vs Asynchronous in Multi Core

![2.png](https://learnwithinnoskrit.s3.amazonaws.com/go-concurrency/2.png)

![3.png](https://learnwithinnoskrit.s3.amazonaws.com/go-concurrency/3.png)

Check [Code Examples](https://github.com/aakashverma1124/go-concurrency/blob/main/introduction)

### Conclusion

|              | Single Core                     | Multi Core              |
|--------------|---------------------------------|-------------------------|
| Synchronous  | Neither Concurrent nor Parallel | Parallel                |
| Asynchronous | Concurrent                      | Concurrent and Parallel |


### Deadlock

Deadlock is a state in which a task starts waiting for something that will not happen. As a result, it stops progressing.

```
process A {
     update variable x
     update variable y
}

process B {
     update variable y
     update variable x
}
```
Suppose two processes (process A and process B) are running simultaneously.
Since both processes are executing, process A will acquire a lock on variable `x`, and process B will acquire a lock on variable `y`.

![deadlock](https://learnwithinnoskrit.s3.amazonaws.com/go-concurrency/deadlock.png)

Check [Deadlock Code Examples](https://github.com/aakashverma1124/go-concurrency/blob/main/deadlock/deadlock.go)
