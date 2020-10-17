package namenode

// NameNode is struct maintaining name-meta
type NameNode struct {
	fileName string
}

// File is interface for file path
type File struct {
	name string
}

func (nameNode *NameNode) setName(fileName File) {
	nameNode.fileName = fileName.name
}

func (nameNode *NameNode) getName() string {
	return nameNode.fileName
}

// NewNameNode new instance of name node
func NewNameNode() *NameNode {
	return new(NameNode)
}
