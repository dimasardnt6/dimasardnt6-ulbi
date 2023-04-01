package controller

import (
	"errors"
	"fmt"

	cek "github.com/aiteung/presensi"
	"github.com/dimasardnt6/dimasardnt6-ulbi/config"
	"github.com/dimasardnt6/kemahasiswaan/model"
	"github.com/dimasardnt6/kemahasiswaan/module"
	"github.com/gofiber/fiber/v2"

	inimodullatihan "github.com/indrariksa/be_presensi/module"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"net/http"
)

func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"github_repo": "https://github.com/dimasardnt6/dimasardnt6-ulbi",
		"message":     "You are at the root endpoint 😉",
		"success":     true,
	})
}

//GetFunction

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

//InsertFunction

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

//GetFunction

func GetAllKemahasiswaan(c *fiber.Ctx) error {
	ps := module.GetAllKemahasiswaan(config.Ulbimongoconn, "kemahasiswaan")
	return c.JSON(ps)
}
func GetAllDataMahasiswa(c *fiber.Ctx) error {
	ps := module.GetAllDataMahasiswa(config.Ulbimongoconn, "data_mahasiswa")
	return c.JSON(ps)
}
func GetAllKeuanganMahasiswa(c *fiber.Ctx) error {
	ps := module.GetAllKeuanganMahasiswa(config.Ulbimongoconn, "keuangan_mahasiswa")
	return c.JSON(ps)
}
func GetAllNilaiMahasiswa(c *fiber.Ctx) error {
	ps := module.GetAllNilaiMahasiswa(config.Ulbimongoconn, "nilai_mahasiswa")
	return c.JSON(ps)
}

//Latihan W-6

func GetAllPresensi(c *fiber.Ctx) error {
	ps := inimodullatihan.GetAllPresensi(config.Ulbimongoconn, "presensi")
	return c.JSON(ps)
}
func GetPresensiID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := inimodullatihan.GetPresensiFromID(objID, config.Ulbimongoconn, "presensi")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}
