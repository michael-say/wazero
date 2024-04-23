package wasm

import (
	"testing"

	"github.com/tetratelabs/wazero/internal/testing/require"
)

func TestMemoryInstance_Backup_Restore(t *testing.T) {
	m := &MemoryInstance{Max: 0xff, Buffer: make([]byte, 0)}

	reqInitialState := func() {
		require.Equal(t, uint32(0), m.Min)
		require.Equal(t, uint32(0xff), m.Max)
		require.Equal(t, uint32(1), m.Cap)
		require.Equal(t, int(1*MemoryPageSize), len(m.Buffer))
		require.Equal(t, uint8(1), m.Buffer[0])
	}

	result, ok := m.Grow(1)
	require.True(t, ok)
	require.Equal(t, uint32(0), result)

	m.Buffer[0] = 1
	reqInitialState()

	backup := m.Backup()

	result, ok = m.Grow(1)
	require.True(t, ok)
	require.Equal(t, uint32(1), result)

	m.Buffer[0] = 0xff
	require.Equal(t, uint32(2), m.Cap)
	require.Equal(t, uint8(0xff), m.Buffer[0])
	require.Equal(t, int(2*MemoryPageSize), len(m.Buffer))

	m.Restore(backup)
	reqInitialState()

}
