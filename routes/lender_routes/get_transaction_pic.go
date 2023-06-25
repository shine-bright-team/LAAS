package lender_routes

import (
	"github.com/gofiber/fiber/v2"
)

func GetTransactionPic(c *fiber.Ctx) error {
	return c.Redirect("https://www.mydegage.com/wp-content/uploads/2020/12/SCB-slip.jpg")

	//i := strings.Index(img, ",")
	//if i < 0 {
	//	log.Fatal("no comma")
	//}
	//dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(img[i+1:]))
	//f, err := os.Create("image.png")
	//io.Copy(f, dec)
	//return c.Send(dec.Read())

}
