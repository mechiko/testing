package entity

type Repo interface {
	DbService() DbService
	Start() error
	CheckVersionDb() error
	// InsertRules() error
	// GetRequests() Requests
	// GetRules() Rules
	// GetNodesout() Nodesout
	// GetRequestRests() RequestRests
	// GetRests() Rests
	// GetClients() Clients
	// GetForm1() Form1
	// GetForm2() Form2
	// GetTasks() Tasks
	GetAppState() AppState
}
