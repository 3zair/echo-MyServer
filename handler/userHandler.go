/*
 * MIT License
 *
 * Copyright (c) 2017 Tang Xiaoji.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"MyServer/module"
	"MyServer/config"
	_ "github.com/go-sql-driver/mysql"
	"MyServer/sqlHelper"
	"MyServer/utils"
	"fmt"
)

func LoginHandler(c echo.Context) error {
	user := new(module.User)

	if err := c.Bind(user); err != nil {
		return err
	}

	rowMap, err := sqlHelper.FetchRow(sqlHelper.Db, "SELECT * FROM `user` WHERE `name` = ? LIMIT 0, 100", user.Name)

	if err != nil {
		return err
	}

	//数据库里查到的密码
	password := (*rowMap)["password"]

	sess := utils.GlobalSessions.SessionStart(c.Response().Writer, c.Request())

	if !utils.CompareHash([]byte(password), user.Password) {
		return c.JSON(config.ErrIncorrectPass, user)
	}

	sess.Set("login", user.Name)

	status_json := &module.Err{
		Status: config.ErrSucceed,
		Data: user.Name,
	}

	module.PutUser(user.Name, user.Name)

	module.LogOnline()

	return c.JSON(http.StatusOK, status_json)
}

func RegisterHandler(c echo.Context) error {
	var status = config.ErrSucceed

	name := c.QueryParam("name")
	password := c.QueryParam("password")
	repeat := c.QueryParam("repeat")

	fmt.Println(name + " " + password + " " + repeat)

	if name == "" || password == "" || repeat == "" {
		status = config.ErrInvalidParam
	} else if password != repeat {
		status = config.ErrIncorrectPass
	} else {
		//保存至数据库
		pass, err := utils.GenerateHash(password)

		if err != nil {
			return err
		}

		sqlHelper.Insert(sqlHelper.Db, "INSERT INTO user (name, password) VALUES (?, ?)", name, string(pass))
	}

	status_json := &module.Err{
		Status: status,
	}

	return c.JSONPretty(http.StatusOK, status_json, "  ")
}

func Logout(c echo.Context) error {
	status := config.ErrSucceed

	sess := utils.GlobalSessions.SessionStart(c.Response().Writer, c.Request())
	name := sess.Get("login")
	err := module.RemoveUser(name.(string))

	if err != nil {
		status = config.ErrLoginRequired
		return err
	}

	sess.Delete("login")

	status_json := &module.Err{
		Status: status,
	}

	module.LogOnline()

	return c.JSONPretty(http.StatusOK, status_json, " ")
}

func ReviseInfo(c echo.Context) error {
	return nil
}
