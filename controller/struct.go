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

//struct antrian puskesmas 
type User struct {
	// ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Fullname        string             `bson:"fullname,omitempty" json:"fullname,omitempty" example:"Budiono"`
	Email           string             `bson:"email,omitempty" json:"email,omitempty" example:"budiono@gmail.com"`
	Password        string             `bson:"password,omitempty" json:"password,omitempty" example:"bdn68"`
	Confirmpassword string             `bson:"confirmpass,omitempty" json:"confirmpass,omitempty" example:"bdn68"`
}

type Token struct {
	Token_String string `bson:"tokenstring,omitempty" json:"tokenstring,omitempty"`
}

type Pasien struct {
	// ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	Nama_Pasien   string             `bson:"nama_pasien,omitempty" json:"nama_pasien,omitempty" example:"Budiono"`
	Nomor_Ktp     string             `bson:"nomor_ktp,omitempty" json:"nomor_ktp,omitempty" example:"3217060601020998"`
	Alamat        string             `bson:"alamat,omitempty" json:"alamat,omitempty" example:"Cimahi"`
	Nomor_Telepon string             `bson:"nomor_telepon,omitempty" json:"nomor_telepon,omitempty" example:"089647129899"`
	Tanggal_Lahir string             `bson:"tanggal_lahir,omitempty" json:"tanggal_lahir,omitempty" example:"18 Januari 2002"`
	Jenis_Kelamin string             `bson:"jenis_kelamin,omitempty" json:"jenis_kelamin,omitempty" example:"Laki-Laki"`
}

type Antrian struct {
	// ID                  primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	Poli                Poliklinik         `bson:"poli,omitempty" json:"poli,omitempty"`
	Identitas_Pasien    Pasien             `bson:"identitas_pasien,omitempty" json:"identitas_pasien,omitempty"`
	Nomor_Antrian       int                `bson:"nomor_antrian,omitempty" json:"nomor_antrian,omitempty"`
	// Tanggal_Pendaftaran primitive.DateTime `bson:"tanggal_pendaftaran,omitempty" json:"tanggal_pendaftaran,omitempty"`
	Status_Antrian      string             `bson:"status_antrian,omitempty" json:"status_antrian,omitempty" example:"Menunggu"`
}

type Poliklinik struct {
	// ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	Kode_Poliklinik  string             `bson:"kode_poliklinik,omitempty" json:"kode_poliklinik,omitempty" example:"PLUM"`
	Nama_Poliklinik  string             `bson:"nama_poliklinik,omitempty" json:"nama_poliklinik,omitempty" example:"PoliklinikUmum"`
	Deskripsi        string             `bson:"deskripsi,omitempty" json:"deskripsi,omitempty" example:"Menyediakan Layanan Kesehatan Umum"`
	Identitas_Dokter Dokter             `bson:"dokter,omitempty" json:"dokter,omitempty"`
}

type Dokter struct {
	// ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	Nama_Dokter  string             `bson:"nama_dokter,omitempty" json:"nama_dokter,omitempty" example:"Dr.William"`
	Spesialisasi string             `bson:"spesialisasi,omitempty" json:"spesialisasi,omitempty" example:"Dokter Umum"`
}