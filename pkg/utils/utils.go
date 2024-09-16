/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	// valid "github.com/asaskevich/govalidator/v11"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/labstack/echo/v4"
)

func PrepareRoutesTable(routes []*echo.Route, caption string) table.Writer {
	t := table.NewWriter()
	t.SetTitle(caption)
	t.AppendHeader(table.Row{"Method", "Path", "Handler"})

	for _, route := range routes {
		t.AppendRow(table.Row{route.Method, route.Path, route.Name})
	}
	t.SortBy([]table.SortBy{
		{Name: "Path", Mode: table.Asc},
		{Name: "Method", Mode: table.Asc},
	})
	t.AppendFooter(table.Row{"TOTAL", len(routes), ""})
	return t
}

func PrettyJSONString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func SetTableBorderStyle(t table.Writer, border bool) {
	t.SetStyle(table.StyleRounded)
	if border {
		t.Style().Options.DrawBorder = false
		t.Style().Options.SeparateColumns = false
		t.Style().Options.SeparateFooter = false
		t.Style().Options.SeparateHeader = false
		t.Style().Options.SeparateRows = false
	}
}

func ResolveArgs(args []string) map[string]string {
	m := make(map[string]string)
	for _, arg := range args {
		parts := strings.Split(arg, "=")
		if len(parts) < 2 {
			panic(fmt.Sprintf("Invalid argument (format: <key>=<value>): %s", arg))
		}
		m[parts[0]] = parts[1]
	}
	return m
}

func IntFromStr(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func StrInArr(s string, arr []string) bool {
	for _, a := range arr {
		if a == s {
			return true
		}
	}
	return false
}

func Exit(val int) {
	os.Exit(val)
}

// // IsEmpty checks if a string is empty
// func IsEmpty(str string) (bool, string) {
// 	if valid.HasWhitespaceOnly(str) && str != "" {
// 		return true, "Must not be empty"
// 	}

// 	return false, ""
// }

// // ValidateRegister func validates the body of user for registration
// func ValidateRegister(u *models.User) {
// 	e := &models.UserErrors{}
// 	e.Err, e.Username = IsEmpty(u.Username)

// 	if !valid.IsEmail(u.Email) {
// 		e.Err, e.Email = true, "Must be a valid email"
// 	}

// 	re := regexp.MustCompile("\\d") // regex check for at least one integer in string
// 	if !(len(u.Password) >= 8 && valid.HasLowerCase(u.Password) && valid.HasUpperCase(u.Password) && re.MatchString(u.Password)) {
// 		e.Err, e.Password = true, "Length of password should be atleast 8 and it must be a combination of uppercase letters, lowercase letters and numbers"
// 	}

// 	return e
// }
