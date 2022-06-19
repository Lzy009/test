package main

import (
	"errors"
	"fmt"
	"net/http"
	"path"
	"path/filepath"
	"testmy/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Video struct {
	//gorm.Model
	ID   string //path
	Type string
	//Auth int64
}

func Registry(id string) {

}
func upload(c *gin.Context, db *gorm.DB) error {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 0,
			"msg":  "err",
			"data": "null",
		})
		return err
	}
	//time.Unix().String()
	file_type := path.Ext(file.Filename)
	dir := "./static/20220518"
	//if err := os.MkdirAll(dir, 0666); err != nil {
	//	c.String(http.StatusBadRequest, "Mkdir失败")
	//	return
	//}
	fileUnixName := file.Filename
	//strconv.FormatInt(time.Now().UnixNano(), 10)
	data := Video{ID: fileUnixName, Type: file_type}
	db.Create(&data)
	saveDir := path.Join(dir, fileUnixName)
	Err := c.SaveUploadedFile(file, saveDir)
	return Err
}

func solve(c *gin.Context) {

	db, err := gorm.Open("mysql", fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=True&loc=Local", utils.Psd))
	if err != nil {
		return
	}
	defer db.Close()
	db.AutoMigrate(&Video{})
	if err != nil {
		fmt.Printf("upload err:%s\n", err)
	}
	if err := upload(c, db); err != nil {
		fmt.Printf("solve err:%s", err)
	}
}
func Push(c *gin.Context, db *gorm.DB) error {
	name := c.Param("name")
	var video Video
	db.Where("id = ?", name).First(&video)
	dir := "./static/20220518"
	suf := video.ID
	path := filepath.Join(dir, suf)
	checkPathStr := filepath.Ext(path)
	if checkPathStr == "" {
		return errors.New("no such file")
	}
	c.File(path)
	return nil
}
func download(c *gin.Context) {
	db, err := gorm.Open("mysql", fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=True&loc=Local", utils.Psd))
	if err != nil {
		return
	}
	defer db.Close()
	if err != nil {
		fmt.Printf("db open err:%s\n", err)
	}
	if err := Push(c, db); err != nil {
		fmt.Printf("Push err :%s\n", err)
	}
}

type Test struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
}

func test1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"pkj": "tset1",
	})
}
func test2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"pkj": "tset2",
	})
	//c.Writer(http.StatusOK, "123")
}
func test3(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"sec": "test3",
	})
}
func main() {
	//s := "1 = 1"
	//var Out Video
	r := gin.Default()
	r.GET("/test/t1", test1)
	r.Use(test2)
	r.GET("/test/t2", test3)
	r.Run()
	//db, _ := gorm.Open("mysql", fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=True&loc=Local", utils.Psd))
	//db.AutoMigrate(&Test{})
	//db.Create(&Test{})
	//dbRes := db.Model(&Video{}).Where("id = ?").First(&Out)
	//fmt.Println(dbRes.Error)
	// r := gin.Default()
	// r.GET("/upload", solve)
	// r.GET("/download/:name", download)

	// r.Run()

	//video := Video{ID: "1", Auth: int64(1)}
	//db.AutoMigrate(&Video{})
	//d := db.Create(video)
	//db.Table("video").Create(video)
	//fmt.Println(d.RowsAffected)
	//other.Print()
	//r := gin.Default()
	//r.GET("/upload", upload)
	//r.Run()
}
