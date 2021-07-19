# Atomic Counters

The primary mechanism for managing state in Go is communication over chanel. Here we’ll look at using the `sync/atomic` package for atomic counters accessed by multiple goroutines.