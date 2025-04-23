package main
 
import "fmt"
 
// BIT represents a Fenwick Tree
type BIT struct {
    tree []int
}
 
// NewBIT creates a new Fenwick Tree with the given size
func NewBIT(size int) *BIT {
    return &BIT{tree: make([]int, size+1)}
}

// Update updates the value at index i to val in the Fenwick Tree
func (bt *BIT) Update(i, val int) {
    for i < len(bt.tree) {
        bt.tree[i] += val
        i += i & -i // i += i & (-i) in C++ style (i & -i gets the least significant 1-bit)
    }
}

// Query returns the sum of elements from index 0 to i (inclusive) in the Fenwick Tree
func (bt *BIT) Query(i int) int {
    sum := 0
    for i > 0 {
        sum += bt.tree[i]
        i -= i & -i // i -= i & (-i) in C++ style (i & -i gets the least significant 1-bit)
    }
    return sum
}

func main() {
    n := 10 // Size of the array (elements are indexed from 1 to n)
    bit := NewBIT(n) // Create a Fenwick Tree of size n+1 (to handle index from 1 to n)
    data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Original array values (indexed from 1 to n)
    for i := 1; i <= n; i++ { // Initialize the Fenwick Tree with original values (optional)
        bit.Update(i, data[i-1]) // Note: Indexing starts from 1 in the original array but from 0 in the BIT tree. So we adjust accordingly.
    }
    // Query sum from index 1 to 5 (inclusive)
    fmt.Println("Sum from index 1 to 5:", bit.Query(5)) // Output: Sum from index 1 to 5: 15 (1+2+3+4+5)
    // Update value at index 3 to 21 and query again from index 1 to 5 (inclusive)
    bit.Update(3, 17) // Increase value at index 3 by 17 (original value was 4)
    fmt.Println("Updated sum from index 1 to 5:", bit.Query(5)) // Output: Updated sum from index 1 to 5: 32 (1+2+3+(4+17)+5)
}