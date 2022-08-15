package q

//import "fmt"

// func main() {
// 	arr := []int{5, 7, 1, 9, 2, 33, 21, 30}
// 	pq := &MaxPQ{}
// 	for _, val := range arr {
// 		pq.Insert(val)
// 	}
// 	fmt.Println(pq.DelMax())

// }

//优先级队列 父亲大于左右孩子，小则不配做父需要下沉 孩子比父节点大则不应该做孩子需要上浮.使用切片，0位置舍弃 所以位置>=1
type MaxPQ []int

//x位置元素的父元素位置
func (m *MaxPQ) Parent(x int) int {
	return x / 2
}

//x位置元素的左孩子位置
func (m *MaxPQ) Left(x int) int {
	return x * 2
}

//x位置元素的右孩子位置
func (m *MaxPQ) Right(x int) int {
	return x*2 + 1
}

//交互两个位置元素值
func (m *MaxPQ) Swap(i int, j int) {
	q := *m
	q[i], q[j] = q[j], q[i]
}

//比较两个位置元素大小
func (m *MaxPQ) Less(i int, j int) bool {
	q := *m
	return q[i] < q[j]
}

//上浮x位置元素 比较自己的父节点大则上浮
func (m *MaxPQ) Swim(x int) {
	for x > 1 && m.Less(m.Parent(x), x) { //x最多到堆顶
		m.Swap(m.Parent(x), x) //交换父子位置上的数值
		x = m.Parent(x)        //当前数的位置变为父亲的位置
	}
}

//下沉x位置元素 只要比左右孩子其中1个小就要下沉
func (m *MaxPQ) Down(x int) {
	q := *m
	for m.Left(x) < len(q) { //x左子树如果超过边界就不可能再下沉了
		max := m.Left(x)                                    //假设左子树最大
		if m.Right(x) < len(q) && m.Less(max, m.Right(x)) { //比较左右子树,得到最大子树
			max = m.Right(x)
		}
		//如果x大于左右孩子则不需要下沉了
		if m.Less(max, x) {
			break
		}
		//否则下沉，交换
		m.Swap(max, x)
		x = max
	}
}

//删除顶部最大元素  交换顶部和最底部元素，并删除底部元素，顶部元素下沉到合适位置
func (m *MaxPQ) DelMax() int {
	q := *m
	max := q[1]
	n := len(q)
	m.Swap(1, n-1)
	*m = q[:n-1]
	m.Down(1)
	return max
}

//插入元素 从尾部插入，进行上浮
func (m *MaxPQ) Insert(v int) {
	*m = append(*m, v)
	m.Swim(len(*m) - 1)
}
