package wasm

import "github.com/tetratelabs/wazero/api"

func (m *MemoryInstance) Backup() *api.MemoryBackup {
	backup := api.MemoryBackup{
		Min:    m.Min,
		Max:    m.Max,
		Cap:    m.Cap,
		Buffer: make([]byte, len(m.Buffer)),
	}
	copy(backup.Buffer[0:], m.Buffer[0:])
	return &backup
}

// Restore restores the memory state from backup
func (m *MemoryInstance) Restore(backup *api.MemoryBackup) {
	m.Max = backup.Max
	m.Min = backup.Min
	m.Cap = backup.Cap
	m.Buffer = make([]byte, len(backup.Buffer))
	copy(m.Buffer[0:], backup.Buffer[0:])
}
