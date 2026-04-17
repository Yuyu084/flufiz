package model

import "time"

type Pets struct {
Pets []Pet `json:"pets"`
}
type Pet struct {
Name           string            `json:"name"`
Type           PetType           `json:"type"`
PropertyHealth PetPropertyHealth `json:"property_health"`
PropertyEnergy PetPropertyEnergy `json:"property_energy"`
Mood           PetMood           `json:"mood"`
LastPlayed     time.Time         `json:"last_played"` 
}
type PetType int
const (
PetTypeCat PetType = iota
PetTypeDog
)
func (p *Pet) UpdateMoodByHealth() {
if p.PropertyHealth.Value <= 30 {
p.Mood.SetMood(MoodAngry)
} else if p.PropertyHealth.Value < 45 {
p.Mood.SetMood(MoodSad)
} else {
p.Mood.SetMood(MoodHappy)
}
}
