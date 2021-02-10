# Testing

For running automated unit tests in Go we need `testing` package. It is intended to be used in concert with the "go test" command, which automates execution of any function of the form  

### Example:
```
    func TestXxx(*testing.T)
```

where Xxx does not start with a lowercase letter. The function name serves to identify the test routine.

Within these functions, use the Error, Fail or related methods to signal failure.

To write a new test suite, create a file whose name ends _test.go that contains the TestXxx functions as described here. Put the file in the same package as the one being tested. The file will be excluded from regular package builds but will be included when the "go test" command is run. For more detail, run "go help test" and "go help testflag".

A simple test function looks like this:

```
func TestAbs(t *testing.T) {
    got := Abs(-1)
    if got != 1 {
        t.Errorf("Abs(-1) = %d; want 1", got)
    }
}
```

## Benchmarks

Functions of the form

```
func BenchmarkXxx(*testing.B)
```
are considered benchmarks, and are executed by the "go test" command when its -bench flag is provided. Benchmarks are run sequentially.

### Example:

```
func BenchmarkRandInt(b *testing.B) {
    for i := 0; i < b.N; i++ {
        rand.Int()
    }
}
```

The benchmark function must run the target code b.N times. During benchmark execution, b.N is adjusted until the benchmark function lasts long enough to be timed reliably. The output
