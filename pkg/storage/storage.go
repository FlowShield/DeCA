package storage

// ExecCloser 存取数据，并关闭
type ExecCloser interface {
	Put(data string) (string, error)
	Get(cid string) (string, error)
	Close() error
}
