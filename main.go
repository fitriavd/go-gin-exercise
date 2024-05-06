package main

import (
	"encoding/json"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

type (
	Siswa struct {
		NIS    int    `json:"nis"`
		Nama   string `json:"nama"`
		Kelas  int    `json:"kelas"`
		Gender string `json:"gender"`
	}

	Guru struct {
		NIG    int    `json:"nig"`
		Nama   string `json:"nama"`
		Gender string `json:"gender"`
	}

	MataPelajaran struct {
		ID    int    `json:"id"`
		Nama  string `json:"nama"`
		NIG   int    `json:"nig"`
		Kelas int    `json:"kelas"`
	}
)

var students = []Siswa{
	{
		NIS:    1,
		Nama:   "Azmi",
		Kelas:  4,
		Gender: "L",
	},
	{
		NIS:    2,
		Nama:   "Dhaby",
		Kelas:  5,
		Gender: "L",
	},
	{
		NIS:    3,
		Nama:   "Fitria",
		Kelas:  6,
		Gender: "P",
	},
}

var Teachers = []Guru{
	{
		NIG:    1,
		Nama:   "Guru Pertama",
		Gender: "L",
	},
	{
		NIG:    2,
		Nama:   "Guru Kedua",
		Gender: "P",
	},
	{
		NIG:    3,
		Nama:   "Guru Ketiga",
		Gender: "P",
	},
}

var Lessons = []MataPelajaran{
	{
		ID:    1,
		Nama:  "IPA",
		NIG:   1,
		Kelas: 4,
	},
	{
		ID:    2,
		Nama:  "IPA",
		NIG:   2,
		Kelas: 5,
	},
	{
		ID:    3,
		Nama:  "IPA",
		NIG:   3,
		Kelas: 6,
	},
	{
		ID:    4,
		Nama:  "IPS",
		NIG:   1,
		Kelas: 4,
	},
	{
		ID:    5,
		Nama:  "MTK",
		NIG:   2,
		Kelas: 5,
	},
}

func StudentGet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get all data siswa
		ctx.Writer.WriteHeader(200)
		dataSiswaJson, err := json.Marshal(students)
		if err != nil {
			panic(err)
		}
		_, err = ctx.Writer.Write(dataSiswaJson)
		if err != nil {
			panic(err)
		}
		return
	}
}

func StudentPost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get 1 data siswa by nis
		// ambil data siswa dari body
		defer ctx.Request.Body.Close()
		dataBody, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			panic(err)
		}
		nis := struct {
			NIS int `json:"nis"`
		}{}
		err = json.Unmarshal(dataBody, &nis)
		if err != nil {
			panic(err)
		}

		dataSiswa := Siswa{NIS: 0}
		for _, val := range students {
			if val.NIS == nis.NIS {
				dataSiswa = val
			}
		}
		if dataSiswa.NIS == 0 {
			ctx.Writer.WriteHeader(404)
			_, err := ctx.Writer.Write([]byte("Data not found"))
			if err != nil {
				panic(err)
			}
			return
		}
		dataSiswaJson, err := json.Marshal(dataSiswa)
		if err != nil {
			panic(err)
		}
		_, err = ctx.Writer.Write(dataSiswaJson)
		if err != nil {
			panic(err)
		}
		return
	}
}

func TeacherGet() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get all data guru
		ctx.Writer.WriteHeader(200)
		dataGuruJson, err := json.Marshal(Teachers)
		if err != nil {
			panic(err)
		}
		_, err = ctx.Writer.Write(dataGuruJson)
		if err != nil {
			panic(err)
		}
		return
	}
}

func TeacherPost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get 1 data guru by nig
		// ambil data guru dari body
		defer ctx.Request.Body.Close()
		dataBody, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			panic(err)
		}
		nig := struct {
			NIG int `json:"nig"`
		}{}
		err = json.Unmarshal(dataBody, &nig)
		if err != nil {
			panic(err)
		}

		dataGuru := Guru{NIG: 0}
		for _, val := range Teachers {
			if val.NIG == nig.NIG {
				dataGuru = val
			}
		}
		if dataGuru.NIG == 0 {
			ctx.Writer.WriteHeader(404)
			_, err := ctx.Writer.Write([]byte("Data not found"))
			if err != nil {
				panic(err)
			}
			return
		}
		dataGuruJson, err := json.Marshal(dataGuru)
		if err != nil {
			panic(err)
		}
		_, err = ctx.Writer.Write(dataGuruJson)
		if err != nil {
			panic(err)
		}
		return
	}
}

func Lesson() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get 1 data mapel by id
		// ambil data mapel dari body
		defer ctx.Request.Body.Close()
		dataBody, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			panic(err)
		}
		id := struct {
			ID int `json:"id"`
		}{}
		err = json.Unmarshal(dataBody, &id)
		if err != nil {
			panic(err)
		}

		dataMapel := MataPelajaran{ID: 0}
		for _, val := range Lessons {
			if val.ID == id.ID {
				dataMapel = val
			}
		}
		if dataMapel.ID == 0 {
			ctx.Writer.WriteHeader(404)
			_, err := ctx.Writer.Write([]byte("Data not found"))
			if err != nil {
				panic(err)
			}
			return
		}
		dataMapelJson, err := json.Marshal(dataMapel)
		if err != nil {
			panic(err)
		}
		_, err = ctx.Writer.Write(dataMapelJson)
		if err != nil {
			panic(err)
		}
		return
	}
}

func LessonByNIS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Ambil NIS dari body request
		defer ctx.Request.Body.Close()
		dataBody, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			panic(err)
		}
		nisData := struct {
			NIS int `json:"nis"`
		}{}
		err = json.Unmarshal(dataBody, &nisData)
		if err != nil {
			panic(err)
		}

		// Cari kelas siswa dengan NIS yang sesuai
		var kelasSiswa int
		for _, siswa := range students {
			if siswa.NIS == nisData.NIS {
				kelasSiswa = siswa.Kelas
				break
			}
		}

		// Cari pelajaran yang sesuai dengan kelas siswa
		var lessonList []MataPelajaran
		for _, lesson := range Lessons {
			if lesson.Kelas == kelasSiswa {
				lessonList = append(lessonList, lesson)
			}
		}

		// Jika tidak ada pelajaran yang ditemukan, kirim status 404
		if len(lessonList) == 0 {
			ctx.Writer.WriteHeader(404)
			_, err := ctx.Writer.Write([]byte("Data not found"))
			if err != nil {
				panic(err)
			}
			return
		}

		// Kirim data pelajaran sebagai JSON
		lessonListJson, err := json.Marshal(lessonList)
		if err != nil {
			panic(err)
		}
		ctx.Writer.WriteHeader(200)
		_, err = ctx.Writer.Write(lessonListJson)
		if err != nil {
			panic(err)
		}
		return
	}
}

func GetGin() *gin.Engine {
	r := gin.Default()

	r.GET("/students", StudentGet())
	r.POST("/students", StudentPost())

	r.GET("/teachers", TeacherGet())
	r.POST("/teachers", TeacherPost())

	r.POST("/lesson", Lesson())
	r.POST("/lessons", LessonByNIS())

	return r
}

func main() {
	//http.ListenAndServe("localhost:8080", GetMux())
	if err := GetGin().Run("localhost:8080"); err != nil {
		log.Println(err)
	}
}
