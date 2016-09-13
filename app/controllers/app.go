package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	//var dict map[string]string
	//dict["greeting"] = "Hello, World! from Go"
	greeting := "Hello, World! from Go"
	//return c.Render(dict)
	return c.Render(greeting)
}

func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("Your name is required!")
	c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}
	return c.Render(myName)
}
