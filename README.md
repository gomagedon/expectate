---
# Expectate: A lightweight testing utility for golang
---

# Quick Start

### Install:

```
go get github.com/gomagedon/expectate
```

```golang
package mylibrary_test

import (
  "testing"
  "github.com/gomagedon/expectate"

  "github.com/foo/mylibrary"
)

func TestProject(t *testing.T) {
  expect := expectate.Expect(t)

  expect(mylibrary.HelloWorld()).ToBe("Hello world!\n")
}
```

## Test Equality of Two Structs

```golang

type Person struct {
  Name string
  Age int
  Occupation string
}

func TestPeopleAreTheSame(t *testing.T) {
  expect := expectate.Expect(t)
  person1 := Person{
    Name: "John Doe",
    Age: 31,
    Occupation: "Electrician",
  }
  person2 := Person{
    Name: "John Smith",
    Age: 31,
    Occupation: "Electrician",
  }

  expect(person1).ToEqual(person2)
}

```

Output:

```
--- FAIL: TestPeopleAreTheSame
    expect.go:46:   expectate_test.Person{
        -       Name:       "John Smith",
        +       Name:       "John Doe",
                Age:        31,
                Occupation: "Electrician",
          }
```

Note that expectate uses `google/go-cmp` for testing strict equality. Custom cmp options are not supported for the `ToEqual()` method, but you can always call the cmp library directly.

# Simple API

Expectate uses only 4 methods!

- ### [`Expect()`](#expect) <br>
- ### [`expect().ToBe()`](#tobe)
- ### [`expect().ToEqual()`](#toequal)
- ### [`expect().NotToBe()`](#nottobe)
- ### [`expect().NotToEqual()`](#nottoequal)

---

<h2 id="expect"><code>Expect()</code></h2>

The `Expect()` method is at the top level of the `expectate` package. It takes `*testing.T` as a parameter and returns an `ExpectorFunc` type.

Here's an example:

```golang
func TestSomething(t *testing.T) {
  expect := expectate.Expect(t)

  expect(something).ToBe(somethingElse)
  // ...some other expectations
}
```

Alternatively:

```golang
func TestSomething(t *testing.T) {
  expectate.Expect(t)(something).ToBe(somethingElse)
  // ...some other expectations
}
```

---

<h2 id="tobe"><code>expect().ToBe()</code></h2>

The `expect().ToBe()` method exists on the `Expector` type (which is what `expect()` returns). It takes any value and performs a simple equality check with that value and the initial value passed to `expect()`. If the two values are equal, the test passes. If not, it calls `t.Fatal()` with a generic error message.

#### Here's an example of a passing test:

```golang
func TestFooIsFoo(t *testing.T) {
  expect := expectate.Expect(t)

  expect("foo").ToBe("foo")
}
```

Result:

```
--- PASS (ok)
```

#### Here's an example of a failing test:

```golang
func TestFooIsBar(t *testing.T) {
  expect := expectate.Expect(t)

  expect("foo").ToBe("bar")
}
```

Result:

```
--- FAIL: TestFooIsBar
    expect.go:34: foo is not bar
```

---

<h2 id="toequal"><code>expect().ToEqual()</code></h2>

The `expect().ToEqual()` method exists on the `Expector` type (which is what `expect()` returns). It takes any value and performs a _deep_ equality check using [go-cmp](https://github.com/google/go-cmp) with that value and the initial value passed to `expect()`. If the two values are equal, the test passes. If not, it calls `t.Fatal()` with a go-cmp diff.

#### Here's an example of a passing test:

```golang
func TestPost1EqualsPost2(t *testing.T) {
  expect := expectate.Expect(t)

  post1 := Post{
    Title:   "Post",
    Content: "Content of Post",
    Likes:   2,
  }
  post2 := Post{
    Title:   "Post",
    Content: "Content of Post",
    Likes:   2,
  }

  expect(post1).ToEqual(post2)
}
```

Result:

```
--- PASS (ok)
```

#### Here's an example of a failing test:

```golang
func TestPost1EqualsPost2(t *testing.T) {
  expect := expectate.Expect(t)

  post1 := Post{
    Title:   "Post 1",
    Content: "Content of Post 1",
    Likes:   2,
  }
  post2 := Post{
    Title:   "Post 2",
    Content: "Content of Post 2",
    Likes:   1,
  }

  expect(post1).ToEqual(post2)
}
```

Output:

```
--- FAIL: TestPost1EqualsPost2
    expect.go:43:   main.Post{
        -       Title:   "Post 2",
        +       Title:   "Post 1",
        -       Content: "Content of Post 2",
        +       Content: "Content of Post 1",
        -       Likes:   1,
        +       Likes:   2,
          }
```

---

<h2 id="nottobe"><code>expect().NotToBe()</code></h2>

The `expect().NotToBe()` method exists on the `Expector` type (which is what `expect()` returns). It has the opposite behavior of `expect().ToBe()` It takes any value and performs a simple inequality check with that value and the initial value passed to `expect()`. If the two values are not equal, the test passes. If they are, it calls `t.Fatal()` with a generic error message.

#### Here's an example of a passing test:

```golang
func TestFooIsNotBar(t *testing.T) {
  expect := expectate.Expect(t)

  expect("foo").NotToBe("bar")
}
```

Result:

```
--- PASS (ok)
```

#### Here's an example of a failing test:

```golang
func TestFooIsNotFoo(t *testing.T) {
  expect := expectate.Expect(t)

  expect("foo").NotToBe("foo")
}
```

Result:

```
--- FAIL: TestFooIsBar
    expect.go:34: foo is foo
```

---

<h2 id="nottoequal"><code>expect().NotToEqual()</code></h2>

The `expect().NotToEqual()` method exists on the `Expector` type (which is what `expect()` returns). It has the opposite behavior as `expect().ToEqual()`. It takes any value and performs a _deep_ inequality check using [go-cmp](https://github.com/google/go-cmp) with that value and the initial value passed to `expect()`. If the two values are not equal, the test passes. If they are, it calls `t.Fatal()` with a generic error message.

#### Here's an example of a passing test:

```golang
func TestPost1_DoesNotEqual_Post2(t *testing.T) {
  expect := expectate.Expect(t)

  post1 := Post{
    Title:   "Post 1",
    Content: "Content of Post 1",
    Likes:   2,
  }
  post2 := Post{
    Title:   "Post 2",
    Content: "Content of Post 2",
    Likes:   1,
  }

  expect(post1).NotToEqual(post2)
}
```

Result:

```
--- PASS (ok)
```

#### Here's an example of a failing test:

```golang
func TestPost1_DoesNotEqual_Post2(t *testing.T) {
  expect := expectate.Expect(t)

  post1 := Post{
    Title:   "Post",
    Content: "Content of Post",
    Likes:   2,
  }
  post2 := Post{
    Title:   "Post",
    Content: "Content of Post",
    Likes:   2,
  }

  expect(post1).NotToEqual(post2)
}
```

Output:

```
--- FAIL: TestPost1_DoesNotEqual_Post2
    expect.go:57: {Post 1 Content of Post 1 2} equals {Post 1 Content of Post 1 2}
```
