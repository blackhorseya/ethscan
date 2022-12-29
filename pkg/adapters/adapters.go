package adapters

// Restful restful adapters
type Restful interface {
	InitRouting() error
}

// Grpc grpc adapter
type Grpc interface {
	// Start to run
	Start() error

	// Stop to end
	Stop() error
}

// Kafka is a consumer adapters
type Kafka interface {
	// Start to listen event
	Start() error

	// Stop to listen
	Stop() error
}

// Cronjob is a cronjob adapters
type Cronjob interface {
	// Start to run
	Start() error

	// Stop to end
	Stop() error
}

// CLI is a command line interface adapters
type CLI interface {
	// Execute serve caller to execute command
	Execute() error
}
