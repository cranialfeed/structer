package main

import (
	"github.com/cranialfeed/structer"
	"github.com/joho/godotenv"
)

func main() {

	// load environment variables
	godotenv.Load("./example/example.env")

	// instantiate struct
	user := &User{}

	// load env and default values
	structer.SetDefaults(user)

	// output object to console
	structer.PrettyPrint(user)

}

type ContactType struct {
	ID      int    `default:"1" struct:"id"`
	Caption string `default:"Mobile Phone Number" json:"number"`
	TypeID  string ` json:"type"`
}
type Contact struct {
	ID      int    `json:"id"`
	Caption string `json:"caption"`
	Type    string `json:"type"`
}
type User struct {
	ID               int       `default:"0" env:"EmployeeID" json:"id"`
	Name             string    `default:"" env:"EmployeeName" json:"name"`
	EmploymentStatus string    `default:"Employed" json:"employment_status"`
	Contacts         []Contact `default:"[{\"ID\":4, \"Caption\":\"817-273-3746\", \"Type\":\"Mobile Phone\"},{\"ID\":5, \"Caption\":\"hpotter@hogwartz.com\", \"Type\":\"Email\"}]" json:"contacts"`
}
