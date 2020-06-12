package main

import (
)

type User struct {
    ID      string `gorm:"primary_key" json:"id"`
    Username  string `gorm:"primary_key" json:"username"`
}

type Version struct {
    Version string `json:"version"`
}

func (s *User) TableName() string {
    return "auth_user"
}

func gormConnect() *gorm.DB {
    DBMS     := "mysql"
    USER     := "root"
    PASS     := "rooooot"
    PROTOCOL := "tcp(172.20.10.4:3306)"
    DBNAME   := "handson"

    CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
    db,err := gorm.Open(DBMS, CONNECT)

    if err != nil {
      panic(err.Error())
    }
    return db
}

func main() {
    listener, err := net.Listen("tcp", ":0")
    if err != nil {
        panic(err)
    }

    var port_number int = listener.Addr().(*net.TCPAddr).Port+1
    var port_number_str string

    port_number_str = strconv.Itoa(port_number)

    fmt.Println("Using port:", listener.Addr().(*net.TCPAddr).Port+1)

    config := consulapi.DefaultConfig()
    consul, err := consulapi.NewClient(config)
    if err != nil {
    log.Fatalln(err)
    }
    registration := new(consulapi.AgentServiceRegistration)
    registration.ID = "api-go-"+port_number_str //replace with service id 
    registration.Name = "api-go" //replace with service name
    address := "127.0.0.1"
    registration.Address = address
    registration.Port = port_number
    registration.Check = new(consulapi.AgentServiceCheck)
    registration.Check.HTTP = fmt.Sprintf("http://%s:%v/users",
    address, port_number_str)
    registration.Check.Interval = "5s"
    registration.Check.Timeout = "3s"
    registration.Check.DeregisterCriticalServiceAfter = "1m"
    consul.Agent().ServiceRegister(registration)

    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)
    router, err := rest.MakeRouter(
        rest.Get("/users", GetAllUsers),
        rest.Get("/version", GetVersion),
    )
    if err != nil {
        log.Fatal(err)
    }

    api.SetApp(router)

    log.Fatal(http.ListenAndServe(":"+port_number_str, api.MakeHandler()))

    fmt.Println("Exit!!!")
}

func GetAllUsers(w rest.ResponseWriter, r *rest.Request) {
    db := gormConnect()
    defer db.Close()

    // 全件取得
    var allUsers []User
    db.Find(&allUsers)
    fmt.Println(allUsers)

    w.WriteHeader(http.StatusOK)
    w.WriteJson(&allUsers)
}

func GetVersion(w rest.ResponseWriter, r *rest.Request) {
    version := Version{"v3.0.0"}

    w.WriteJson(version)
}