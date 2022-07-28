package main

import "app"

func init() {

}
func main() {
	println("\r\n[VERSION INFO]")
	println("-----------------------------------------------------------------")
	println("BINARY:", app.Binary)
	println("LANGUAGE:", app.Language)
	println("LANGUAGE_VERSION:", app.LanguageVersion)
	println("OS:", app.Os)
	println("ARCH:", app.Arch)
	println("REPO:", app.Repository)
	println("VERSION:", app.Version)
	println("BUID_BY:", app.Build)
	println("BUID_DATE:", app.BuildDate)
}
