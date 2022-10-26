package main
import "fmt"

type Node struct {
 value int
 next *Node
}


//This is a healper struct containing info on stack 
type StackLL struct {
 head *Node //keeps track of top of the stack
 size int
}

// Return the Sizeof the stack at any time
func (s *StackLL) Size() int {
    return s.size
}

//Check whether Stack is empty or not 
func (s *StackLL) IsEmpty() bool {
    return s.size == 0
}

// Peek() function 
func (s *StackLL) Peek() (int) {
    if s.IsEmpty() {
        fmt.Println("StackEmptyException")
        return 0
    }
    return s.head.value
}

// Push() function: Add the element from front of LL
func (s *StackLL) Push(value int) {
    s.head = &Node{value, s.head} //initialize the current head to new node next
    /* Alternatively you can do this way*/
    //temp := &Node{value,s.head}
    //s.head = temp
    s.size++
}

// Pop() function
func (s *StackLL) Pop() (int, bool) {
    if s.IsEmpty() {
        fmt.Println("StackEmptyException")
        return 0, false
    }
    value := s.head.value
    s.head = s.head.next
    s.size--
    return value, true
}

// Print() function 
func (s *StackLL) Print() {
    temp := s.head
    fmt.Print("Values stored in stack are: ")
    //we use for keyword instead of while as we don't have while keyword
    for temp != nil { 
        fmt.Print(temp.value, " ")
        temp = temp.next
    } 
    fmt.Println()
}

func main() {
	var stack StackLL // create a stack variable of type StackLinkedList
    /* Adding items to stack */
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
  
    stack.Print() // Print the stack
    fmt.Println("Checking length of stack:")
    fmt.Println(stack.Size()) // Get Length
    fmt.Println("Removing an Item:")
    stack.Pop() // Remove an item
    stack.Print() 
    fmt.Println("Getting Item at Top of Linked List")
    fmt.Println(stack.Peek()) // Get Top-most item
}

