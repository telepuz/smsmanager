package promexporter

func (p *PromExporter) IncMessageReceiveCounter() {
	MessageReceiveCounter.Inc()
}

func (p *PromExporter) IncMessageSendCounter() {
	MessageSendCounter.Inc()
}

func (p *PromExporter) SetDatabaseMessagesGauge(n int) {
	DatabaseMessagesGauge.Set(float64(n))
}

func (p *PromExporter) IncErrMessageReceiveCounter() {
	ErrMessageReceiveCounter.Inc()
}

func (p *PromExporter) IncErrMessageSendCounter() {
	ErrMessageSendCounter.Inc()
}

func (p *PromExporter) IncErrDatabaseCounter() {
	ErrDatabaseCounter.Inc()
}
