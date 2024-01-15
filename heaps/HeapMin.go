package heaps

type MinHeap struct {
	array []int
}

func minheap_constructor(value int) MinHeap {
	minHeap := MinHeap{}
	minHeap.insert_min_heap(value)
	return minHeap
}

func (m *MinHeap) insert_min_heap(value int) {
	m.array = append(m.array, value)
	m.minHeapifyUp(len(m.array) - 1)
}

func (m *MinHeap) minHeapifyUp(index int) {
	if index == 0 {
		return
	}

	for m.array[parent(index)] > m.array[index] {
		m.Swap(parent(index), index)
		index = parent(index)
	}
}

func (m *MinHeap) delete_min_heap() int {
	extracted := m.array[0]

	if len(m.array) == 0 {
		return -1
	}

	m.array[0] = m.array[len(m.array)-1]
	m.array = m.array[:]
	m.array = m.array[:len(m.array)-1]

	m.minHeapifyDown(0)

	return extracted
}

func (m *MinHeap) minHeapifyDown(index int) {
	lastIndex := len(m.array) - 1
	l, r := leftChild(index), rightChild(index)

	if index >= lastIndex || l >= lastIndex {
		return
	}

	lV := m.array[l]
	rV := m.array[r]
	v := m.array[index]

	if lV > rV && v > rV {
		m.Swap(r, index)
		m.minHeapifyDown(r)
	} else if rV > lV && v > lV {
		m.Swap(l, index)
		m.minHeapifyDown(l)
	}
}

func parent(idx int) int {
	return (idx - 1) / 2
}

func leftChild(idx int) int {
	return (idx * 2) + 1
}

func rightChild(idx int) int {
	return (idx * 2) + 2
}

func (m *MinHeap) Swap(i1, i2 int) {
	m.array[i1], m.array[i2] = m.array[i2], m.array[i1]
}
