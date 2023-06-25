package lender_routes

import (
	"github.com/gofiber/fiber/v2"
)

func GetTransactionPic(c *fiber.Ctx) error {
	return c.Redirect("https://www.mydegage.com/wp-content/uploads/2020/12/SCB-slip.jpg")
	//resp, _ := http.Get("https://cloud.thistine.com/apps/files_sharing/publicpreview/XeNPMkDS8kJMbzw?file=/&fileId=1850&x=1920&y=1080&a=true")
	//defer resp.Body.Close()
	//
	//if body, err := ioutil.ReadAll(resp.Body); err != nil {
	//	return c.Status(fiber.StatusInternalServerError).SendString("There is an error from our side please try again later")
	//} else {
	//	return c.Send(body)
	//}

	//i := strings.Index(img, ",")
	//if i < 0 {
	//	log.Fatal("no comma")
	//}
	//dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(img[i+1:]))
	//f, err := os.Create("image.png")
	//io.Copy(f, dec)
	//return c.Send(dec.Read())

}
