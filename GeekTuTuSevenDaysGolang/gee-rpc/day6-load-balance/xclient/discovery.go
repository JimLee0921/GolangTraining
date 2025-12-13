package xclient

import (
	"errors"
	"math"
	"math/rand"
	"sync"
	"time"
)

// SelectMode 代表不同的负载均衡策略
type SelectMode int

// 进实现 Random 和 RoundRobin 策略
const (
	RandomSelect SelectMode = iota
	RoundRobinSelect
)

type Discovery interface {
	Refresh() error                      // 若有注册中心，实现定期刷新服务列表
	Update(servers []string) error       // 手动更新实例列表
	Get(mode SelectMode) (string, error) // 根据策略选择一个实例地址
	GetAll() ([]string, error)           // 获取所有实例地址列表
}

// MultiServersDiscovery 不依赖注册中心，仅使用一个 静态列表 + 本地维护 的实现
type MultiServersDiscovery struct {
	mu      sync.RWMutex
	servers []string   // 服务实例地址列表
	index   int        // 用于轮询算法记录上一次选到的位置
	r       *rand.Rand // 用于随机选择时的随机数生成器
}

func NewMultiServerDiscovery(servers []string) *MultiServersDiscovery {
	// 直接保存传入的服务器数组
	d := &MultiServersDiscovery{
		servers: servers,
		r:       rand.New(rand.NewSource(time.Now().UnixNano())), // 使用时间戳设定随机数种子
	}
	d.index = d.r.Intn(math.MaxInt32 - 1) // index 记录 round robin 算法已经轮询到的位置，为避免每次从0，初始化设置一个值
	return d
}

// 编译器检查
var _ Discovery = (*MultiServersDiscovery)(nil)

// Refresh 实现接口，暂时空实现
func (d *MultiServersDiscovery) Refresh() error {
	return nil
}

// Update 更新服务器地址列表
func (d *MultiServersDiscovery) Update(servers []string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.servers = servers
	return nil
}

// Get 根据负载均衡策略选择一个服务器地址
func (d *MultiServersDiscovery) Get(mode SelectMode) (string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	n := len(d.servers)
	// 没有可用的服务
	if n == 0 {
		return "", errors.New("rpc discovery: no available servers")
	}
	// 从以定义的服务进行查找
	switch mode {
	case RandomSelect:
		// 随机策略，直接随机性返回一个服务器地址
		return d.servers[d.r.Intn(n)], nil
	case RoundRobinSelect:
		// 轮询策略，每次返回下一个 server 从末尾回绕
		s := d.servers[d.index%n]
		d.index = (d.index + 1) % n
		return s, nil
	default:
		return "", errors.New("rpc discovery: not supported select mode")
	}
}

// GetAll 返回所有服务器地址
func (d *MultiServersDiscovery) GetAll() ([]string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	// 返回一份拷贝
	servers := make([]string, len(d.servers), len(d.servers))
	copy(servers, d.servers)
	return servers, nil
}
