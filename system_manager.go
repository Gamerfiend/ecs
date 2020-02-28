package ecs

// SystemManager handles the access to each system.
type SystemManager struct {
	systems map[string]System
	isPaused map[string]bool
}

// NewSystemManager creates a new SystemManager and returns its address.
func NewSystemManager() *SystemManager {
	return &SystemManager{
		systems: map[string]System{},
		isPaused: map[string]bool{},
	}
}

// Add systems to the SystemManager.
func (m *SystemManager) Add(systems ...System) {
	for _, system := range systems {
		m.systems[system.Name()] = system
		m.isPaused[system.Name()] = false
	}
}

func (m *SystemManager) Pause(systemName string) {
	if _, exists := m.systems[systemName]; exists {
		m.isPaused[systemName] = true
	}
}

func (m *SystemManager) Resume(systemName string) {
	if _, exists := m.systems[systemName]; exists {
		m.isPaused[systemName] = false
	}
}

// Systems returns the system, which are internally stored.
func (m *SystemManager) Systems() (map[string]System, map[string]bool) {
	return m.systems, m.isPaused
}
