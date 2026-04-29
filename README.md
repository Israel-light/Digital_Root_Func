# Digital Root Function

A Go implementation of the **digital root** algorithm that repeatedly sums the digits of a number until a single digit is obtained.

## What is Digital Root?

The **digital root** (also called the repeated digit sum) of a non-negative integer is the single digit obtained by an iterative process of summing digits:

1. Sum all digits of the number
2. If the result has more than one digit, repeat step 1
3. Continue until you get a single digit

### Examples

- **467** → 4 + 6 + 7 = **17** → 1 + 7 = **8**
- **21** → 2 + 1 = **3**
- **1000** → 1 + 0 + 0 + 0 = **1**
- **1** → **1** (already single digit)
- **-36** → 36 → 3 + 6 = **9** (handles negatives by taking absolute value)

## Features

✨ **Simple & Efficient** - Clean recursive implementation
✨ **Handles Negatives** - Converts to absolute value
✨ **Well-Tested** - Includes multiple test cases in `main()`

## Usage

### Basic Example

```go
package main

import (
	"fmt"
)

func main() {
	result := singleDigit(467)
	fmt.Println(result)  // Output: 8
}
```

### Available Function

```go
func singleDigit(n int) int
```

**Parameters:**
- `n` (int) - The number to calculate the digital root for

**Returns:**
- (int) - The single digit digital root

## How It Works

The `singleDigit()` function:

1. Converts negative numbers to their absolute value
2. Returns immediately if the number is already single-digit (< 10)
3. Sums all digits by converting to string and iterating through characters
4. Recursively calls itself on the sum if it's greater than 9
5. Returns the final single digit

## Running the Code

```bash
go run main.go
```

**Output:**
```
8
3
1
1
9
```

## Algorithm Complexity

- **Time Complexity:** O(log n) where n is the input number
- **Space Complexity:** O(log n) due to recursion depth

## Mathematical Note

The digital root has a mathematical property: For any positive integer n, the digital root equals:
- `(n - 1) % 9 + 1` (when n ≠ 0)

This could be used for a more efficient O(1) implementation, but the current approach is more intuitive.

## License

This project is open source and available under the MIT License.

## Author

Created by [Israel-light](https://github.com/Israel-light)
