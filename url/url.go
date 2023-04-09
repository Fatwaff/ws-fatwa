package url

import (
	"github.com/Fatwaff/ws-fatwa/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Home) 
	page.Get("/semua-presensi", controller.GetSemuaPresensi)
	page.Get("/semua-mahasiswa", controller.GetSemuaMahasiswa)
	page.Get("/semua-kelas", controller.GetSemuaKelas)
	page.Get("/semua-prodi", controller.GetSemuaProdi)
	page.Get("/semua-matkul", controller.GetSemuaMataKuliah)
	page.Get("/semua-dosen", controller.GetSemuaDosen)
	page.Get("/semua-ruangan", controller.GetSemuaRuangKuliah)
	page.Get("/presensi", controller.GetAllPresensi)
	page.Get("/mahasiswa", controller.GetMahasiswa)
	page.Get("/kelas", controller.GetKelas)
	page.Get("/matkul", controller.GetMatkul)
	page.Post("/ins_presensi", controller.InsertDataPresensi)
	page.Post("/ins_mahasiswa", controller.InsertDataMahasiswa)
	page.Post("/ins_kelas", controller.InsertDataKelas)
	page.Post("/ins_matkul", controller.InsertDataMatkul)
	//
	page.Get("/presensi2", controller.GetAllPresensi2) //menampilkan seluruh data presensi
	page.Get("/presensi2/:id", controller.GetPresensiID) //menampilkan data presensi berdasarkan id
}
