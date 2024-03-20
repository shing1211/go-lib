package config

type RabbitMQConfig struct {
	RabbitMQHost string
	RabbitMQPort string
	RabbitMQUser string
	RabbitMQPwd  string
}

type RabbitMQPublisherConfig struct {
	RabbitMQPublishAPIHost               string
	RabbitMQPublishAPIPort               string
	RabbitMQPublishAPIEndpointPrefix     string
	RabbitMQPublishAPIEndpointKline      string
	RabbitMQPublishAPIEndpointAggTrade   string
	RabbitMQPublishAPIEndpointPriceDepth string
}
