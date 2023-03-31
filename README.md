# Go Unit Test and Fuzz

In this talk, we will begin by introducing Collatz problem that we can start write or functions and tests. Then move on to others Go testing techniques, such as table-driven tests, activating the race detector in tests, and writing testable examples. Finally, we will cover how to use Go benchmarks to test performance and how to incorporate random inputs in tests with Go fuzz.

[This Talk Slides](https://docs.google.com/presentation/d/1VZQsDjs_zEpofQwAGpBORSVySciiTpSK9Q-_dxLQ-Ik/edit?usp=sharing)

## Table of Contents

- Introduction to [go functions and tests](https://github.com/andref5/gotesttalk/tree/go-functions-and-tests) with Collatz Problem.

    The Collatz problem is a simple mathematical problem that involves iterating through a series of calculations to reach the number 1. In this talk, we will use the Collatz problem as a way to illustrate basic testing techniques in Go.

- **\o_ The particular test technique of Go [table driven](https://github.com/andref5/gotesttalk/tree/table-driven).**

    Table-driven tests are a way of testing a function by defining a table of input and expected output values. This technique is particularly useful when you need to test a function with a large number of inputs or when you need to test a function with complex input and output types.

- Activating the [race detector](https://github.com/andref5/gotesttalk/tree/race-detector) in Your Tests.

    The Go language has a built-in race detector that can help you identify and debug race conditions in your code. By activating the race detector in your tests, you can ensure that your code is thread-safe and free from race conditions.

- Checking performance like a test with [go benchmarks](https://github.com/andref5/gotesttalk/tree/go-benchmarks).

    Go benchmarks are a way of measuring the performance of your code in a controlled environment. By writing benchmarks for your code, you can ensure that your code is performant and that it meets the requirements of your application.

- Writing [testable examples](https://github.com/andref5/gotesttalk/tree/testable-examples). to your documentation.

    Testable examples are a way of incorporating examples into your code documentation that can be used as both documentation and test cases. By using testable examples, you can ensure that your code examples are always up-to-date and that they provide accurate documentation of your code.

- Random inputs in your tests with [go fuzz](https://github.com/andref5/gotesttalk/tree/go-fuzz).

    Go fuzz is a tool that allows you to generate random inputs for your tests. By using Go fuzz, you can ensure that your code can handle unexpected inputs and that it is resilient to edge cases.


## References

- [The Go Programming Language (Brian W. Kernighan, Alan Donovan)](https://www.amazon.com.br/Go-Programming-Language-Brian-Kernighan/dp/0134190440)
- [100 Go Mistakes and How to Avoid Them (Teiva Harsanyi)](https://www.amazon.com.br/100-Go-Mistakes-Avoid-Them/dp/1617299596)
- https://pkg.go.dev/testing
- [Where did the idea for collatz problem come from](https://github.com/ServiceWeaver/weaver/tree/main/examples/collatz)