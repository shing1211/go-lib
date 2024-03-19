package config

// Define a custom type for application's deployment environment
type DeployEnv struct {
	Dev string
	Prd string
}

// Create a variable of the struct type to represent the enum
var Env DeployEnv = DeployEnv{
	Dev: "development",
	Prd: "production",
}
