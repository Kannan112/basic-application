package domain

type Ticket struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	ClinicID uint   `json:"clinic_id"`
	Clinic   Clinic `json:"clinic" gorm:"foreignKey:ClinicID"`
	UHCD     string `json:"uhcd,omitempty"`
	TimeUTC  string `json:"time_utc"`
}

type Clinic struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}
