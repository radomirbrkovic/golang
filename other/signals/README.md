# Signals | [Source](https://gobyexample.com/signals)

Sometimes we'd like our Go programs to intelligently handle [Unix Signals](https://en.wikipedia.org/wiki/Signal_(IPC)). For example, we might want a server o gracefully shutdown when it receives a `SIGTERM`, or a command-line tool to stop processing input if it receives a `SIGINT`. Here's how to handle signals in Go with channels.  