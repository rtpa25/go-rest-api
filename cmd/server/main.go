package main

import "fmt"

//*App - the struct which contains things like pointers to db connections
type App struct{}

//*Run- sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up our App")
	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error setting up the app")
		fmt.Println(err)
	}
}
