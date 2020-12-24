---

# Expectate: A lightweight testing utility for golang

This testing library's intended purpose is to provide 'expect-like' syntax to golang without using the heavy-handed ginkgo toolset.

---

# Principles

## 1. Idiomatic
  Well, sort of. More idiomatic than Ginkgo for sure! But I guess you could argue that you're *supposed* to write a bunch of if statements in your tests, as that is the Go way. I just think the expect syntax is cool, and if you don't like it you can just *go* away!
  
## 2. Super Readable
  Taking heavy inspiration from JavaScript's jest library, this should make it perfectly clear what your assertions are.
  
## 3. Not For Everything!
  Think about whether or not you want your fail messages to be specific or not, because this library does not allow custom failure messages!
  Don't forget, this is just syntactic sugar, and it's only a matter of one or two extra lines to replicate the behavior.
  
---
  
# Examples

The best thing any assertion library is for is testing the equality of two structs, right?<br>
Well, you'll be happy to know that, unlike testify, this uses go-cmp instead of reflect.DeepEquals!<br>
...Not impressed?<br>
Well, take go ahead and google go-cmp to figure out why. Also note that the use of reflect.DeepEquals is an outstanding issue in testify...

## Equality of two structs

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
Note that when using go-cmp structs must have only exported fields! Unexported fields will result in a panic because it can't compare them!

But that's okay! You should really separate your pure data structures from your objects so that you never have to compare anything with unexported fields!

# Documentation

Expectate only has 4 methods!

- ### [`Expect()`](#expect) <br>
- ### [`expect().ToBe()`](#tobe)
- ### [`expect().ToEqual()`](#toequal)
- ### [`expect().NotToBe()`](#nottobe)
- ### [`expect().NotToEqual()`](#nottoequal)

---

<h2 id="expect"><code>Expect()</code></h2>

The `Expect()` method is at the top level of the `expectate` package. It takes `*testing.T` as a parameter and returns an `ExpectorGenerator` type.

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
