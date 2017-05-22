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

package module

import (
	"fmt"
	"github.com/pkg/errors"
)

var Online = make(map[string]string)

func PutUser(key, value string) {
	Online[key] = value
}

func RemoveUser(key string) error {
	found := CheckIsOnline(key)
	if !found {
		return errors.New("The key: " + key + " is not exist.")
	}
	delete(Online, key)
	return nil
}

func CheckIsOnline(key string) bool {
	found := false

	if _, ok := Online[key]; ok {
		found = true
	}

	return found
}

func LogOnline() {
	for k, v := range Online {
		fmt.Println(k, v)
	}
}
