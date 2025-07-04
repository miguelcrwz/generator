<div align="center">

<pre>
 ┌────────────────────────────────────────────────────┐
 │  user@dev:~/projects$ ./generator                  │
 │                                                    │
 │       ██████╗    █████╗   ███████╗  ███████╗       │
 │       ██╔══██╗  ██╔══██╗  ██╔════╝  ██╔════╝       │
 │       ██████╔╝  ███████║  ███████╗  ███████╗       │
 │       ██╔═══╝   ██╔══██║  ╚════██║  ╚════██║       │
 │       ██║       ██║  ██║  ███████║  ███████║       │
 │       ╚═╝       ╚═╝  ╚═╝  ╚══════╝  ╚══════╝       │
 │                                                    │
 │  > random password generator · go 1.23             │
 └────────────────────────────────────────────────────┘
</pre>

![type](https://img.shields.io/badge/type-study%20project-2ea043?style=for-the-badge&labelColor=0d1117)
![version](https://img.shields.io/badge/version-1.0.0-58a6ff?style=for-the-badge&labelColor=0d1117)
![license](https://img.shields.io/badge/license-mit-8b949e?style=for-the-badge&labelColor=0d1117)
![stack](https://img.shields.io/badge/stack-go%201.23-d29922?style=for-the-badge&labelColor=0d1117)

</div>

---

## `$ cat about.txt`

```
generator is a command-line tool that asks for a length and prints a
random password with exactly that many characters.

two slots are reserved up front: one uppercase letter and one digit,
drawn on their own. the other length-2 bytes come from an 81-char pool
(lowercase, uppercase, digits and 19 symbols), the two reserved ones
are appended, and rand.Shuffle mixes the whole slice so they do not
always land at the end.

built as a hands-on study of go — byte slices, math/rand and ansi
colors in the terminal.
```

## `$ ls features/`

```
[x] any length from 2 up, typed at the prompt
[x] always at least one uppercase letter and one digit
[x] 81-char pool: a-z, A-Z, 0-9 and !@#$%^&*()+?><:{}[]
[x] shuffled after the forced characters go in
[x] colored prompt, label and error with ansi escape codes
[ ] math/rand seeded with the clock, not crypto/rand — not for secrets
[ ] no lowercase letter or symbol is guaranteed, only upper + digit
[ ] length < 2 or no number read: error shown as the password, exit 0
[ ] trailing junk is ignored — "12abc" and "3.5" read as 12 and 3
[ ] no upper limit — a huge length panics in make() or runs out of memory
[ ] ansi codes always printed, no tty check — pipes and files get them too
[ ] rand.Seed is deprecated since go 1.20
[ ] "LENGHT" typo in the prompt and the error message
[ ] no flags (length only via the prompt), no tests
```

## `$ cat stack.lock`

```
runtime     go 1.23.3 (go.mod)
framework   none — stdlib only (fmt, math/rand, time)
random      math/rand, seeded with time.Now().UnixNano()
infra       none, runs locally
tests       none yet
tooling     go run / go build
```

## `$ ./install.sh`

```bash
# clone
git clone https://github.com/miguelcrwz/generator.git
cd generator

# dependencies (stdlib only — just needs go 1.23.3+)
go version
```

## `$ ./run`

```bash
# development — asks for the length on stdin
go run .

# production — keep -o: the module is github.com/miguelcrwz,
# so a plain go build names the binary miguelcrwz
go build -o generator . && ./generator

# piped — the prompt still prints, so the password lands on the same line
echo 16 | go run .
```

```
Enter the PASSWORD LENGHT: 16
Generated PASSWORD: <o{7AVm&Z%5%M>rF
```

```
Enter the PASSWORD LENGHT: 1
Generated PASSWORD: ✘ PASSWORD LENGHT MUST BE AT LEAST 2
```

```
Enter the PASSWORD LENGHT: Generated PASSWORD: ]QLr562:JPhKV9j@
```

## `$ cat generate.flow`

```
stdin ─> fmt.Scan(&length)
  │
  ├─ no digits or int overflow ──> length stays 0, falls into < 2
  ├─ length < 2 ─────────────────> red error returned as the password
  │
  └─ ok (no upper limit)
       ├─ seed math/rand with time.Now().UnixNano()
       ├─ draw 1 uppercase + 1 digit ───────> mandatory
       ├─ make([]byte, length-2) ───────────> a huge length crashes here
       ├─ fill length-2 bytes from the 81-char pool
       ├─ append mandatory, rand.Shuffle the slice
       └─ print "Generated PASSWORD:" ──────> stdout
```

```
999999999999999999 ─> panic: runtime error: makeslice: len out of range
100000000000000 ────> fatal error: runtime: cannot allocate memory
```

## `$ tree`

```
generator/
├── main.go            # generatePassword + main: prompt, read, print
├── preview/
│   └── gopher.gif     # the gopher that celebrates exit 0
├── go.mod             # module github.com/miguelcrwz, go 1.23.3
├── .gitignore         # binaries, test output, go.work, editor/os files
├── LICENSE
└── README.md
```

## `$ git log --oneline roadmap`

```
a1b2c3d  switch to crypto/rand and drop the deprecated rand.Seed
e4f5g6h  reject bad input with an error and exit 1, fix "LENGHT"
i7j8k9l  guarantee a lowercase letter and a symbol too
m0n1o2p  take the length from a -n flag, keep the prompt as fallback
q3r4s5t  table tests for lengths and character classes
```

## `$ cat LICENSE | head -1`

```
MIT © 2025 Miguel Cruz
```

## `$ whoami`

```
miguel cruz
├── github    https://github.com/miguelcrwz
├── linkedin  https://linkedin.com/in/miguelcrwz
└── email     miguwlcrwz@gmail.com
```

<div align="center">

<pre>
$ exit
[process exited with code 0]
</pre>

<img src="preview/gopher.gif" alt="dancing gopher" width="128">

</div>
