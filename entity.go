package ecs

// Component contains only the data (no behaviour at all).
// The Name() method must be implemented, because the EntityManager
// uses it to filter the entities by component names.
type Component interface {
	Name() (name string)
}

// entity is simply a composition of one or more components with an id.
type entity struct {
	Components map[string]Component
	Tags       map[string]bool
	Name       string
	ID         int64
}

// Checks for existence of a component
func (e *entity) Has(name string) bool {
	if _, contains := e.Components[name]; contains {
		return true
	}

	return false
}

// Get a specific component by name, will return nil if the component doesn't exists
func (e *entity) Get(name string) Component {
	if e.Has(name) {
		return e.Components[name]
	}

	return nil
}

// Add a component.
func (e *entity) Add(components ...Component) {
	for _, component := range components {
		e.Components[component.Name()] = component
	}
}

func (e *entity) AddTag(tags ...string) {
	for _, tag := range tags {
		if !e.HasTag(tag) {
			e.Tags[tag] = true
		}
	}
}

func (e *entity) Remove(component string) {
	if e.Has(component) {
		delete(e.Components, component)
	}
}

func (e *entity) RemoveTag(tag string) {
	if e.HasTag(tag) {
		e.Tags[tag] = false
	}
}

func (e *entity) HasTag(name string) bool {
	if _, contains := e.Tags[name]; contains {
		return e.Tags[name]
	}

	return false
}
