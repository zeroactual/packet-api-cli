package util

type project struct {
	Id   string
	Name string
}

func GetProject() (*project, error) {
	var project project

	// Grab Project ID for later use
	err := handle("https://api.packet.net/project", GET, nil, &project)

	return &project, err
}