type MyCalendar struct {
	bst *BST
}

func Constructor() MyCalendar {
	return MyCalendar{&BST{}}
}

func (this *MyCalendar) Book(start int, end int) bool {
	if floor, ok := this.bst.floor(start); ok {
		floorEnd := this.bst.get(floor)
		if floorEnd != -1 {
			if floorEnd > start {
				return false
			}
		}
	}

	if ceil, ok := this.bst.ceil(start); ok {
		if end > ceil {
			return false
		}
	}

	this.bst.insert(start, end)

	return true
}

type Node struct {
	key, val int
	left     *Node
	right    *Node
}

type BST struct {
	root *Node
}

func (b *BST) insert(key, val int) {
	b.root = insertHelper(b.root, key, val)
}

func insertHelper(x *Node, key, val int) *Node {
	if x == nil {
		return &Node{key, val, nil, nil}
	}
	if x.key == key {
		x.val = val
	} else if x.key < key {
		x.right = insertHelper(x.right, key, val)
	} else {
		x.left = insertHelper(x.left, key, val)
	}
	return x
}

func (b *BST) get(key int) int {
	x := getHelper(b.root, key)
	if x != nil {
		return x.val
	}
	return -1
}

func getHelper(x *Node, key int) *Node {
	if x == nil {
		return nil
	}
	if x.key == key {
		return x
	} else if key < x.key {
		return getHelper(x.left, key)
	} else {
		return getHelper(x.right, key)
	}
}

func (b *BST) floor(key int) (int, bool) {
	x := floorHelper(b.root, key)
	if x == nil {
		return 0, false
	}
	return x.key, true
}

func floorHelper(x *Node, key int) *Node {
	if x == nil {
		return nil
	}
	if x.key == key {
		return x
	} else if x.key > key {
		return floorHelper(x.left, key)
	} else {
		t := floorHelper(x.right, key)
		if t != nil {
			return t
		}
	}
	return x
}

func (b *BST) ceil(key int) (int, bool) {
	x := ceilHelper(b.root, key)
	if x == nil {
		return 0, false
	}
	return x.key, true
}

func ceilHelper(x *Node, key int) *Node {
	if x == nil {
		return nil
	}
	if x.key == key {
		return x
	} else if x.key < key {
		return ceilHelper(x.right, key)
	} else {
		t := ceilHelper(x.left, key)
		if t != nil {
			return t
		}
	}
	return x
}