package promexporter

import "github.com/prometheus/client_golang/prometheus"

var (
	MessageReceiveCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "message_receive_count",
		Help: "Total number of received SMS",
	})

	MessageSendCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "message_send_count",
		Help: "Total number of sent messsage",
	})

	DatabaseMessagesGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "database_messages_count",
		Help: "Total number of messages in database",
	})

	ErrMessageReceiveCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "message_error_receive_count",
		Help: "Total number of error received SMS",
	})

	ErrMessageSendCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "message_error_send_count",
		Help: "Total number of error sent messsage",
	})

	ErrDatabaseCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "message_error_database_count",
		Help: "Total number of error with database",
	})

	ErrSendSMSCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "send_sms_error_count",
		Help: "Total number of error with sending sms",
	})
)
