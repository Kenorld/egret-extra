package www

import (
	"bitbucket.org/kenorld/eject-core"
	"gopkg.in/mgo.v2/bson"

	"bitbucket.org/kenorld/xlh-server/core/models"
)

type (
	user struct {
	}
)

// GET /users
func (u *user) List(ctx *eject.Context) {
	// Db := db.MgoDb{}
	// Db.Init()

	// result := []models.User{}
	// if err := Db.C("users").Find(nil).All(&result); err != nil {
	//ctx.RenderJSON("1")
	// } else {
	// 	ctx.RenderJSON(&result)
	// }

	ctx.RenderText("users")
	// Db.Close()
}

// GET /users/:param1
func (u *user) Show(ctx *eject.Context) {
	// Db := db.MgoDb{}
	// Db.Init()

	//	result, _ := models.User{}, ctx.Param("id")

	// if err := Db.C("people").Find(bson.M{"id": id}).One(&result); err != nil {
	//ctx.RenderJSON(result)
	// } else {
	// 	ctx.RenderJSON(result)
	// }
	ctx.RenderText("user id: " + ctx.Param("id"))
	// Db.Close()
}

// PUT /users
func (u *user) Create(ctx *eject.Context) {
	// newUsername := string(ctx.Form("username"))
	// myDb.InsertUser(newUsername)
	//println(newUsername + " has been inserted to database")
	//ctx.RenderJSON(map[string]interface{}{"response": true})

	ctx.RenderMarkdownFile("public/readme.md")
	// // Update
	// colQuerier := bson.M{"name": "Ale"}
	// change := bson.M{"$set": bson.M{"phone": "+86 99 8888 7777", "timestamp": time.Now()}}
	// err = c.Update(colQuerier, change)
	// if err != nil {
	// 	panic(err)
	// }

}

// POST /users/:param1
func (u *user) Update(ctx *eject.Context) {
	id := ctx.Param("id")
	usr := models.User{}
	err := ctx.Read(&usr)

	if err != nil {
		ctx.RenderJSON("4")
		panic(err.Error())
	}

	usr.ID = bson.ObjectId(id)

	// Db := db.MgoDb{}
	// Db.Init()

	// Insert
	// if err := Db.C("people").Insert(&usr); err != nil {
	// 	ctx.RenderJSON("5")
	// } else {
	// 	ctx.RenderJSON(map[string]interface{}{"response": true})
	// }

	// Db.Close()

}

// DELETE /users/:param1
func (u *user) Trash(ctx *eject.Context) {
}

var User = &user{}
