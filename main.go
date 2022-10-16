package main

import (
	"github.com/Package/onedrive"
	"github.com/Package/system"
)

func main() {
	onedrive.OneDriver("test.txt", "https://1drv.ms/t/s!Aq1pYjY0sJwZgZtZzj2f2hjK9lDQ")
	system.CheckOS()
}
