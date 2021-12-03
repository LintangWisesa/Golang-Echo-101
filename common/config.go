package common

type Configuration struct {
	JwtIssuer         string
	JwtSignature      string
	JwtMinutesExpired string
}

var Config = Configuration{JwtIssuer: "Simplinnovation", JwtSignature: "321@noitavonnilpmiS", JwtMinutesExpired: "60"}
