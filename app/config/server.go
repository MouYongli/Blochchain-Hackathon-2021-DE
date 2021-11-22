package config

type Server struct {
	Host 		string	`json:"host"`
	Port 		int		`json:"port"`
	Model 		string 	`json:"model"`
	Uploads		string	`json:"uploads"`
	Downloads	string	`json:"downloads"`
}
