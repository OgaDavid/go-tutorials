# Variables and Data Types

`Basic Types`

| Category         | Type                                     | Example              | Notes                               |
| ---------------- | ---------------------------------------- | -------------------- | ----------------------------------- |
| Integer          | `int`, `int8`, `int16`, `int32`, `int64` | `age := 25`          | `int` adjusts to system (32/64-bit) |
| Unsigned Integer | `uint`, `uint8`, etc.                    | `count := uint(5)`   | No negatives allowed                |
| Float            | `float32`, `float64`                     | `pi := 3.14159`      | `float64` is preferred              |
| Complex          | `complex64`, `complex128`                | `z := complex(3, 4)` | Rarely used                         |
| Boolean          | `bool`                                   | `isActive := true`   | only `true` or `false`              |
| String           | `string`                                 | `name := "David"`    | Immutable                           |

`Derived Types`

| Type      | Example                                       | Description                                 |
| --------- | --------------------------------------------- | ------------------------------------------- |
| Array     | `[3]int{1, 2, 3}`                             | Fixed-size, same type                       |
| Slice     | `[]int{1, 2, 3}`                              | Dynamic array (most used)                   |
| Map       | `map[string]int{"A": 1, "B": 2}`              | Key-value pairs                             |
| Struct    | `type Person struct { name string; age int }` | Custom type with multiple fields            |
| Pointer   | `var p *int`                                  | Stores memory address of a value            |
| Function  | `func add(a int, b int) int { return a + b }` | Functions are first-class citizens          |
| Interface | `interface {}`                                | Defines method sets (advanced but powerful) |
