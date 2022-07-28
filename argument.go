package app

import (
	"errors"
	"os"
	"strings"
)

var args = os.Args[1:]

/*
	Enter the list of allowed arguments,
	if there exists an argument that is not in the above list, it will return an error.

	This function is optional, if you don't call it, the next operations will still happen normally.
*/
func allow_arguments(allow ...string) error {
	var regexPrefix = func(list []string, arg string) bool {
		for _, element := range list {
			if strings.HasPrefix(arg, element) {
				return true
			}
		}
		return false
	}
	for _, arg := range args {
		if !regexPrefix(allow, arg) {
			list := func() (s string) {
				for i, arg := range allow {
					if i == 0 {
						s = "[" + arg
						continue
					}
					s += "," + arg
				}
				return s + "]"
			}()
			println("log/ regexPrefix(", list, ",", arg, ") == false")
			return errors.New("argument " + arg + " do not support, please enter \"command\" --help to about.")
		}
	}
	return nil
}

func get_argument(full string, shorten ...string) (value string, ok bool, err error) {
	for _, arg := range args {
		// kiểm tra đối số đầy đủ có tồn tại không
		if strings.HasPrefix(arg, full) {
			// kiểm tra đối số đã xét trước đó chưa
			if ok || len(value) > 0 {
				println("log/ argument", full, "is duplicated.")
				return "", false, errors.New("input arguments are duplicated, please enter \"command\" --help to about")
			}
			// lấy giá trị của đối số nếu có
			if arr := strings.Split(arg, "="); len(arr) == 2 {
				value = arr[1]
			}
			// nhắc rằng đối số tồn tại
			ok = true
		}
		// kiểm tra đối số rút gọn có được đặt không
		for _, short := range shorten {
			if arg == short {
				// kiểm tra đối số đã xét trước đó chưa
				if ok || len(value) > 0 {
					println("log/ argument", shorten, "is duplicated.")
					return "", false, errors.New("input arguments are duplicated, please enter \"command\" --help to about")
				}
				// nhắc rằng đối số tồn tại
				ok = true
			}
		}
	}
	return value, ok, func() error {
		if ok {
			return err
		} else {
			return nil
		}
	}()
}
