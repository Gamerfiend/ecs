package ecs

// SystemManager handles the access to each system.
type SystemManager struct {
	systems []System
	isPaused map[string]bool
}

// NewSystemManager creates a new SystemManager and returns its address.
func NewSystemManager() *SystemManager {
	return &SystemManager{
		systems: []System{},
		isPaused: map[string]bool{},
	}
}

// Add systems to the SystemManager.
func (m *SystemManager) Add(systems ...System) {
	for _, system := range systems {
		m.systems = append(m.systems, system)
		m.isPaused[system.Name()] = false
	}
}

func (m *SystemManager) Pause(systemName string) {
	m.isPaused[systemName] = true
}

func (m *SystemManager) Resume(systemName string) {
	m.isPaused[systemName] = false
}

// Systems returns the system, which are internally stored.
func (m *SystemManager) Systems() ([]System, map[string]bool) {
	return m.systems, m.isPaused
}
