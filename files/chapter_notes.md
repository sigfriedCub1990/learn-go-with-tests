# NOTES

### Mindset and goals:
- **Write the test we want to see**. Think about how we'd like to use the code we're going to write from a consumer's point of view.
- Focus on _what_ and _why_, but don't get distracted by the _how_.

[Leveling-up TDD](https://deniseyu.github.io/leveling-up-tdd/)


### Addendum

#### Data structure

```go 
type Post struct {
	Title, Description, Body string
	Tags                     []string
}
```

#### File example

```md
Title: Hello, TDD world!
Description: First post on our wonderful blog
Tags: tdd, go
---
Hello world!

The body of posts starts after the `---`
```
