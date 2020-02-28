package ecs

// Engine is simple a composition of an EntityManager and a SystemManager.
// It handles the stages Setup(), Run() and Teardown() for all the systems.
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

// Run calls the Process() method for each System
// until ShouldEngineStop is set to true.
func (e *Engine) Run() {
	systems, isPaused := e.systemManager.Systems()
	for name, system := range systems {
		if isPaused[name] {
			system.Process(e.entityManager)
		}
	}
}

// Setup calls the Setup() method for each System
// and initializes ShouldEngineStop and ShouldEnginePause with false.
func (e *Engine) Setup() {
	systems, _ := e.systemManager.Systems()
	for _, sys := range systems {
		sys.Setup()
	}
}

// Teardown calls the Teardown() method for each System.
func (e *Engine) Teardown() {
	systems, _ := e.systemManager.Systems()
	for _, sys := range systems {
		sys.Teardown()
	}
}


func (e *Engine) PauseSystem(systemType string) {
	e.systemManager.Pause(systemType)
}

