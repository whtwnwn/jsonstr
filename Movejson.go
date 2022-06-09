package jsonstr

import (
	"errors"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"strconv"
	"strings"
)

func Movejson(jsonstr []byte, dstpath string, sourecepath string) (string, error) {
	count := strings.Count(dstpath, "#")
	if count > 1 {
		str := fmt.Sprintf("\"dstpath must only on # or *,but dstpath is %d\"", count)
		return "", errors.New(str)
	}
	count = strings.Count(sourecepath, "#")
	if count > 1 {
		str := fmt.Sprintf("\"sourecepath must only on # or *,but sourecepath is %d\"", count)
		return "", errors.New(str)
	}
	if strings.Contains(dstpath, "#") {
		dstpath = strings.Replace(dstpath, "#", "*", 1)
	}
	if strings.Contains(sourecepath, "#") {
		sourecepath = strings.Replace(sourecepath, "#", "*", 1)
	}
	jObj, err := gabs.ParseJSON(jsonstr)
	if err != nil {
		return "", err
	}
	if count == 1 {
		for k, v := range jObj.Path(sourecepath).Children() {
			ks := strconv.Itoa(k)
			path := strings.Replace(dstpath, "*", ks, 1)
			value := v.Data()
			jObj.SetP(value, path)
		}
	} else {
		value := jObj.Path(sourecepath).Data()
		jObj.SetP(value, dstpath)
	}
	return jObj.String(), nil
}
func MapingAndSplit(jsonstr []byte, srcpath string, dstpath map[string]string, mapping map[string]map[string]string) (string, error) {
	count := strings.Count(srcpath, "#")
	if count > 1 {
		str := fmt.Sprintf("\"srcpath must only on # or *,but dstpath is %d\"", count)
		return "", errors.New(str)
	}
	jObj, err := gabs.ParseJSON(jsonstr)
	if err != nil {
		return "", err
	}
	for key, value := range dstpath {
		count = strings.Count(value, "#")
		if count > 1 {
			str := fmt.Sprintf("\"%s must only on # or *,but dstpath is %d\"", key, count)
			return "", errors.New(str)
		}
		if count == 1 {
			for k, v := range jObj.Path(srcpath).Children() {
				ks := strconv.Itoa(k)
				path := strings.Replace(value, "*", ks, 1)
				values, ok := mapping[v.Data().(string)]
				if !ok {
					return "", errors.New("mapping is null")
				}
				jObj.SetP(values[key], path)
			}
		} else {
			v := jObj.Path(srcpath).Data()
			values, ok := mapping[v.(string)]
			if !ok {
				return "", errors.New("mapping is null")
			}
			jObj.SetP(values[key], value)
		}
	}
	return jObj.String(), nil
}
func Mapping(jsonstr []byte, srcpath string, mapping map[string]string) (string, error) {
	count := strings.Count(srcpath, "#")
	if count > 1 {
		str := fmt.Sprintf("\"srcpath must only on # or *,but dstpath is %d\"", count)
		return "", errors.New(str)
	}
	jObj, err := gabs.ParseJSON(jsonstr)
	if err != nil {
		return "", err
	}
	if count == 1 {
		for k, v := range jObj.Path(srcpath).Children() {
			ks := strconv.Itoa(k)
			path := strings.Replace(srcpath, "*", ks, 1)
			values, ok := mapping[v.Data().(string)]
			if !ok {
				return "", errors.New("mapping is null")
			}
			jObj.SetP(values, path)
		}
	} else {
		tmp := jObj.Path(srcpath).Data()
		values, ok := mapping[tmp.(string)]
		if !ok {
			return "", errors.New("mapping is null")
		}
		jObj.SetP(values, srcpath)
	}
	return jObj.String(), nil
}
