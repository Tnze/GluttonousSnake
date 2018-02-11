package ai

import "testing"

func TestPushAndpopMin(t *testing.T) {
	queue := NewQueue()
	queue.push(NewCell([2]int{7, 0}, 5, 0, nil))
	queue.push(NewCell([2]int{2, 0}, 1, 0, nil))
	queue.push(NewCell([2]int{3, 0}, 7, 0, nil))
	queue.push(NewCell([2]int{4, 0}, 8, 0, nil))
	c, _ := queue.popMinF()
	if c.Pos != [2]int{2, 0} {
		t.Error("error", c)
	}
	c, _ = queue.popMinF()
	if c.Pos != [2]int{7, 0} {
		t.Error("error", c)
	}
	queue.push(NewCell([2]int{3, 0}, 1, 0, nil))
	c, _ = queue.popMinF()
	if c.Pos != [2]int{3, 0} {
		t.Error("error", c)
	}
	c, _ = queue.popMinF()
	if c.Pos != [2]int{3, 0} {
		t.Error("error", c)
	}
}
func TestPushAndpop(t *testing.T) {
	queue := NewQueue()
	queue.push(NewCell([2]int{7, 0}, 5, 0, nil))
	queue.push(NewCell([2]int{2, 0}, 1, 0, nil))
	queue.push(NewCell([2]int{3, 0}, 7, 0, nil))
	queue.push(NewCell([2]int{4, 0}, 8, 0, nil))
	c, _ := queue.pop()
	if c.Pos != [2]int{7, 0} {
		t.Error("error", c)
	}
	c, _ = queue.pop()
	if c.Pos != [2]int{2, 0} {
		t.Error("error", c)
	}
	queue.push(NewCell([2]int{3, 0}, 1, 0, nil))
	c, _ = queue.pop()
	if c.Pos != [2]int{3, 0} {
		t.Error("error", c)
	}
	c, _ = queue.pop()
	if c.Pos != [2]int{4, 0} {
		t.Error("error", c)
	}
}

func TestGet(t *testing.T) {
	queue := NewQueue()
	queue.push(NewCell([2]int{7, 0}, 5, 0, nil))
	queue.push(NewCell([2]int{2, 0}, 1, 0, nil))
	queue.push(NewCell([2]int{3, 0}, 7, 0, nil))
	queue.push(NewCell([2]int{4, 0}, 8, 0, nil))
	c := queue.get([2]int{2, 0})
	if c.Pos != [2]int{2, 0} {
		t.Error("error", c)
	}
}

func TestExist(t *testing.T) {
	queue := NewQueue()
	queue.push(NewCell([2]int{7, 0}, 5, 0, nil))
	queue.push(NewCell([2]int{2, 0}, 1, 0, nil))
	queue.push(NewCell([2]int{3, 0}, 7, 0, nil))
	queue.push(NewCell([2]int{4, 0}, 8, 0, nil))
	c := queue.exist([2]int{2, 0})
	if !c {
		t.Error("error", c)
	}
}
