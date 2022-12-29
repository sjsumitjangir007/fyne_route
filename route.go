package fyne_route

import "fyne.io/fyne/v2"

func init() {
	// Initialization of route
}

// Page structure
type Page struct {
	container *fyne.Container
}

// Route
type Route struct {
	container *fyne.Container
	pages     map[string]Page
}

func (r *Route) SetContainer(container *fyne.Container) *fyne.Container {
	r.container = container
	return r.container
}

func (r *Route) AddPage(state string, container *fyne.Container) {
	if r.pages == nil {
		r.pages = map[string]Page{}
	}
	r.pages[state] = Page{
		container: container,
	}
}

func (r *Route) Start() {
	for key, value := range r.pages {
		if key == "/" {
			r.container.Add(value.container)
		}
	}
}

func (r *Route) Load(path string) {
	r.container.RemoveAll()
	for key, value := range r.pages {
		if key == path {
			r.container.Add(value.container)
		}
	}
	r.container.Refresh()
}
