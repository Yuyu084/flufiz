package model

import "time"

type PetPropertyHealth struct {
	Value       int       `json:"value"`
	LastUpdated time.Time `json:"last_updated"`
}

func (p *PetPropertyHealth) GetCurrentValue() int {
	return p.Value
}

func (p *PetPropertyHealth) UpdateHealth(newValue int) {
	p.Value = newValue
	p.LastUpdated = time.Now()
}

// PetPropertyEnergy - энергичность питомца
type PetPropertyEnergy struct {
	Value       int       `json:"value"`
	LastUpdated time.Time `json:"last_updated"`
}

func (e *PetPropertyEnergy) GetCurrentValue() int {
	return e.Value
}

func (e *PetPropertyEnergy) UpdateEnergy(newValue int) {
	e.Value = newValue
	e.LastUpdated = time.Now()
}

// PetMood - настроение питомца
type PetMood struct {
	Mood        string    `json:"mood"`
	LastUpdated time.Time `json:"last_updated"`
}

const (
	MoodHappy = "happy"
	MoodSad   = "sad"
	MoodAngry = "angry"
)

func (m *PetMood) GetCurrentMood() string {
	return m.Mood
}

func (m *PetMood) SetMood(mood string) {
	m.Mood = mood
	m.LastUpdated = time.Now()
}

func (m *PetMood) GetMoodEmoji() string {
	switch m.Mood {
	case MoodHappy:
		return "😊"
	case MoodSad:
		return "😢"
	case MoodAngry:
		return "😠"
	default:
		return "😐"
	}
}
