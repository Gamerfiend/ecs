package ecs_test

import (
	"github.com/andygeiss/assert"
	"github.com/andygeiss/assert/is"
	"github.com/andygeiss/ecs"
	"testing"
)

func TestSystemManager_Systems_Should_Have_No_Entity_At_Start(t *testing.T) {
	m := ecs.NewSystemManager()
	systems, _ := m.Systems()
	assert.That(t, len(systems), is.Equal(0))
}

type MockSystem struct {}

func (s *MockSystem) Name() string {
	return "MockSystem"
}
func (s *MockSystem) Process(entityManager *ecs.EntityManager) {}
func (s *MockSystem) Setup()                                   {}
func (s *MockSystem) Teardown()                                {}

func TestSystemManager_Systems_Should_Have_One_System_After_Adding_One_System(t *testing.T) {
	m := ecs.NewSystemManager()
	s := &MockSystem{}
	m.Add(s)
	systems, _ := m.Systems()
	assert.That(t, len(systems), is.Equal(1))
	assert.That(t, systems[0], is.Equal(s))
}

func TestSystemManager_Systems_Should_Have_Two_System_After_Adding_Two_System(t *testing.T) {
	m := ecs.NewSystemManager()
	s1 := &MockSystem{}
	s2 := &MockSystem{}
	m.Add(s1, s2)
	systems, _ := m.Systems()
	assert.That(t, len(systems), is.Equal(2))
	assert.That(t, systems[0], is.Equal(s1))
	assert.That(t, systems[1], is.Equal(s2))
}
