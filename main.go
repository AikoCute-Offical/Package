package main

import "github.com/AikoCute-Offical/GoPackage/onedrive"

// DownloadFile download file from onedrive
func OneDriver(patch string, url string) {
	onedrive.DownloadFile(patch, url)
}

func main() {
	OneDriver("test.txt", "https://1drv.ms/t/s!Aq1pYjY0sJwZgZtZzj2f2hjK9lDQ")
}
