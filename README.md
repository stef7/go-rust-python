This is a challenge proposed by ChatGPT for learning Go, Rust and Python. I asked it for a progressive challenge that would help me, as an experienced Node developer, bump up against the key differences and idiosyncrasies between the languages and the mental models behind them, especially compared to Node.

It suggested a JSON CLI tool into 3 laddered mini-challenges, each progressively forcing more "language culture shock" and deeper learning, keeping things manageable while hitting all the key quirks.

Recursion is explicitly covered in Challenge 3 for Rust, since it's not organically needed otherwise.

# Progressive Challenge Series: JSON CLI Tool

## Challenge 1: Read and Parse a JSON File
**Goal:** Read a local `users.json` file and print each _(top-level)_ user's name and age.

### You’ll Learn:
* How to handle CLI args
* Basic file I/O
* Basic JSON parsing
* Structs / type annotations
* Printing and string formatting
* Error handling boilerplate

### Differences you’ll notice:
* **Go:** Need to decode JSON into a `struct`; error returns.
* **Rust:** Borrow checker, ownership of file/contents, need to derive `Deserialize`.
* **Python:** Dynamic parsing is easier, but formatting errors still trip up.

## Challenge 2: Filter and Sort
**Goal:** Filter _(top-level)_ for `active == true` and `age > 30`, then sort by `lastLogin` and print names.

### You’ll Learn:
* Collection filtering and sorting idioms
* Closures / lambdas
* Date parsing
* More robust error handling and conditionals
* Comparison operators and field access quirks

### Gotchas:
* **Go:** No generics for sorting; must implement `sort.Interface`.
* **Rust:** Lifetimes and `.iter().filter().collect()` can be surprising.
* **Python:** List comprehensions feel like home—too easy? Maybe.
* **Date parsing:** All three are clunky. You’ll be Googling formats.

## Challenge 3: Use Recursion to Find the Most Active Tree of Users
**Goal:** Recursively walk the _full_ structure to find the most recently active user over 30.

### You’ll Learn:
* Recursion
* Optional / nullable values (children may be missing)
* Tree traversal
* Maximum value tracking

### Rust-Specific Gotcha:
* You’ll **struggle** with recursive structs needing `Box` to break infinite size.
* Rust’s stack isn’t deep—so large trees need tail recursion or iteration.
* Ownership/borrowing of subtrees will make you cry (a little).
