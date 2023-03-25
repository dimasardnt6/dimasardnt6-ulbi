package controller

import (
	"github.com/aiteung/musik"
	cek "github.com/aiteung/presensi"
	"github.com/dimasardnt6/dimasardnt6-ulbi/config"
	"github.com/gofiber/fiber/v2"

	"net/http"

	"github.com/dimasardnt6/kemahasiswaan/model"
	"github.com/dimasardnt6/kemahasiswaan/module"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetPresensi(c *fiber.Ctx) error {
	ps := cek.GetPresensiCurrentMonth(config.Ulbimongoconn)
	return c.JSON(ps)
}
func GetKemahasiswaan(c *fiber.Ctx) error {
	kemahasiswaan := module.GetKemahasiswaanFromNpm("1214054", config.Ulbimongoconn, "kemahasiswaan")
	return c.JSON(kemahasiswaan)
}
func GetDataMahasiswa(c *fiber.Ctx) error {
	mahasiswa := module.GetDataMahasiswaFromNpm("1214054", config.Ulbimongoconn, "data_mahasiswa")
	return c.JSON(mahasiswa)
}
func GetKeuanganMahasiswa(c *fiber.Ctx) error {
	keuangan := module.GetKeuanganMahasiswaFromNomorHp("6289647129890", config.Ulbimongoconn, "keuangan_mahasiswa")
	return c.JSON(keuangan)
}
func GetNilaiMahasiswa(c *fiber.Ctx) error {
	nilai := module.GetNilaiMahasiswaFromNama("Dimas Ardianto", config.Ulbimongoconn, "nilai_mahasiswa")
	return c.JSON(nilai)
}

func InsertKemahasiswaan(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var kemahasiswaan model.Kemahasiswaan
	if err := c.BodyParser(&kemahasiswaan); err != nil {
		return err
	}
	insertedID := module.InsertKemahasiswaan(db, "kemahasiswaan",
		kemahasiswaan.Identitas_Mahasiswa,
		kemahasiswaan.Status_Keuangan,
		kemahasiswaan.Nilai_Mahasiswa)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertDataMahasiswa(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var mahasiswa model.Mahasiswa
	if err := c.BodyParser(&mahasiswa); err != nil {
		return err
	}
	insertedID := module.InsertDataMahasiswa(db, "data_mahasiswa",
		mahasiswa.Npm,
		mahasiswa.Nama,
		mahasiswa.Nomor_Handphone,
		mahasiswa.Prodi,
		mahasiswa.Jurusan,
		mahasiswa.Kelas)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertKeuanganMahasiswa(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var keuangan model.Keuangan
	if err := c.BodyParser(&keuangan); err != nil {
		return err
	}
	insertedID := module.InsertKeuanganMahasiswa(db, "keuangan_mahasiswa",
		keuangan.Biodata,
		keuangan.Total_Pembayaran)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertNilaiMahasiswa(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var nilai model.Nilai
	if err := c.BodyParser(&nilai); err != nil {
		return err
	}
	insertedID := module.InsertNilaiMahasiswa(db, "nilai_mahasiswa",
		nilai.Biodata_Mahasiswa,
		nilai.Matakuliah,
		nilai.Nilai_Angka,
		nilai.Nilai_Huruf)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}
