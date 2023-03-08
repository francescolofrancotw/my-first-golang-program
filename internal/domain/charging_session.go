package domain

// Unused Example

type chargingSessionStatus int

const (
	Started chargingSessionStatus = iota
	Ended   chargingSessionStatus = iota
)

type ChargingSession struct {
	status chargingSessionStatus
}

func StartChargingSession() ChargingSession {
	return ChargingSession{
		status: Started,
	}
}

func (cs ChargingSession) EndChargingSession() ChargingSession {
	if cs.status != Started {
		panic("Status must be started")
	}

	cs.status = Ended

	return cs
}
