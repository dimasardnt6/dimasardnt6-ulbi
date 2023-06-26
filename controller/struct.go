package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Kemahasiswaan struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Identitas_Mahasiswa Mahasiswa          `bson:"identitas,omitempty" json:"identitas,omitempty"`
	Status_Keuangan     Keuangan           `bson:"status_keuangan,omitempty" json:"status_keuangan,omitempty"`
	Nilai_Mahasiswa     Nilai              `bson:"nilai_mhs,omitempty" json:"nilai_mhs,omitempty"`
}

type Mahasiswa struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Npm             string             `bson:"npm,omitempty" json:"npm,omitempty"`
	Nama            string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Nomor_Handphone string             `bson:"no_hp,omitempty" json:"no_hp,omitempty"`
	Prodi           ProgramStudi       `bson:"prodi,omitempty" json:"prodi,omitempty"`
	Jurusan         string             `bson:"jurusan,omitempty" json:"jurusan,omitempty"`
	Kelas           string             `bson:"kelas,omitempty" json:"kelas,omitempty"`
}

type ProgramStudi struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Kode_Prodi string             `bson:"kode_prodi,omitempty" json:"kode_prodi,omitempty"`
	Nama_Prodi string             `bson:"nama,omitempty" json:"nama,omitempty"`
}

type Keuangan struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Biodata          Mahasiswa          `bson:"bio_mahasiswa,omitempty" json:"biodata,omitempty"`
	Total_Pembayaran int                `bson:"total_pembayaran,omitempty" json:"total_pembayaran,omitempty"`
}

type Nilai struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Biodata_Mahasiswa Mahasiswa          `bson:"bio_mhs,omitempty" json:"bio_mhs,omitempty"`
	Matakuliah        Matakuliah         `bson:"matakuliah,omitempty" json:"matakuliah,omitempty"`
	Nilai_Angka       int                `bson:"nilai_angka,omitempty" json:"nilai_angka,omitempty"`
	Nilai_Huruf       string             `bson:"nilai_huruf,omitempty" json:"nilai_huruf,omitempty"`
}

type Matakuliah struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_Matkul string             `bson:"nama_matkul,omitempty" json:"nama_matkul,omitempty"`
	Nama_Dosen  string             `bson:"dosen,omitempty" json:"dosen,omitempty"`
}

// struct presensi
type Karyawan struct {
	Nama        string     `bson:"nama,omitempty" json:"nama,omitempty" example:"Tes Swagger"`
	PhoneNumber string     `bson:"phone_number,omitempty" json:"phone_number,omitempty" example:"08123456789"`
	Jabatan     string     `bson:"jabatan,omitempty" json:"jabatan,omitempty" example:"Anonymous"`
	Jam_kerja   []JamKerja `bson:"jam_kerja,omitempty" json:"jam_kerja,omitempty"`
	Hari_kerja  []string   `bson:"hari_kerja,omitempty" json:"hari_kerja,omitempty" example:"Senin,Selasa,Rabu,Kamis,Jumat,Sabtu,Minggu"`
}

type JamKerja struct {
	Durasi     int      `bson:"durasi,omitempty" json:"durasi,omitempty" example:"8"`
	Jam_masuk  string   `bson:"jam_masuk,omitempty" json:"jam_masuk,omitempty" example:"08:00"`
	Jam_keluar string   `bson:"jam_keluar,omitempty" json:"jam_keluar,omitempty" example:"16:00"`
	Gmt        int      `bson:"gmt,omitempty" json:"gmt,omitempty" example:"7"`
	Hari       []string `bson:"hari,omitempty" json:"hari,omitempty" example:"Senin,Selasa,Rabu,Kamis,Jumat,Sabtu,Minggu"`
	Shift      int      `bson:"shift,omitempty" json:"shift,omitempty" example:"2"`
	Piket_tim  string   `bson:"piket_tim,omitempty" json:"piket_tim,omitempty" example:"Piket Z"`
}

type Presensi struct {
	Longitude    float64 `bson:"longitude,omitempty" json:"longitude,omitempty" example:"123.11"`
	Latitude     float64 `bson:"latitude,omitempty" json:"latitude,omitempty" example:"123.11"`
	Location     string  `bson:"location,omitempty" json:"location,omitempty" example:"Bandung"`
	Phone_number string  `bson:"phone_number,omitempty" json:"phone_number,omitempty" example:"08123456789"`
	//Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Checkin string   `bson:"checkin,omitempty" json:"checkin,omitempty" example:"MASUK"`
	Biodata Karyawan `bson:"biodata,omitempty" json:"biodata,omitempty"`
}
