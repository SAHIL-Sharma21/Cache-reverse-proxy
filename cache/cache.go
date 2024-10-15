package cache

var Cache map[string]string

func Init() {
	Cache = make(map[string]string)
}
