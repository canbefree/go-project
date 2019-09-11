package mydb

//UserActive 用户表
type UserActive struct {
	user User
	ActiveDb
}

var activeDb = ActiveDb{
	dsn:       "",
	connected: false,
	db:        nil,
}

func newUserActive() *UserActive {
	return &UserActive{
		ActiveDb:activeDb,
	}
}


func (active *UserActive) getTableName() string {
	return "t_user"
}

type User struct {
	id         int32
	user       string
	pwd        string
	created_at string
	updated_at string
}

func GetOne1() {
	userActive := newUserActive()
	userActive.QueryOne()
}
