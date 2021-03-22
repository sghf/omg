package drivergroup

//
// T groups drivers sharing some properties.
// A resourceset is a collection of resources having the same drivergroup and subset.
//
type T int

const (
	Unknown T = 1 << iota
	IP
	Volume
	Disk
	FS
	Share
	Container
	App
	Sync
	Task
)

func New(s string) T {
	switch s {
	case "ip":
		return IP
	case "volume":
		return Volume
	case "disk":
		return Disk
	case "fs":
		return FS
	case "share":
		return Share
	case "container":
		return Container
	case "app":
		return App
	case "sync":
		return Sync
	case "task":
		return Task
	default:
		return Unknown
	}
}

func (t T) Name() string {
	switch t {
	case IP:
		return "ip"
	case Volume:
		return "volume"
	case Disk:
		return "disk"
	case FS:
		return "fs"
	case Share:
		return "share"
	case Container:
		return "container"
	case App:
		return "app"
	case Sync:
		return "sync"
	case Task:
		return "task"
	default:
		return "unknown"
	}
}
