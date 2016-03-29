package executil

import "sync"

type multifile struct {
	sync.Mutex
}
