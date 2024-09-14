package check

import (
	"peterszarvas94/blog/pkg/custom"
	"peterszarvas94/blog/pkg/errors"
	"peterszarvas94/blog/pkg/fileutils"
)

var protectedRoutes = []string{"/404", "/static", "/tag", "/category"}

func CheckContentDir() error {
	for route := range *custom.Routes {
		for _, protectedRoute := range protectedRoutes {
			if route == protectedRoute {
				return &errors.ExistingProtectedRouteError{
					Route: route,
				}
			}
		}
	}

	files := fileutils.GetFiles()

	for _, file := range files {
		for route := range *custom.Routes {
			if file.Fileroute == route {
				return &errors.ProtectedRouteError{
					Filename: file.Path,
					Route:    file.Fileroute,
					Kind:     "custom",
				}
			}
		}
	}

	for _, file := range files {
		for _, protectedRoute := range protectedRoutes {
			if file.Fileroute == protectedRoute {
				return &errors.ProtectedRouteError{
					Filename: file.Path,
					Route:    file.Fileroute,
					Kind:     "protected",
				}
			}
		}
	}

	return nil
}
