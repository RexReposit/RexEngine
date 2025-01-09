package src

type NodeInterface interface {
	Init()
	Process()
	getName() string
	setParent(node *Node)
	setInit(state bool)
}

type Node struct {
	Name     string
	Children map[string]NodeInterface
	Parent   *Node
	isInit   bool
}

func NewRootNode() *Node {
	return &Node{
		Name:     "root",
		Children: make(map[string]NodeInterface),
		Parent:   nil,
		isInit:   false,
	}
}

func (n *Node) Init() {
	n.setInit(true)
	for _, child := range n.Children {
		child.setParent(n)
		child.Init()
		child.setInit(true)
	}
}

func (n *Node) Process() {
	for _, child := range n.Children {
		child.Process()
	}
}

func (n *Node) setInit(state bool) {
	if n.isInit {
		return
	}
	n.isInit = state
}

func (n *Node) setParent(node *Node) {
	n.Parent = node
}

func (n *Node) getName() string {
	return n.Name
}

func (n *Node) AddChild(child NodeInterface) {
	n.Children[child.getName()] = child
}

func (n *Node) RemoveChild(child NodeInterface) {
	delete(n.Children, child.getName())
}
