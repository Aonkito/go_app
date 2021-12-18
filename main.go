package main

import "todo_app/app/controllers"

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	// fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.Password = "testtest"
	// u.CreateUser()

	// user, err := models.GetUser(2)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println("Name:", user.Name, "\nEmail:", user.Email)

	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println("Name:", u.Name, "\nEmail:", u.Email)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println("Name:", u.Name, "\nEmail:", u.Email)

	// todos, _ := user.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// todo, _ := models.GetTodo(1)
	// todo.Content = "New Content"
	// todo.UpdateTodo()
	// todo, _ = models.GetTodo(1)
	// fmt.Println(todo)
	// todo, _ := models.GetTodo(1)
	// todo.DeleteTodo()

	controllers.StartMainServer()
	// user, _ := models.GetUserByEmail("test@baobao.com")

	// fmt.Println(user)
	// session, _ := user.CreateSession()
	// fmt.Println(session)

	// valid, _ := session.CheckSession()
	// fmt.Println(valid)
}
