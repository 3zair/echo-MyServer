package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"MyServer/module"
	"MyServer/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"MyServer/sqlHelper"
	"MyServer/utils"
)

func LoginHandler(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

func RegisterHandler(c echo.Context) error {
	var status = config.Success

	name := c.QueryParam("name")
	password := c.QueryParam("password")
	repeat := c.QueryParam("repeat")

	if name == "" || password == "" || repeat == "" {
		status = config.ParamErr
	} else if password != repeat {
		status = config.PasswordNotMarry
	} else {
		//保存至数据库
		db, err := sql.Open("mysql", "root:123456@/MyServer?charset=utf8")
		if err != nil {
			panic(err.Error())
		}

		defer db.Close()

		sqlHelper.Insert(db, "INSERT INTO user (name, password) VALUES (?, ?)", name, utils.MD5(password))
	}

	status_json := module.Register{
		Status: status,
	}

	return c.JSONPretty(http.StatusOK, status_json, "  ")
}
