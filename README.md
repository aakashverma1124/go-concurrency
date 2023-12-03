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

Check [Code Examples](https://github.com/aakashverma1124/go-concurrency/blob/main/deadlock/deadlock.go)
