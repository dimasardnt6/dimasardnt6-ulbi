package url

import (
	"github.com/dimasardnt6/dimasardnt6-ulbi/controller"

	"github.com/gofiber/swagger" // swagger handler

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
	page.Get("/kemahasiswaan/:id", controller.GetDataKemahasiswaanFromID) //menampilkan data kemahasiwaan berdasarkan id

	page.Get("/presensi", controller.GetAllPresensi)    //menampilkan seluruh data presensi
	page.Get("/presensi/:id", controller.GetPresensiID) //menampilkan data presensi berdasarkan id

	page.Post("/insrt-kemahasiswaan", controller.InsertKemahasiswaan)
	page.Post("/insrt-mahasiswa", controller.InsertDataMahasiswa)
	page.Post("/insrt-keuangan", controller.InsertKeuanganMahasiswa)
	page.Post("/insrt-nilai", controller.InsertNilaiMahasiswa)
	page.Post("/ins-kemahasiswaan", controller.InsertDataKemahasiswaan)

	page.Put("/upd-kemahasiswaan/:id", controller.UpdateDataKemahasiswaan)
	page.Delete("/delete-kemahasiswaan/:id", controller.DeleteKemahasiswaanByID)

	page.Post("/ins", controller.InsertData)
	page.Put("/upd/:id", controller.UpdateData)
	page.Delete("/delete/:id", controller.DeletePresensiByID)

	page.Get("/docs/*", swagger.HandlerDefault)

	// Endpoint Antrian Puskesmas
	// Get Endpoint
	page.Get("/all-user", controller.GetAllUser)
	page.Get("/all-pasien", controller.GetAllPasien)
	page.Get("/all-antrian", controller.GetAllAntrian)
	page.Get("/all-poliklinik", controller.GetAllPoliklinik)
	page.Get("/all-dokter", controller.GetAllDokter)
	page.Get("/user/:id", controller.GetUserFromID)             //menampilkan data pasien berdasarkan id
	page.Get("/pasien/:id", controller.GetPasienFromID)         //menampilkan data pasien berdasarkan id
	page.Get("/antrian/:id", controller.GetAntrianFromID)       //menampilkan data antrian berdasarkan id
	page.Get("/poliklinik/:id", controller.GetPoliklinikFromID) //menampilkan data poliklinik berdasarkan id
	page.Get("/dokter/:id", controller.GetDokterFromID)         //menampilkan data dokter berdasarkan id

	// Insert Endpoint
	page.Post("/ins-pasien", controller.InsertPasien)
	page.Post("/ins-antrian", controller.InsertAntrian)
	page.Post("/ins-poliklinik", controller.InsertPoliklinik)
	page.Post("/ins-dokter", controller.InsertDokter)

	// Update Endpoint
	page.Put("/upd-pasien/:id", controller.UpdatePasien)
	page.Put("/upd-antrian/:id", controller.UpdateAntrian)
	page.Put("/upd-poliklinik/:id", controller.UpdatePoliklinik)
	page.Put("/upd-dokter/:id", controller.UpdateDokter)

	// Delete Endpoint
	page.Delete("/delete-pasien/:id", controller.DeletePasienByID)
	page.Delete("/delete-antrian/:id", controller.DeleteAntrianByID)
	page.Delete("/delete-poliklinik/:id", controller.DeletePoliklinikByID)
	page.Delete("/delete-dokter/:id", controller.DeleteDokterByID)

	// Signup-Login Endpoint
	page.Post("/sign-up", controller.SignUp)
	page.Post("/sign-in", controller.SignIn)
}
