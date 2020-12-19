package namenode

import (
	"sync"
	"time"
)

// todo add mutex to handle concurrency

// LockManagerInterface exposes API that grants lease to the client for a file
type LockManagerInterface interface {
	GetNewLockManager()
	Renew()
	revoke()
	Grant()
	HasLock()
	monitor()
}

// LockManager asdf
type LockManager struct {
	fileToMetaMap map[string]fileLockMeta
	mu            sync.Mutex
}

type fileLockMeta struct {
	client    string
	timeStamp int64
}

// HasLock checks if file is locked before giving read permission
func (lm *LockManager) HasLock(file string) bool {
	lm.mu.Lock()
	defer lm.mu.Unlock()
	_, present := lm.fileToMetaMap[file]
	if present {
		return true
	}
	return false
}

func (lm *LockManager) revoke(client string, file string) {
	lm.mu.Lock()
	defer lm.mu.Unlock()
	delete(lm.fileToMetaMap, file)
}

// Renew is called to Renew the lock for a client on a file
func (lm *LockManager) Renew(client string, file string) bool {
	lm.mu.Lock()
	defer lm.mu.Unlock()
	fileMeta, present := lm.fileToMetaMap[file]
	if present {
		if fileMeta.client == client {
			meta := fileLockMeta{client: client, timeStamp: time.Now().Unix()}
			lm.fileToMetaMap[file] = meta
			return true
		}
		return false
	}
	return false
}

// Grant is called to Grant the lock for a client on a file
func (lm *LockManager) Grant(client string, file string) bool {
	lm.mu.Lock()
	defer lm.mu.Unlock()
	_, present := lm.fileToMetaMap[file]
	if present {
		return false
	}
	meta := fileLockMeta{client: client, timeStamp: time.Now().Unix()}
	lm.fileToMetaMap[file] = meta
	return true
}

// seperate gorutines and checks every 10 mins if
// lock is expired and revokes frees the client
func (lm *LockManager) monitor() {
	delay := 5 * time.Minute
	time.Sleep(delay)
	lm.mu.Lock()
	defer lm.mu.Unlock()
	for file, fileMeta := range lm.fileToMetaMap {
		if time.Since(time.Unix(fileMeta.timeStamp, 0)) > delay {
			lm.revoke(fileMeta.client, file)
		}
	}
	lm.monitor()
}

// GetNewLockManager returns reference to initilized LockManager
func GetNewLockManager() *LockManager {
	fileToMetaMap := make(map[string]fileLockMeta)
	lm := LockManager{fileToMetaMap: fileToMetaMap}
	go lm.monitor()
	return &lm
}
