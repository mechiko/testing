package entity

type Task interface {
	Run()
	Info() TaskInfo
	Clear()
	Drop() error
}

type TaskList interface {
	Add(Task) error
	Del(Task) error
	Clear()
	List() []Task
	Load() error
}

type TaskInfo struct {
	Start string
	End   string
	Since string
	Name  string
	State string
	Stage string
	Rem   string
	Ticks string
	Id    int64
}

type TasksItem struct {
	Task
	Id             int64  `json:"id"`
	Automatic      int    `json:"automatic"`
	Fsrar_Id       string `json:"fsrar_id"`
	Name           string `json:"name"`
	Start_On       string `json:"start_on"`
	After_Of       string `json:"after_of"`
	If_Exists      string `json:"if_exists"`
	Started        string `json:"started"`
	Ended          string `json:"ended"`
	Active         int    `json:"active"`
	Error          string `json:"error"`
	Items          int    `json:"items"`
	Progress       int    `json:"progress"`
	Progress_Error int    `json:"progress_error"`
	Ticks          int64  `json:"ticks"`
	State_Txt      string `json:"state_txt"`
	State_Json     string `json:"state_json"`
	Stage          int    `json:"stage"`
}

type TasksList struct {
	Total int64       `json:"total"`
	Items []TasksItem `json:"rows"`
}

// это интерфейс к таблице в БД
type Tasks interface {
	Insert(*TasksItem) error
	GetById(int64) (*TasksItem, error)
	GetAll() (*TasksList, error)
	GetActive() (*TasksList, error)
	Update(*TasksItem) error
	Delete(ti *TasksItem) error
	UpdateJsonById(js string, id int64) (result error)
	GetJsonById(id int64) (js string, result error)
}
