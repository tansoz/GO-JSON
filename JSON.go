package GO_JSON

import (
	"fmt"
)

func Parse (json string)*MapList{
	if (json == "") {
		return nil
	}
	inStr := false
	r := Map()
	FLOOR_T := false // false => {}, true => []
	STR_T := false   // false => ", true => '
	deep := -1;
	KEY := ""
	isDown := false
	intKEY := 0
	VALUE := ""
	TMP := ""
	TYPE := -1 // 0 => 键,  1 => 值, 2 => 子
	get := true
	first := false
	i := 0
	for i = range json {
		char := json[i];
		get = true
		if (isDown) {
			isDown = false
			deep -= 1
		}
		switch (char) {
		case '{':
			if (!inStr) {
				deep += 1
				if !first{
					FLOOR_T = false
					first = true
				}
				if(deep==0){
					TYPE = 0
					get = false
				}else{
					TYPE = 2
				}
			}
			break
		case '[':
			if (!inStr) {
				deep += 1
				if !first{
					FLOOR_T = true
					first = true
				}
				if(deep==0){
					get = false
					KEY = fmt.Sprintf("%d",intKEY)
					intKEY++
					TYPE = 1
				}else{
					TYPE = 2
				}
			}
			break
		case '}':
			if (!inStr) {
				isDown = true
				//FLOOR_T = false
				if(deep==0){
					get = false
					if(first&&!FLOOR_T){
						first = false
					}else{
						isDown = false
						get = true
					}
				}

			}
			break
		case ']':
			if (!inStr) {
				isDown = true
				//FLOOR_T = true
				if(deep==0){
					get = false
					if(first&&FLOOR_T){
						first = false
					}else{
						isDown = false
						get = true
					}
				}
			}
			break
		case '"':
			if (!inStr) {
				inStr = true
				get = false
				STR_T = false
			} else if (inStr && !STR_T && i-1 >= 0 && json[i-1] != '\\') {
				inStr = false
				get = false
			}
			if(TYPE==2){
				get = true
			}
			break
		case ' ':
			if(!inStr&&deep==0){
				get = false
			}
		case ':':
			if (!inStr) {
				if(deep==0){
					TYPE = 1
					get = false
				}
			}
			break
		case ',':
			if(!inStr&&deep==0){
				if(TYPE==2){
					r.Add(KEY,Parse(TMP))
				}else{
					r.Add(KEY,VALUE)
				}
				KEY = ""
				VALUE = ""
				TMP = ""
				if(deep==0){
					if(FLOOR_T){
						KEY = fmt.Sprintf("%d",intKEY)
						intKEY++
						TYPE = 1
					}else{
						TYPE = 0
					}
					get = false
				}else{
					TYPE = 2
				}
			}
		}
		if(get&&first){

			if(TYPE==0){
				KEY += string([]byte{char})
			}else if(TYPE==1){
				VALUE += string([]byte{char})
			}else if(TYPE==2){
				TMP += string([]byte{char})
			}
		}
	}
	if(i>1){
		if(TYPE==2){
			r.Add(KEY,Parse(TMP))
		}else{
			r.Add(KEY,VALUE)
		}
	}
	return r
}
