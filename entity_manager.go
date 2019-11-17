package ecs

// EntityManager handles the access to each entity.
type EntityManager struct {
	entities []*entity
	index []int
	entityID int64
}

// NewEntityManager creates a new EntityManager and returns its address.
func NewEntityManager() *EntityManager {
	return &EntityManager{
		entities: []*entity{},
	}
}

// Builds a new entity
func (m *EntityManager) NewEntity(name string) *entity {
	entity := &entity{Components: map[string]Component{}, Name: name, ID: m.getNextID(),}
	m.add(entity)
	return entity
}

func (m *EntityManager) getNextID() int64 {
	m.entityID = m.entityID + 1
	return m.entityID
}

// Add entries to the manager.
func (m *EntityManager) add(entity *entity) {
	m.entities = append(m.entities, entity)
}

// Entities returns all the entities.
func (m *EntityManager) Entities() (entities []*entity) {
	return m.entities
}

// FilterBy returns the mapped entities, which components name matched.
func (m *EntityManager) FilterBy(components ...string) (entities []*entity) {
	for _, e := range m.entities {
		count := 0
		wanted := len(components)
		// Simply increase the count if the component could be found.
		for _, name := range components {
			for _, c := range e.Components {
				if c.Name() == name {
					count++
				}
			}
		}
		// Add the entity to the filter list, if all components are found.
		if count == wanted {
			entities = append(entities, e)
		}
	}
	return
}

// Get a specific entity by id.
func (m *EntityManager) Get(name string) (entity *entity) {
	for _, e := range m.entities {
		if e.Name == name {
			return e
		}
	}
	return
}

// Remove a specific entity.
func (m *EntityManager) Remove(entity *entity) {
	for i, e := range m.entities {
		if e.Name == entity.Name {
			copy(m.entities[i:], m.entities[i+1:])
			m.entities[len(m.entities)-1] = nil
			m.entities = m.entities[:len(m.entities)-1]
			break
		}
	}
}
