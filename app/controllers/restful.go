package controllers

type RestFul interface {
	List()
	Show()
	Store()
	Update()
	Destroy()
}
