package container

// 用于传递资源限制配置的结构体，包括内存限制，cpu时间片权重，cpu核心数
type ResourceConfig struct {
	MemoryLimit string
	CpuShare    string
	CpuSet      string
}

// subsystem接口
type Subsystem interface {
	Name() string                               //subsystem名称，例如cpu,memory
	Set(path string, res *ResourceConfig) error //设置某个cgroup在这个subsystem中的资源限制
	Apply(path string, pid int) error           //将进程加入某个cgroup中
	Remove(path string) error                   //移除某个cgroup
}

var (
	SubsystemsIns = []Subsystem{
		&MemorySubSystem{},
	}
)
