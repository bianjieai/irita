package domain

type QueueMetaData struct {
	Timestamp uint64
	Height    uint64
}

type Context struct {
	height    uint64
	chainName string

	queue []QueueMetaData
}

func NewContext(height uint64, chainName string) *Context {
	return &Context{height: height, chainName: chainName, queue: []QueueMetaData{}}
}

// Height return the current height of the chain
func (ctx *Context) Height() uint64 {
	return ctx.height
}

func (ctx *Context) IncrHeight() {
	ctx.height += 1
}

func (ctx *Context) DecrHeight() {
	ctx.height -= 1
}

func (ctx *Context) SetHeight(height uint64) {
	ctx.height = height
}

// ChainName return the ChainName of the chain
func (ctx *Context) ChainName() string {
	return ctx.chainName
}

func (ctx *Context) Queue() []QueueMetaData {
	return ctx.queue
}

func (ctx *Context) SetQueue(queue []QueueMetaData) {
	ctx.queue = queue
}

func (ctx *Context) PushQueue(queue QueueMetaData) {
	ctx.queue = append(ctx.queue, queue)
}
