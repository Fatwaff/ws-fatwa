package url

import (
	"github.com/Fatwaff/ws-fatwa/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
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
	page.Get("/semua-test", controller.GetSemuaDataRuangan)
	page.Get("/presensi", controller.GetAllPresensi)
	page.Get("/mahasiswa", controller.GetMahasiswa)
	page.Get("/kelas", controller.GetKelas)
	page.Get("/matkul", controller.GetMatkul)
	page.Post("/presensi", controller.InsertDataPresensi)
	page.Post("/mahasiswa", controller.InsertDataMahasiswa)
	page.Post("/kelas", controller.InsertDataKelas)
	page.Post("/matkul", controller.InsertDataMatkul)
	// tb ws
	page.Get("/user", controller.GetAllUser)
	page.Post("/user", controller.InsertUser)
	page.Post("/signup", controller.SignUp)
	page.Post("/login", controller.LogIn)
	page.Post("/auth", controller.Authenticated)
	// tb ltmn
	page.Get("/kamtibmas", controller.GetAllKamtibmas)
	//
	page.Get("/presensi2", controller.GetAllPresensi2) //menampilkan seluruh data presensi
	page.Get("/presensi2/:id", controller.GetPresensiID) //menampilkan data presensi berdasarkan id
	page.Post("/presensi2", controller.InsertDataPresensi2) //post
	page.Put("/presensi2/:id", controller.UpdateDataPresensi) //update
	page.Delete("/presensi2/:id", controller.DeletePresensiByID)//delete
	page.Get("/docs/*", swagger.HandlerDefault)
}
