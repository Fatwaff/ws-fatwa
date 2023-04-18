package controller

import (
	"errors"
	"fmt"
	"net/http"

	inimodel "github.com/Fatwaff/presensi_mahasiswa/model"
	inimodule "github.com/Fatwaff/presensi_mahasiswa/module"
	"github.com/Fatwaff/ws-fatwa/config"
	"github.com/aiteung/musik"
	cek "github.com/aiteung/presensi"
	"github.com/gofiber/fiber/v2"
	inimodullatihan "github.com/indrariksa/be_presensi/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"github_repo": "https://github.com/Fatwaff/ws-fatwa",
		"message":     "You are at the root endpoint 😉",
		"success":     true,
	})
}

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetPresensi(c *fiber.Ctx) error {
     ps := cek.GetPresensiCurrentMonth(config.Ulbimongoconn)
     return c.JSON(ps)
}

func GetAllPresensi(c *fiber.Ctx) error {
	ps := inimodule.GetAllPresensiFromKehadiran("masuk", config.Ulbimongoconn, "presensi")
	return c.JSON(ps)
}

func GetAllPresensi2(c *fiber.Ctx) error {
	ps := inimodullatihan.GetAllPresensi(config.Ulbimongoconn2, "presensi")
	return c.JSON(ps)
}

func GetSemuaMahasiswa(c *fiber.Ctx) error {
	ps := inimodule.GetAllMahasiswa(config.Ulbimongoconn, "mahasiswa")
	return c.JSON(ps)
}
func GetSemuaKelas(c *fiber.Ctx) error {
	ps := inimodule.GetAllKelas(config.Ulbimongoconn, "kelas")
	return c.JSON(ps)
}
func GetSemuaProdi(c *fiber.Ctx) error {
	ps := inimodule.GetAllProdi(config.Ulbimongoconn, "prodi")
	return c.JSON(ps)
}
func GetSemuaMataKuliah(c *fiber.Ctx) error {
	ps := inimodule.GetAllMataKuliah(config.Ulbimongoconn, "matkul")
	return c.JSON(ps)
}
func GetSemuaDosen(c *fiber.Ctx) error {
	ps := inimodule.GetAllDosen(config.Ulbimongoconn, "dosen")
	return c.JSON(ps)
}
func GetSemuaRuangKuliah(c *fiber.Ctx) error {
	ps := inimodule.GetAllRuangKuliah(config.Ulbimongoconn, "ruang")
	return c.JSON(ps)
}
func GetSemuaPresensi(c *fiber.Ctx) error {
	ps := inimodule.GetAllPresensi(config.Ulbimongoconn, "presensi")
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

func GetMahasiswa(c *fiber.Ctx) error {
	ps := inimodule.GetMahasiswaFromNpm(1214039, config.Ulbimongoconn, "mahasiswa")
	return c.JSON(ps)
}
func GetKelas(c *fiber.Ctx) error {
	ps := inimodule.GetKelasFromKodeKelas("TI-B2", config.Ulbimongoconn, "kelas")
	return c.JSON(ps)
}
func GetMatkul(c *fiber.Ctx) error {
	ps := inimodule.GetMatkulFromKodeMatkul(21711, config.Ulbimongoconn, "matkul")
	return c.JSON(ps)
}

func InsertDataPresensi(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var presensi inimodel.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return err
	}
	insertedID := inimodule.InsertPresensi(db, "presensi",
		presensi.Kehadiran,
		presensi.Biodata,
		presensi.Mata_kuliah)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}
func InsertDataMahasiswa(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var mahasiswa inimodel.Mahasiswa
	if err := c.BodyParser(&mahasiswa); err != nil {
		return err
	}
	insertedID := inimodule.InsertMahasiswa(db, "mahasiswa",
		mahasiswa.Nama,
		mahasiswa.Npm,
		mahasiswa.Nama_kelas,
		mahasiswa.Jurusan)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}
func InsertDataKelas(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var kelas inimodel.Kelas
	if err := c.BodyParser(&kelas); err != nil {
		return err
	}
	insertedID := inimodule.InsertKelas(db, "kelas",
		kelas.Kode_kelas,
		kelas.Nama_kelas,
		kelas.Kapasitas)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}
func InsertDataMatkul(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var matkul inimodel.MataKuliah
	if err := c.BodyParser(&matkul); err != nil {
		return err
	}
	insertedID := inimodule.InsertMatkul(db, "matkul",
		matkul.Kode_matkul,
		matkul.Nama_matkul,
		matkul.Sks,
		matkul.Dosen_pengajar,
		matkul.Jadwal_kuliah,
		matkul.Ruang_kuliah)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}