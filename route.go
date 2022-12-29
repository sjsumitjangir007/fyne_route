package fyne_route

import "fyne.io/fyne/v2"

func init() {
	// Initialization of route
}

// Page structure
type Page struct {
	Container *fyne.Container
}

// Route
type Route struct {
	container *fyne.Container
	Pages     map[string]Page
}

func (r *Route) SetContainer(container *fyne.Container) *fyne.Container {
	r.container = container
	return r.container
}

func (r *Route) Start() {
	for key, value := range r.Pages {
		if key == "/" {
			r.container.Add(value.Container)
		}
	}
}

func (r *Route) Load(path string) {
	r.container.RemoveAll()
	for key, value := range r.Pages {
		if key == path {
			r.container.Add(value.Container)
		}
	}
	r.container.Refresh()
}
