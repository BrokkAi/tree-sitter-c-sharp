# Brokk's C# Grammar for Tree-sitter

[![CI][ci]](https://github.com/BrokkAi/tree-sitter-c-sharp/actions/workflows/ci.yml)
[![crates][crates]](https://crates.io/crates/brokk-tree-sitter-c-sharp)
[![docs.rs][docs]](https://docs.rs/brokk-tree-sitter-c-sharp)

This is the **Brokk-owned and independently maintained fork** of
[`tree-sitter/tree-sitter-c-sharp`](https://github.com/tree-sitter/tree-sitter-c-sharp),
a C# grammar for [Tree-sitter](https://tree-sitter.github.io/tree-sitter/).
Brokk maintains this fork for use in its code-intelligence tooling and publishes
the Rust package as
[`brokk-tree-sitter-c-sharp`](https://crates.io/crates/brokk-tree-sitter-c-sharp).
It may intentionally diverge from upstream to support correctness and language
coverage required by Brokk.

## Installation

Add the Brokk-maintained Rust crate to your project:

```sh
cargo add brokk-tree-sitter-c-sharp@=0.23.6
```

Or add it directly to `Cargo.toml`:

```toml
[dependencies]
brokk-tree-sitter-c-sharp = "=0.23.6"
```

The grammar is based upon the Roslyn grammar with changes in order to:

- Deal with differences between the parsing technologies
- Work around some bugs in that grammar
- Handle `#if`, `#else`, `#elif`, `#endif` blocks
- Support syntax highlighting/parsing of fragments
- Simplify the output tree
- Reduce parser state count and complexity
- Be in-line with tree-sitter's convention where applicable

### Status

Comprehensively supports C# 1 through 14.0 with the following exceptions:

- [ ] `var` and `await` cannot be used as identifiers everywhere they are valid
- [ ] File-based apps preprocessor directives (`#:property`, `#:package`, `#:sdk`, `#:project`) are not yet recognized

### References

- [Official C# 8 Draft Language Spec](https://github.com/dotnet/csharpstandard/tree/draft-v8/standard) provides chapters that formally define the language grammar.
- [Roslyn C# language grammar export](https://github.com/dotnet/roslyn/blob/master/src/Compilers/CSharp/Portable/Generated/CSharp.Generated.g4)
- [SharpLab](https://sharplab.io) (web-based syntax tree playground based on Roslyn)

[ci]: https://img.shields.io/github/actions/workflow/status/BrokkAi/tree-sitter-c-sharp/ci.yml?logo=github&label=CI
[crates]: https://img.shields.io/crates/v/brokk-tree-sitter-c-sharp?logo=rust
[docs]: https://img.shields.io/docsrs/brokk-tree-sitter-c-sharp?logo=docs.rs
