package laji

func diff_about_var_type() {
	var e struct {
		name string
	}
	type f struct {
		name string
	}
	e.name = "18"
	var g f
	g.name = "19"
}

func what_make_do() {
	type student struct {
		name string
		age  string
	}

	stus := make(map[string]student)
	stus["12"] = student{name: "小明", age: "19"}
}

func what_new_do() {
	i := new(int)
	*i = 12
}

//适配类型
type ilog interface {
	getLog()
}

//适配类型
type jlog interface {
	getLog()
}

type dbLog struct {
}

func (log dbLog) getLog() {
}

type cacheLog struct {
}

func (log cacheLog) getLog() {
}

func getLog(t int) interface{} {
	if t == 1 {
		return dbLog{}
	}
	if t == 2 {
		return cacheLog{}
	}
	return nil
}

func what_point_do() {
	//i 类型为 interface
	var i interface{} = 12
	if _, ok := i.(int); !ok {
		panic("error")
	}
	log := getLog(3)
	//如果类型为nil不会报错
	if vs, ok := log.(ilog); !ok {
		panic("error2")
	}else{
		vs.getLog()
	}
}
