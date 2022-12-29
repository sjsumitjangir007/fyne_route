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
	Container *fyne.Container
	Pages     map[string]Page
}

func (r *Route) Start() *fyne.Container {
	for key, value := range r.Pages {
		if key == "/" {
			r.Container.Add(value.Container)
		}
	}
	return r.Container
}

func (r *Route) Load(path string) {
	r.Container.RemoveAll()
	for key, value := range r.Pages {
		if key == path {
			r.Container.Add(value.Container)
		}
	}
	r.Container.Refresh()
}
