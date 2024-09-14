package check

import (
	"peterszarvas94/blog/pkg/custom"
	"peterszarvas94/blog/pkg/errors"
	"peterszarvas94/blog/pkg/fileutils"
)

var protectedRoutes = []string{"/404", "/static", "/tag", "/category"}

func CheckContentDir() error {
	for route := range *custom.Routes {
		protectedRoutes = append(protectedRoutes, route)
	}

	files := fileutils.GetFiles()
	for _, file := range files {
		for _, protectedRoute := range protectedRoutes {
			if file.Fileroute == protectedRoute {
				return &errors.ProtectedRouteError{
					Filename: file.Path,
					Route:    file.Fileroute,
				}
			}
		}
	}

	return nil
}
