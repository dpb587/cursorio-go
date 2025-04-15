# cursorio-go

Utilities for referencing byte and line+column offsets in UTF-8 streams.

## Usage

Import the module and refer to the code's documentation ([pkg.go.dev](https://pkg.go.dev/github.com/dpb587/cursorio-go/cursorio)).

```go
import "github.com/dpb587/cursorio-go/cursorio"
```

Some sample use cases and starter snippets can be found in the [`examples` directory](examples).

<details><summary><code>examples$ go run ./<strong>line-dump</strong> <<<'<strong>A-𝄞-Clef</strong>'</code></summary>

```
A
^ byte-offset 0; byte-count 1; text-range L1C1:L1C2

A-
 ^ byte-offset 1; byte-count 1; text-range L1C2:L1C3

A-𝄞
  ^ byte-offset 2; byte-count 4; text-range L1C3:L1C4

A-𝄞-
   ^ byte-offset 6; byte-count 1; text-range L1C4:L1C5
==== range L1C1:L1C5;0x0:0x7

A-𝄞-C
    ^ byte-offset 7; byte-count 1; text-range L1C5:L1C6
 ==== range L1C2:L1C6;0x1:0x8

A-𝄞-Cl
     ^ byte-offset 8; byte-count 1; text-range L1C6:L1C7
  ==== range L1C3:L1C7;0x2:0x9

A-𝄞-Cle
      ^ byte-offset 9; byte-count 1; text-range L1C7:L1C8
   ==== range L1C4:L1C8;0x6:0xa

A-𝄞-Clef
       ^ byte-offset 10; byte-count 1; text-range L1C8:L1C9
    ==== range L1C5:L1C9;0x7:0xb

```

</details>

More complex usage can be seen from importers like [inspecthtml-go](https://github.com/dpb587/inspecthtml-go), [inspectjson-go](https://github.com/dpb587/inspectjson-go), and [rdfkit-go](https://github.com/dpb587/rdfkit-go).

## Primitives

The `TextLineColumn` is a pair of `int64` values representing a line and its column. Within code it is 0-based, but its string form is 1-based and intended for humans. That is, the very first symbol of a stream starts from `TextLineColumn{0, 0}` which is printed as `L1C1`.

The `Offset` interface represents a position within a stream and is implemented by:

* `ByteOffset` as an `int64` value and formatted as `0x%x`.
* `TextOffset` as a `ByteOffset` + `TextLineColumn` tuple and formatted as `L%dC%d;0x%x`.

The `OffsetRange` interface represents a selection within a stream marked by two offsets. The `ByteOffsetRange` and `TextOffsetRange` implementations both contain two fields, `From` (inclusive) and `Until` (exclusive), for their respective offsets.

## Text Writer

The `TextWriter` supports tracking the lines and columns of a Unicode document. It acts as a standard `io.Writer` with getter functions for the current offsets, but offers several additional functions which may be more useful to lower-level tokenizer/scanner-type implementations.

* `WriteRunesForOffset` will write a slice of runes and return a `TextOffset`.
* `WriteRunesForOffsetRange` will write a slice of runes and return their `TextOffsetRange`.

> [!NOTE]
> As a reminder, Unicode makes line and column tracking non-trivial with its multi-byte code points and grapheme clusters. Put another way, N-bytes != N-runes != N-"columns" of printed symbols. This tries to abstract those complexities.

In code, use `NewTextWriter` to create an instance with an initial offset and begin writing.

```go
w := cursorio.NewTextWriter(cursorio.TextOffset{})
_ = w.WriteRunesForOffsetRange([]rune{0x1f477, 0x1f3fc})
// cursorio.TextOffsetRange{
//   From:cursorio.TextOffset{Byte:0, LineColumn:cursorio.TextLineColumn{0, 0}},
//   Until:cursorio.TextOffset{Byte:8, LineColumn:cursorio.TextLineColumn{0, 1}},
// }
```

## License

[MIT License](LICENSE)
