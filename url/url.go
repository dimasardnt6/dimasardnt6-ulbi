package url

import (
	"github.com/dimasardnt6/dimasardnt6-ulbi/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Home)
	page.Get("/kemahasiswaan", controller.GetKemahasiswaan)
	page.Get("/data-mahasiswa", controller.GetDataMahasiswa)
	page.Get("/keuangan-mahasiswa", controller.GetKeuanganMahasiswa)
	page.Get("/nilai-mahasiswa", controller.GetNilaiMahasiswa)
	page.Post("/insrt-kemahaiswaan", controller.InsertKemahasiswaan)
	page.Post("/insrt-mahasiswa", controller.InsertDataMahasiswa)
	page.Post("/insrt-keuangan", controller.InsertKeuanganMahasiswa)
	page.Post("/insrt-nilai", controller.InsertNilaiMahasiswa)
}
