package circularbuffer

import "testing"

func TestCircularBuffer_Insert(t *testing.T) {
	cb := NewCircularBuffer()

	// Insert values into the buffer
	cb.Insert(1)
	cb.Insert(2)
	cb.Insert(3)

	// Check if the values were inserted correctly
	expected := "[1, 2, 3]"
	result := cb.Print()
	if result != expected {
		t.Errorf("Insertion failed. Expected %v, got %v", expected, result)
	}
}

func TestCircularBuffer_Remove(t *testing.T) {
	cb := NewCircularBuffer()

	// Insert values into the buffer
	cb.Insert(1)
	cb.Insert(2)
	cb.Insert(3)

	// Remove a value from the buffer
	cb.Remove(2)

	// Check if the value was removed correctly
	expected := "[1, 3]"
	result := cb.Print()
	if result != expected {
		t.Errorf("Removal failed. Expected %v, got %v", expected, result)
	}
}

func TestCircularBuffer_Print(t *testing.T) {
	cb := NewCircularBuffer()

	// Insert values into the buffer
	cb.Insert(1)
	cb.Insert(2)
	cb.Insert(3)

	// Check if the print function returns the correct string representation
	expected := "[1, 2, 3]"
	result := cb.Print()
	if result != expected {
		t.Errorf("Print failed. Expected %v, got %v", expected, result)
	}
}
