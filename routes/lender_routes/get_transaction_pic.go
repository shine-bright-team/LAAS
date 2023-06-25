package lender_routes

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"image"
	"image/jpeg"
	"net/http"
)

func GetTransactionPic(c *fiber.Ctx) error {
	resp, _ := http.Get("https://www.mydegage.com/wp-content/uploads/2020/12/SCB-slip.jpg")
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return c.Status(fiber.StatusInternalServerError).SendString("There is an error from our side please try again later")
	}

	img, _, _ := image.Decode(resp.Body)

	buf := new(bytes.Buffer)

	jpeg.Encode(buf, img, nil)
	c.Set("Content-Type", "image/jpeg")
	return c.Send(buf.Bytes())
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
