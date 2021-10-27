package example

import (
	"fmt"
	"os"

	"github.com/micro/services/clients/go/user"
)

// Read an account by id, username or email. Only one need to be specified.
func ReadAnAccountById() {
	userService := user.NewUserService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := userService.Read(&user.ReadRequest{
		Id: "usrid-1",
	})
	fmt.Println(rsp, err)
}
