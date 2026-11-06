# Go Study Plan — 12-Week Roadmap (Java → Go)

Since you are transitioning from Java, this 12-week accelerated roadmap cuts out general programming theory and focuses heavily on the structural, memory, and concurrency differences between the JVM and the Go runtime.

Here is your week-by-week execution plan.

---

## Month 1: The "Unlearn Java" Phase (Weeks 1–4)

**Goal:** Master structural typing, explicit error handling, and manual memory control. Eliminate Java-isms like deep inheritance, abstract factories, and exception throwing.

| Week 1                       | Week 2                    | Weeks 3 & 4                                |
| ---------------------------- | ------------------------- | ------------------------------------------ |
| Go types & pointers (vs JVM) | Composition & duck typing | Explicit errors & milestone 1 (Config CLI) |

### Week 1: Memory basics & pointers (the anti-JVM shift)

- **Deep dive:** In Java, primitives are values and objects are references. In Go, everything is passed by value unless you explicitly pass a pointer (`*`). Learn when to use pointer receivers vs. value receivers on structs.
- **Java translation:** Understand that a Go struct is a contiguous block of memory, unlike a Java object header plus heap pointers. This results in massive cache locality benefits.
- **Hands-on:** Practice initializing structs, passing them to functions by pointer vs. value, and observing mutations.

### Week 2: Struct composition & structural interface typing

- **Deep dive:** Go has no `extends` or `implements`. Interfaces are satisfied implicitly ("duck typing"). If a struct has a `Read()` method, it is an `io.Reader`. Master struct embedding to achieve composition over inheritance.
- **Java translation:** Interface definition belongs to the package consuming the behavior, not the package implementing it. Scrap your `service/impl/` folder habits entirely.
- **Hands-on:** Refactor a small mock Java-like class hierarchy into flat Go structs utilizing interface polymorphism.

### Week 3: Errors as values & standard I/O

- **Deep dive:** No try/catch. Errors are returned as normal values `(data, error)`. Master the `errors.Is()` and `errors.As()` paradigms for error wrapping. Work with standard I/O streams using `io.Reader`, `io.Writer`, and `bufio`.
- **Java translation:** Stop treating errors as exceptional flow-control mechanisms. In Go, a missing file or failed validation is an expected output value.
- **Hands-on:** Write a program that streams, parses, and modifies a heavy local file without loading the entire content into memory.

### Week 4: Phase milestone 1 — Configuration engine CLI

- **Expected outcome:** You can write structural, concise, framework-free Go code that cleanly processes files, maps data types, and safely handles edge cases.
- **Project specs:** Build a Configuration Engine CLI. It must natively parse configuration inputs from JSON files, YAML files, and environment flags using only standard packages (`flag`, `os`, `encoding/json`). It must cleanly merge these sources together based on priority and return wrapped, explicit semantic errors (e.g. `ErrMissingRequiredConfig`) without calling a single `panic`.

---

## Month 2: Production APIs & testing (Weeks 5–8)

**Goal:** Build highly performant web services using Go's standard library. Learn explicit wiring (no reflection-based dependency injection) and idiomatic testing.

| Week 5                      | Week 6              | Weeks 7 & 8                              |
| --------------------------- | ------------------- | ---------------------------------------- |
| Web services via `net/http` | Context & databases | Table tests, DI & milestone 2 (REST API) |

### Week 5: Frameworkless web servers via `net/http` & middleware

- **Deep dive:** Master the native `http.ServeMux` (routing, path variables, method matching). Learn to write standard HTTP middleware functions by wrapping `http.Handler`. Use Go's modern `slog` package for structured JSON logging.
- **Java translation:** Spring Boot handles routing and filters using heavy annotation reflection. Go achieves this through plain, nested function callbacks. It is faster, transparent, and completely type-safe.
- **Hands-on:** Write a small HTTP server with three routes and custom middleware for authentication headers and request logging.

### Week 6: Database/SQL & context lifecycle

- **Deep dive:** Master `database/sql` connection pools. Master the `context.Context` API for deadlines and cancellation propagation.
- **Java translation:** Java uses thread-interruption or transactional scopes. Go uses the `ctx` argument. Every external database call must receive a context (`QueryContext(ctx, ...)`). If a client cancels their HTTP request, the context automatically signals the SQL database to abort the query instantly, saving resource cycles.
- **Hands-on:** Wire a PostgreSQL or MySQL connection pool, executing transactional reads/writes while safely handling database errors and timed contexts.

### Week 7: Explicit dependency injection & table-driven tests

- **Deep dive:** Write unit tests using Go's `testing` package. Master table-driven tests using anonymous structs to sweep through hundreds of test matrix variables in a single test runner block.
- **Java translation:** No Spring `@Autowired`. You inject dependencies manually via constructor functions (e.g. `func NewService(repo Repository) *Service`). Swap Mockito for lightweight, manually written mock structs or native interface stubs.
- **Hands-on:** Implement unit tests for your database access layer using sub-tests (`t.Run`) and mock out external integrations manually.

### Week 8: Phase milestone 2 — production-ready REST API

- **Expected outcome:** A production-grade REST backend that gracefully manages network lifecycles, structured metrics, database state, and clean boundaries.
- **Project specs:** Build a Task & User Lifecycle Management API. It must utilize raw `net/http` (or a lightweight router like [chi](https://github.com/go-chi/chi)), include JWT validation middleware, connect to a real SQL database, implement graceful shutdown (intercepting `os.Interrupt`/`SIGTERM` to safely flush connection pools before killing the process), and feature greater than 80% code coverage via table-driven tests. Run `go test -race` to prove it runs free of structural data races.

---

## Month 3: Concurrency & performance tuning (Weeks 9–12)

**Goal:** Transition from Java's thread models to CSP concurrency. Master low-level performance metrics, allocation tracking, and gRPC communication frameworks.

| Week 9                    | Week 10                      | Weeks 11 & 12                                  |
| ------------------------- | ---------------------------- | ---------------------------------------------- |
| Goroutines & CSP channels | Sync locks vs. memory limits | Profiling, gRPC & milestone 3 (log aggregator) |

### Week 9: Goroutines & CSP channel architecture

- **Deep dive:** Master Go's M:N runtime scheduler. Understand how goroutines are multiplexed across OS threads. Master unbuffered vs. buffered channels, `select` statements, and worker pool patterns.
- **Java translation:** Even with Java 21 virtual threads, Go channels provide a completely unique paradigm: "Do not communicate by sharing memory; instead, share memory by communicating."
- **Hands-on:** Build an asynchronous worker pipeline that consumes tasks from an incoming channel, processes them across 5 concurrent goroutines, and outputs results into a collection channel.

### Week 10: Low-level concurrency & memory layout (stack vs. heap)

- **Deep dive:** Learn when channels are overkill and when to prefer `sync.Mutex`, `sync.RWMutex`, `sync.WaitGroup`, and `sync.Once`. Learn escape analysis by running `go build -gcflags="-m"`.
- **Java translation:** Java allocates objects aggressively on the heap, letting the garbage collector figure it out later. Go tries to store variables on the fast CPU stack; they only "escape" to the heap if their lifecycle leaves the local function execution frame. Minimizing heap allocations is the secret to senior-level Go performance.
- **Hands-on:** Write a performance-critical thread-safe cache using `sync.RWMutex`. Run escape analysis to verify that your data isn't leaking unnecessarily to the heap.

### Week 11: Application profiling via pprof & gRPC

- **Deep dive:** Instrument an application using the built-in runtime profiler (`net/http/pprof`). Generate and read CPU flame graphs, heap charts, and goroutine allocation profiles. Learn to define a gRPC service definition using Protocol Buffers (`.proto`) to achieve ultra-fast binaries.
- **Java translation:** Swap heavy tools like JProfiler for Go's native, low-overhead command-line `go tool pprof` utility.
- **Hands-on:** Run a benchmark (`go test -bench`), generate a CPU profile, find the hot path causing allocations, optimize the code using `sync.Pool`, and prove the performance increase with a second benchmark comparison.

### Week 12: Phase milestone 3 — high-throughput distributed log aggregator

- **Expected outcome:** Complete mastery over Go's concurrency, memory semantics, performance analysis tools, and cloud-native standards.
- **Project specs:** Build a Distributed Log Aggregator and Stream Service. The service must receive massive volumes of logs via a gRPC endpoint, funnel them into bounded internal channels, use concurrent worker pools to batch and transform the logs, and periodically flush them out to a mocked disk layer. You must use pprof under a heavy load test to isolate, analyze, and tune out any goroutine leaks or memory traps, providing a clean optimization report.
