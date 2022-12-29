package route

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

func (r *Route) start() *fyne.Container {
	for key, value := range r.pages {
		if key == "/" {
			r.container.Add(value.container)
		}
	}
	return r.container
}

func (r *Route) load(path string) {
	r.container.RemoveAll()
	for key, value := range r.pages {
		if key == path {
			r.container.Add(value.container)
		}
	}
	r.container.Refresh()
}
