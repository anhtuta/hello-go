# Map Keys in Go: No `hashCode` Needed

In Go, a type is automatically usable as a map key if it is **comparable**, meaning Go can compare two values of that type using `==`. For structs, Go compares them **field by field** — it does this automatically without you having to write anything.

So for your `Address` struct:

```go
a1 == a2  // true if all fields are equal: Name, city, Pincode
```

Go uses this built-in equality to handle map key lookups internally.

**Why no `hashCode` equivalent?** Go's runtime handles hashing internally. You don't define or expose a hash function — Go computes the hash of a struct key automatically based on its fields.

**Contrast with Java:**

|              | Java                            | Go                               |
| ------------ | ------------------------------- | -------------------------------- |
| Key equality | You must override `equals()`    | Automatic field-by-field `==`    |
| Hashing      | You must override `hashCode()`  | Runtime handles it automatically |
| Requirement  | Both methods must be consistent | Type must just be **comparable** |

**What types are NOT allowed as map keys in Go?** Types that are not comparable:

- `slice`
- `map`
- `function`

If you try to use a struct that contains a slice as a map key, Go will give a compile-time error.

So Go trades flexibility (you can't customize equality logic) for simplicity (nothing to implement).
