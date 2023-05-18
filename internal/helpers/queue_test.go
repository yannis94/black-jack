package helpers

import "testing"

var testQueue *Queue

type person struct {
    name string
    age int
}

func init() {
    testQueue = NewQueue()
}

func TestPeek(t *testing.T) {
    data := "My unit test"
    q := &Queue{ Length: 1, head: &node{ Data: data } }
    peeked, err := q.Peek()

    if err != nil {
        t.Errorf("Cannot peek, error : %s", err.Error())
    } else if peeked != "My unit test" {
        t.Errorf("The peek found uncorresponding data: %s", peeked)
    }
}

func TestEnqueue(t *testing.T) {
    bob := &person{ name: "Bob", age: 20 } 
    john := &person{ name: "John", age: 22 }
    andrew := &person{ name: "Andrew", age: 19 }

    testQueue.Enqueue(*bob)

    if testQueue.Length != 1 {
        t.Errorf("The queue's length should be 1 but it is %d", testQueue.Length)
    }

    testQueue.Enqueue(*john)
    testQueue.Enqueue(*andrew)

    peeked, _:= testQueue.Peek()

    if lastInserted, ok := peeked.(person); !ok {
        t.Errorf("Data enqueueing faild.")
    } else if lastInserted.name != "Bob" && lastInserted.age != 20 {
        t.Errorf("Peek looks for head of the queue wich should be 'Bob, 20 yo' and it is '%s, %d yo'.", lastInserted.name, lastInserted.age)
    }

}

func TestDequeue(t *testing.T) {
    dequeued, err := testQueue.Dequeue()

    if err != nil {
        t.Errorf("Error returned: %s", err.Error())
    } 

    if p, ok := dequeued.(person); !ok {
        t.Errorf("Data extraction type not expected.")
    } else if p.name != "Bob" {
        t.Errorf("The person extract is '%s' and it should be 'Bob'", p.name)
    } else if testQueue.Length != 2 {
        t.Errorf("Queue's length should be 2 and it is %d", testQueue.Length)
    }

    testQueue.Dequeue()
    testQueue.Dequeue()

    _, err = testQueue.Dequeue()

    if err == nil {
        t.Errorf("The dequeue function should return an error because queue's length is 0")
    }

}
