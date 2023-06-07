package main

import "fmt"

type Manager struct {
	name  string
	field string
}

type EngineerLeader struct {
	name     string
	teamSize int
}

type Managment interface {
	peopleManagement() string
}

func (m *Manager) peopleManagement() {
	fmt.Println("Manager is managing ops people")
}

type EngineerManager struct {
	manager   Manager
	engLeader EngineerLeader
}

func (m *EngineerManager) peopleManagement() {
	fmt.Println("Manager is managing developers, my name is", m.manager.name, "and I work together with", m.engLeader.teamSize)
}

func decorator() {
	engManager := EngineerManager{
		Manager{
			"TestManager",
			"Engineering",
		},
		EngineerLeader{teamSize: 10},
	}
	engManager.peopleManagement()
}
