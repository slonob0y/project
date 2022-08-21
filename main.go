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

  router := gin.Default()

  api := router.Group("/api/v1")

  api.POST("/users", userHandler.RegisterUser)

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