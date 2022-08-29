package main

import (
	"startup/handler"
	"startup/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
  dsn := "root:74712331@tcp(127.0.0.1:3306)/startup?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

  if err != nil {
    panic(err)
  }

  userRepository := user.NewRepository(db)
  userService := user.NewService(userRepository)
  userHandler := handler.NewUserHandler(userService)

  // input := user.LoginInput{
  //   Email: "zazhil249@gmail.com",
  //   Password: "password",
  // }

  // user, err := userService.Login(input)
  // if err != nil {
  //   fmt.Println("Terjadi kesalahan")
  //   fmt.Println(err.Error())
  // } 

  // fmt.Println(user.Email)
  // fmt.Println(user.Name)


  // userByEmail, err := userRepository.FindByEmail("zazhil24@gmail.com")
  // if err != nil {
  //   fmt.Println(err.Error())
  // }
  
  // if (userByEmail.ID == 0) {
  //   fmt.Println("User tidak ditemukan")
  // } else {
  //   fmt.Println(userByEmail.Name)
  // }

  router := gin.Default()

  api := router.Group("/api/v1")

  api.POST("/users", userHandler.RegisterUser)
  api.POST("/sessions", userHandler.Login)

  router.Run()


  // userInput := user.RegisterUserInput{}
  // userInput.Name = "Tes simpan dari service"
  // userInput.Email = "tes@gmail.com"
  // userInput.Occupation = "tukang"
  // userInput.Password = "password"

  // userService.RegisterUser(userInput)


  // userRepository.Save(user)

  // fmt.Println("Connection success!")
  // for _, user := range users {
  //   fmt.Println(user.Name)
  //   fmt.Println(user.Occupation)
  //   fmt.Println("=============>")
  // }

  // router.GET("/handler", handler)
  // router.Run()
}
//  func handler(c *gin.Context)  {
//   dsn := "root:74712331@tcp(127.0.0.1:3306)/startup?charset=utf8mb4&parseTime=True&loc=Local"
//   db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

//   if err != nil {
//     panic(err)
//   }

//   var users []user.User

//   db.Find(&users)

//   c.JSON(http.StatusOK, users)
//  }