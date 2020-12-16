---

# Expectate: A lightweight testing utility for golang

This testing library's intended purpose is to provide 'expect-like' syntax to golang without using the heavy-handed ginkgo toolset.

---

# Principles:

## 1. Idiomatic
  It does nothing more than enhance the readability of the already awesome built-in Go testing toolkit!
  
## 2. Super Readable
  Taking heavy inspiration from JavaScript's jest library, this should make it perfectly clear what your assertions are.
  
## 3. Not For Everything!
  Think about whether or not you want your fail messages to be specific or not, because this library does not allow custom failure messages!
  Don't forget, this is just syntactic sugar, and it's only a matter of one or two extra lines to replicate the behavior.
  
---
  
# Examples:

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
--- FAIL: TestPeopleAreTheSame (0.00s)
    expect.go:46:   expectate_test.Person{
        -       Name:       "John Smith",
        +       Name:       "John Doe",
                Age:        31,
                Occupation: "Electrician",
          }
```
Note that when using go-cmp structs must have only exported fields! Unexported fields will result in a panic because it can't compare them!

But that's okay! You should really separate your pure data structures from your objects so that you never have to compare anything with unexported fields!
