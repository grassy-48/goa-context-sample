// +build !test

package env

func Initialize() {
	OnDevelopment = true
	configfilename = "local"
}
