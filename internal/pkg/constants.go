package pkg

const (
	// common constant names for messagebus.Optional properties

	// Client identifier configurations
	Username = "Username"
	Password = "Password"
	ClientId = "ClientId"

	// Connection configuration names
	ConnectTimeout = "ConnectTimeout"
	AutoReconnect  = "AutoReconnect"

	// TLS configuration names
	SkipCertVerify = "SkipCertVerify"
	CertFile       = "CertFile"
	KeyFile        = "KeyFile"
	KeyPEMBlock    = "KeyPEMBlock"
	CertPEMBlock   = "CertPEMBlock"

	// MQTT Specifics
	Qos          = "Qos"
	KeepAlive    = "KeepAlive"
	Retained     = "Retained"
	CleanSession = "CleanSession"

	// NATS specifics
	RetryOnFailedConnect = "RetryOnFailedConnect"
	Format               = "Format"
	QueueGroup           = "QueueGroup"

	// NATS JetStream specifics
	Durable                 = "Durable"
	AutoProvision           = "AutoProvision"
	Deliver                 = "Deliver"
	DefaultPubRetryAttempts = "DefaultPubRetryAttempts"
)
