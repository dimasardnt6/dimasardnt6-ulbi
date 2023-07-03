package controller

import (
	"errors"
	"fmt"
	"time"

	cek "github.com/aiteung/presensi"
	"github.com/dimasardnt6/dimasardnt6-ulbi/config"
	"github.com/dimasardnt6/kemahasiswaan/model"
	inimodule "github.com/dimasardnt6/kemahasiswaan/module"
	"github.com/gofiber/fiber/v2"

	inimodellatihan "github.com/indrariksa/be_presensi/model"
	inimodullatihan "github.com/indrariksa/be_presensi/module"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"net/http"

	// Antrian Puskesmas

	modelantrian "github.com/dimasardnt6/BE-Antrian-Puskesmas/model"
	moduleantrian "github.com/dimasardnt6/BE-Antrian-Puskesmas/module"
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
	kemahasiswaan := inimodule.GetKemahasiswaanFromNpm("1214054", config.Ulbimongoconn, "kemahasiswaan")
	return c.JSON(kemahasiswaan)
}
func GetDataMahasiswa(c *fiber.Ctx) error {
	mahasiswa := inimodule.GetDataMahasiswaFromNpm("1214054", config.Ulbimongoconn, "data_mahasiswa")
	return c.JSON(mahasiswa)
}
func GetKeuanganMahasiswa(c *fiber.Ctx) error {
	keuangan := inimodule.GetKeuanganMahasiswaFromNomorHp("6289647129890", config.Ulbimongoconn, "keuangan_mahasiswa")
	return c.JSON(keuangan)
}
func GetNilaiMahasiswa(c *fiber.Ctx) error {
	nilai := inimodule.GetNilaiMahasiswaFromNama("Dimas Ardianto", config.Ulbimongoconn, "nilai_mahasiswa")
	return c.JSON(nilai)
}

//InsertFunction

func InsertKemahasiswaan(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var kemahasiswaan model.Kemahasiswaan
	if err := c.BodyParser(&kemahasiswaan); err != nil {
		return err
	}
	insertedID := inimodule.InsertKemahasiswaan(db, "kemahasiswaan",
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
	insertedID := inimodule.InsertDataMahasiswa(db, "data_mahasiswa",
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
	insertedID := inimodule.InsertKeuanganMahasiswa(db, "keuangan_mahasiswa",
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
	insertedID := inimodule.InsertNilaiMahasiswa(db, "nilai_mahasiswa",
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

//GetAllFunction

func GetAllKemahasiswaan(c *fiber.Ctx) error {
	ps := inimodule.GetAllKemahasiswaan(config.Ulbimongoconn, "kemahasiswaan")
	return c.JSON(ps)
}
func GetAllDataMahasiswa(c *fiber.Ctx) error {
	ps := inimodule.GetAllDataMahasiswa(config.Ulbimongoconn, "data_mahasiswa")
	return c.JSON(ps)
}
func GetAllKeuanganMahasiswa(c *fiber.Ctx) error {
	ps := inimodule.GetAllKeuanganMahasiswa(config.Ulbimongoconn, "keuangan_mahasiswa")
	return c.JSON(ps)
}
func GetAllNilaiMahasiswa(c *fiber.Ctx) error {
	ps := inimodule.GetAllNilaiMahasiswa(config.Ulbimongoconn, "nilai_mahasiswa")
	return c.JSON(ps)
}

// Latihan W-6

// GetAllPresensi godoc
// @Summary Get All Data Presensi.
// @Description Mengambil semua data presensi.
// @Tags Get All
// @Accept json
// @Produce json
// @Success 200 {object} Presensi
// @Router /presensi [get]
func GetAllPresensi(c *fiber.Ctx) error {
	ps := inimodullatihan.GetAllPresensi(config.Ulbimongoconn2, "presensi")
	return c.JSON(ps)
}

// GetPresensiID godoc
// @Summary Get By ID Data Presensi.
// @Description Ambil per ID data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /presensi/{id} [get]
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
	ps, err := inimodullatihan.GetPresensiFromID(objID, config.Ulbimongoconn2, "presensi")
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

func GetDataKemahasiswaanFromID(c *fiber.Ctx) error {
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
	ps, err := inimodule.GetDataKemahasiswaanFromID(objID, config.Ulbimongoconn, "kemahasiswaan")
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

func GetKemahasiswaanNPM(c *fiber.Ctx) error {
	queryValue := c.Query("npm")
	// i, err := strconv.Atoi(queryValue)
	// if err != nil {
	// 	// ... handle error
	// 	// panic(err)
	// 	return c.SendString("format params kurang tepat")
	// }
	ps := inimodule.GetKemahasiswaanFromNpm(queryValue, config.Ulbimongoconn, "kemahasiswaan")

	return c.JSON(ps)
}

// Week 9

// InsertData godoc
// @Summary Insert data presensi.
// @Description Input data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param request body Presensi true "Payload Body [RAW]"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 500
// @Router /ins [post]
func InsertData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn2
	var presensi inimodellatihan.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := inimodullatihan.InsertPresensi(db, "presensi",
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertDataKemahasiswaan(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var kemahasiswaan model.Kemahasiswaan
	if err := c.BodyParser(&kemahasiswaan); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := inimodule.InsertDataKemahasiswaan(db, "kemahasiswaan",
		kemahasiswaan.Identitas_Mahasiswa,
		kemahasiswaan.Status_Keuangan,
		kemahasiswaan.Nilai_Mahasiswa)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// Week 10

// Update Data

// UpdateData godoc
// @Summary Update data presensi.
// @Description Ubah data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body Presensi true "Payload Body [RAW]"
// @Success 200 {object} Presensi
// @Failure 400
// @Failure 500
// @Router /upd/{id} [put]
func UpdateData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn2

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var presensi inimodellatihan.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = inimodullatihan.UpdatePresensi(db, "presensi",
		objectID,
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

func UpdateDataKemahasiswaan(c *fiber.Ctx) error {
	db := config.Ulbimongoconn

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var kemahasiswaan model.Kemahasiswaan
	if err := c.BodyParser(&kemahasiswaan); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = inimodule.UpdateKemahasiswaan(db, "kemahasiswaan",
		objectID,
		kemahasiswaan.Identitas_Mahasiswa,
		kemahasiswaan.Status_Keuangan,
		kemahasiswaan.Nilai_Mahasiswa)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

// Delete Data

// DeletePresensiByID godoc
// @Summary Delete data presensi.
// @Description Hapus data presensi.
// @Tags Presensi
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete/{id} [delete]
func DeletePresensiByID(c *fiber.Ctx) error {
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

	err = inimodullatihan.DeletePresensiByID(objID, config.Ulbimongoconn2, "presensi")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}

func DeleteKemahasiswaanByID(c *fiber.Ctx) error {
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

	err = inimodule.DeleteKemahasiwaanByID(objID, config.Ulbimongoconn, "kemahasiswaan")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}

// Antrian Puskesmas

// Get All Function

// GetAllUser godoc
// @Summary Get All Data User.
// @Description Mengambil semua data user.
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} User
// @Router /user [get]
func GetAllUser(c *fiber.Ctx) error {
	ps := moduleantrian.GetAllUser(config.Ulbimongoconn3, "data_user")
	return c.JSON(ps)
}

// GetAllPasien godoc
// @Summary Get All Data Pasien.
// @Description Mengambil semua data pasien.
// @Tags Pasien
// @Accept json
// @Produce json
// @Success 200 {object} Pasien
// @Router /pasien [get]
func GetAllPasien(c *fiber.Ctx) error {
	ps := moduleantrian.GetAllPasien(config.Ulbimongoconn3, "data_pasien")
	return c.JSON(ps)
}

// GetAllAntrian godoc
// @Summary Get All Data Antrian.
// @Description Mengambil semua data antrian.
// @Tags Antrian
// @Accept json
// @Produce json
// @Success 200 {object} Antrian
// @Router /antrian [get]
func GetAllAntrian(c *fiber.Ctx) error {
	ps := moduleantrian.GetAllAntrian(config.Ulbimongoconn3, "data_antrian")
	return c.JSON(ps)
}

// GetAllPoliklinik godoc
// @Summary Get All Data Poliklinik.
// @Description Mengambil semua data poliklinik.
// @Tags Poliklinik
// @Accept json
// @Produce json
// @Success 200 {object} Poliklinik
// @Router /poliklinik [get]
func GetAllPoliklinik(c *fiber.Ctx) error {
	ps := moduleantrian.GetAllPoliklinik(config.Ulbimongoconn3, "data_poliklinik")
	return c.JSON(ps)
}

// GetAllDokter godoc
// @Summary Get All Data Dokter.
// @Description Mengambil semua data dokter.
// @Tags Dokter
// @Accept json
// @Produce json
// @Success 200 {object} Dokter
// @Router /dokter [get]
func GetAllDokter(c *fiber.Ctx) error {
	ps := moduleantrian.GetAllDokter(config.Ulbimongoconn3, "data_dokter")
	return c.JSON(ps)
}

// Get From ID

// GetUserFromID godoc
// @Summary Get By ID Data User.
// @Description Ambil per ID data user.
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} User
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /user/{id} [get]
func GetUserFromID(c *fiber.Ctx) error {
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
	ps, err := moduleantrian.GetUserFromID(objID, config.Ulbimongoconn3, "data_user")
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

// GetPasienFromID godoc
// @Summary Get By ID Data Pasien.
// @Description Ambil per ID data pasien.
// @Tags Pasien
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Pasien
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /pasien/{id} [get]
func GetPasienFromID(c *fiber.Ctx) error {
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
	ps, err := moduleantrian.GetPasienFromID(objID, config.Ulbimongoconn3, "data_pasien")
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

// GetAntrianFromID godoc
// @Summary Get By ID Data Antrian.
// @Description Ambil per ID data antrian.
// @Tags Antrian
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Antrian
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /antrian/{id} [get]
func GetAntrianFromID(c *fiber.Ctx) error {
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
	ps, err := moduleantrian.GetAntrianFromID(objID, config.Ulbimongoconn3, "data_antrian")
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

// GetPoliklinikFromID godoc
// @Summary Get By ID Data Poliklinik.
// @Description Ambil per ID data poliklinik.
// @Tags Poliklinik
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Poliklinik
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /poliklinik/{id} [get]
func GetPoliklinikFromID(c *fiber.Ctx) error {
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
	ps, err := moduleantrian.GetPoliklinikFromID(objID, config.Ulbimongoconn3, "data_poliklinik")
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

// GetDokterFormID godoc
// @Summary Get By ID Data Dokter.
// @Description Ambil per ID data dokter.
// @Tags Dokter
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} Dokter
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /dokter/{id} [get]
func GetDokterFromID(c *fiber.Ctx) error {
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
	ps, err := moduleantrian.GetDokterFromID(objID, config.Ulbimongoconn3, "data_dokter")
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

// Login Signup

func SignUp(c *fiber.Ctx) error {
	db := config.Ulbimongoconn3
	var data modelantrian.User
	if err := c.BodyParser(&data); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	_, err := moduleantrian.SignUp(db, "data_user", data)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Akun berhasil disimpan.",
	})
}

// Define a secret key for signing the JWT token
var jwtSecret = []byte("secret-key")

func SignIn(c *fiber.Ctx) error {
	db := config.Ulbimongoconn3
	var data modelantrian.User
	if err := c.BodyParser(&data); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	email, err := moduleantrian.LogIn(db, "data_user", data)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = data.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Set token expiration time to 24 hours

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Selamat datang " + data.Fullname,
		"email":   email,
		"token":   tokenString,
	})
}

func AuthenticateMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization") // Assuming the token is sent in the Authorization header

	// Verify the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"status":  http.StatusUnauthorized,
			"message": "Invalid or expired token",
		})
	}

	// Token is valid, proceed with the request
	return c.Next()
}

// func SignIn(c *fiber.Ctx) error {
// 	db := config.Ulbimongoconn3
// 	var data modelantrian.User
// 	if err := c.BodyParser(&data); err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"status":  http.StatusInternalServerError,
// 			"message": err.Error(),
// 		})
// 	}
// 	user, err := moduleantrian.LogIn(db, "data_user", data)
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"status":  http.StatusInternalServerError,
// 			"message": err.Error(),
// 		})
// 	}
// 	return c.Status(http.StatusOK).JSON(fiber.Map{
// 		"status":  http.StatusOK,
// 		"message": "Selamat datang " + user,
// 	})
// }

// Insert Function

// InsertUser godoc
// @Summary Insert data user.
// @Description Input data user.
// @Tags User
// @Accept json
// @Produce json
// @Param request body User true "Payload Body [RAW]"
// @Success 200 {object} User
// @Failure 400
// @Failure 500
// @Router /ins [post]
func InsertUser(c *fiber.Ctx) error {
	db := config.Ulbimongoconn3
	var data modelantrian.User
	if err := c.BodyParser(&data); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	if data.Fullname == "" || data.Email == "" || data.Password == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Data tidak boleh kosong",
		})
	}
	insertedID, err := moduleantrian.InsertUser(db, "data_user", data)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// InsertPasien godoc
// @Summary Insert data pasieni.
// @Description Input data pasien.
// @Tags Pasien
// @Accept json
// @Produce json
// @Param request body Pasien true "Payload Body [RAW]"
// @Success 200 {object} Pasien
// @Failure 400
// @Failure 500
// @Router /ins [post]
func InsertPasien(c *fiber.Ctx) error {
	db := config.Ulbimongoconn3
	var pasien modelantrian.Pasien
	if err := c.BodyParser(&pasien); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := moduleantrian.InsertPasien(db, "data_pasien",
		pasien.Nama_Pasien,
		pasien.Nomor_Ktp,
		pasien.Alamat,
		pasien.Nomor_Telepon,
		pasien.Tanggal_Lahir,
		pasien.Jenis_Kelamin)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// InsertAntrian godoc
// @Summary Insert data antrian.
// @Description Input data antrian.
// @Tags Antrian
// @Accept json
// @Produce json
// @Param request body Antrian true "Payload Body [RAW]"
// @Success 200 {object} Antrian
// @Failure 400
// @Failure 500
// @Router /ins [post]
func InsertAntrian(c *fiber.Ctx) error {
	db := config.Ulbimongoconn3
	var antrian modelantrian.Antrian
	if err := c.BodyParser(&antrian); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := moduleantrian.InsertAntrian(db, "data_antrian",
		antrian.Poli,
		antrian.Identitas_Pasien,
		antrian.Status_Antrian)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// InsertPoliklinik godoc
// @Summary Insert data poliklinik.
// @Description Input data poliklinik.
// @Tags Poliklinik
// @Accept json
// @Produce json
// @Param request body Poliklinik true "Payload Body [RAW]"
// @Success 200 {object} Poliklinik
// @Failure 400
// @Failure 500
// @Router /ins [post]
func InsertPoliklinik(c *fiber.Ctx) error {
	db := config.Ulbimongoconn3
	var poliklinik modelantrian.Poliklinik
	if err := c.BodyParser(&poliklinik); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := moduleantrian.InsertPoliklinik(db, "data_poliklinik",
		poliklinik.Kode_Poliklinik,
		poliklinik.Nama_Poliklinik,
		poliklinik.Deskripsi,
		poliklinik.Identitas_Dokter)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// InsertDokter godoc
// @Summary Insert data dokteri.
// @Description Input data dokter.
// @Tags dokter
// @Accept json
// @Produce json
// @Param request body dokter true "Payload Body [RAW]"
// @Success 200 {object} dokter
// @Failure 400
// @Failure 500
// @Router /ins [post]
func InsertDokter(c *fiber.Ctx) error {
	db := config.Ulbimongoconn3
	var dokter modelantrian.Dokter
	if err := c.BodyParser(&dokter); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := moduleantrian.InsertDokter(db, "data_dokter",
		dokter.Nama_Dokter,
		dokter.Spesialisasi)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// Update Function

// UpdatePasien godoc
// @Summary Update data pasien.
// @Description Ubah data pasien.
// @Tags Pasien
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body Pasien true "Payload Body [RAW]"
// @Success 200 {object} Pasien
// @Failure 400
// @Failure 500
// @Router /upd/{id} [put]
func UpdatePasien(c *fiber.Ctx) error {
	db := config.Ulbimongoconn3

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var pasien modelantrian.Pasien
	if err := c.BodyParser(&pasien); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = moduleantrian.UpdatePasien(db, "data_pasien",
		objectID,
		pasien.Nama_Pasien,
		pasien.Nomor_Ktp,
		pasien.Alamat,
		pasien.Nomor_Telepon,
		pasien.Tanggal_Lahir,
		pasien.Jenis_Kelamin)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

// UpdateAntrian godoc
// @Summary Update data antrian.
// @Description Ubah data antrian.
// @Tags Antrian
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body Antrian true "Payload Body [RAW]"
// @Success 200 {object} Antrian
// @Failure 400
// @Failure 500
// @Router /upd/{id} [put]
func UpdateAntrian(c *fiber.Ctx) error {
	db := config.Ulbimongoconn3

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var antrian modelantrian.Antrian
	if err := c.BodyParser(&antrian); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = moduleantrian.UpdateAntrian(db, "data_antrian",
		objectID,
		antrian.Poli,
		antrian.Identitas_Pasien,
		antrian.Nomor_Antrian,
		antrian.Status_Antrian)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

// UpdatePoliklinik godoc
// @Summary Update data polklinik.
// @Description Ubah data poliklinik.
// @Tags Poliklinik
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body Poliklink true "Payload Body [RAW]"
// @Success 200 {object} Poliklinik
// @Failure 400
// @Failure 500
// @Router /upd/{id} [put]
func UpdatePoliklinik(c *fiber.Ctx) error {
	db := config.Ulbimongoconn3

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var poliklinik modelantrian.Poliklinik
	if err := c.BodyParser(&poliklinik); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = moduleantrian.UpdatePoliklinik(db, "data_poliklinik",
		objectID,
		poliklinik.Kode_Poliklinik,
		poliklinik.Nama_Poliklinik,
		poliklinik.Deskripsi,
		poliklinik.Identitas_Dokter)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

// UpdateDokter godoc
// @Summary Update data dokter.
// @Description Ubah data dokter.
// @Tags Dokter
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body Dokter true "Payload Body [RAW]"
// @Success 200 {object} Dokter
// @Failure 400
// @Failure 500
// @Router /upd/{id} [put]
func UpdateDokter(c *fiber.Ctx) error {
	db := config.Ulbimongoconn3

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var dokter modelantrian.Dokter
	if err := c.BodyParser(&dokter); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = moduleantrian.UpdateDokter(db, "data_dokter",
		objectID,
		dokter.Nama_Dokter,
		dokter.Spesialisasi)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

// Delete Function

// DeletePasienByID godoc
// @Summary Delete data pasien.
// @Description Hapus data pasien.
// @Tags Pasien
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete/{id} [delete]
func DeletePasienByID(c *fiber.Ctx) error {
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

	err = moduleantrian.DeletePasienByID(objID, config.Ulbimongoconn3, "data_pasien")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}

// DeleteAntrianByID godoc
// @Summary Delete data antrian.
// @Description Hapus data antrian.
// @Tags Antrian
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete/{id} [delete]
func DeleteAntrianByID(c *fiber.Ctx) error {
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

	err = moduleantrian.DeleteAntrianByID(objID, config.Ulbimongoconn3, "data_antrian")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}

// DeletePoliklinikByID godoc
// @Summary Delete data poliklinik.
// @Description Hapus data poliklinik.
// @Tags Poliklinik
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete/{id} [delete]
func DeletePoliklinikByID(c *fiber.Ctx) error {
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

	err = moduleantrian.DeletePoliklinikByID(objID, config.Ulbimongoconn3, "data_poliklinik")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}

// DeleteDokterByID godoc
// @Summary Delete data dokter.
// @Description Hapus data dokter.
// @Tags Dokter
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete/{id} [delete]
func DeleteDokterByID(c *fiber.Ctx) error {
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

	err = moduleantrian.DeleteDokterByID(objID, config.Ulbimongoconn3, "data_dokter")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}

// DeleteUserByID godoc
// @Summary Delete data user.
// @Description Hapus data user.
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /delete/{id} [delete]
func DeleteUserByID(c *fiber.Ctx) error {
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

	err = moduleantrian.DeleteDokterByID(objID, config.Ulbimongoconn3, "data_user")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}
