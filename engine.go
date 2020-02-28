package ecs

// Engine is simple a composition of an EntityManager and a SystemManager.
// It handles the stages Setup(), Process() and Teardown() for all the systems.
type Engine struct {
	entityManager *EntityManager
	systemManager *SystemManager
}

// NewEngine creates a new Engine and returns its address.
func NewEngine(entityManager *EntityManager, systemManager *SystemManager) *Engine {
	return &Engine{
		entityManager: entityManager,
		systemManager: systemManager,
	}
}

// Run calls the Process() method for each System unless the System is paused.
func (e *Engine) Run() {
	for _, system := range e.systemManager.systems {
		if !e.systemManager.isPaused[system.Name()] {
			system.Process(e.entityManager)
		}
	}
}

// Setup calls the Setup() method for each System
func (e *Engine) Setup() {
	for _, sys := range e.systemManager.systems {
		sys.Setup()
	}
}

// Teardown calls the Teardown() method for each System.
func (e *Engine) Teardown() {
	for _, sys := range e.systemManager.systems {
		sys.Teardown()
	}
}


func (e *Engine) PauseSystem(systemType string) {
	e.systemManager.Pause(systemType)
}

