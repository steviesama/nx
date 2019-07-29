package database

// ConnectionInfo contains the fields required to build a database connection
// string which is used to connect to the database.
type ConnectionInfo struct {
	DbIdentity   Identity `json:"DbIdentity"`
	Username     string   `json:"Username"`
	Password     string   `json:"Password"`
	Url          string   `json:"Url"`
	Port         int      `json:"Port"`
	DbName       string   `json:"DbName"`
	MaxIdleConns int      `json:"MaxIdleConns"`
	MaxOpenConns int      `json:"MaxOpenConns"`
}

// Init assigns the intended default values on the ConnectionInfo instance.
func (ci *ConnectionInfo) Init() {
	ci.DbIdentity = ""
	ci.Username = ""
	ci.Password = ""
	ci.Url = ""
	ci.Port = 0
	ci.DbName = ""
	ci.MaxIdleConns = 0
	ci.MaxOpenConns = 151
}
