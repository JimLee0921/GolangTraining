package geecache

import pb "geecache/geecachepb/geecachepb"

// PeerPicker 根据 key 选择应该从哪个远程服务器（peer）拉取数据
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter 跨节点获取数据
type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}
