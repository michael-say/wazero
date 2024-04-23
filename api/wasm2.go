package api

type MemoryBackup struct {
	Buffer        []byte
	Min, Cap, Max uint32
}

type RecoverableMemory interface {

	// Backup returns a complete memory backup
	Backup() *MemoryBackup

	// Restore restores the memory state from backup
	Restore(backup *MemoryBackup)
}
