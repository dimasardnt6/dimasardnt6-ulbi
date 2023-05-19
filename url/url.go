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
	page.Get("/all-kemahasiswaan", controller.GetAllKemahasiswaan)
	page.Get("/all-mahasiswa", controller.GetAllDataMahasiswa)
	page.Get("/all-keuangan", controller.GetAllKeuanganMahasiswa)
	page.Get("/all-nilai", controller.GetAllNilaiMahasiswa)
	page.Get("/kemahasiswaan", controller.GetKemahasiswaanNPM)
	page.Get("/data-mahasiswa", controller.GetDataMahasiswa)
	page.Get("/keuangan-mahasiswa", controller.GetKeuanganMahasiswa)
	page.Get("/nilai-mahasiswa", controller.GetNilaiMahasiswa)
	page.Get("/presensi", controller.GetAllPresensi)    //menampilkan seluruh data presensi
	page.Get("/presensi/:id", controller.GetPresensiID) //menampilkan data presensi berdasarkan id
	page.Post("/insrt-kemahaiswaan", controller.InsertKemahasiswaan)
	page.Post("/insrt-mahasiswa", controller.InsertDataMahasiswa)
	page.Post("/insrt-keuangan", controller.InsertKeuanganMahasiswa)
	page.Post("/insrt-nilai", controller.InsertNilaiMahasiswa)

	page.Post("/ins", controller.InsertData)
	page.Put("/upd/:id", controller.UpdateData)
	page.Delete("/delete/:id", controller.DeletePresensiByID)
}
