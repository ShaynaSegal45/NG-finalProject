package user

type User struct {
	UserName string
	Password string

	//Gender string
	//Image image.Image
	//Preferences []string
}
type Profile struct {
	UserName string
	Gender string
	Age int
	//Image image.Image
	//Preferences []string
}
type Gender struct {
	Gender string
}

