package GO-JSON

import (
	"fmt"
	"strings"
)
type data struct {
	obj interface{}
	obj_t int
}
func (this *data)ToString()string{
	return this.obj.(string)
}
func (this *data)ToMapList()*MapList{
	return this.obj.(*MapList)
}
func (this *data)ToBool()bool{
	if this.obj_t == STRING{
		return parseBool(this.obj.(string))
	}else{
		return this.obj.(bool)
	}
}

func parseBool(s string)bool {
	for i := range s {
		char := s[i];
		if (char == 't'||char == 'T') && i + 3 <= len(s[i:]) && (s[i+1] == 'r'||s[i+1] == 'R') &&(s[i+1] == 'u'||s[i+1] == 'U') &&(s[i+1] == 'e'||s[i+1] == 'E')  {
			return true
		}
	}
	return false
}
func (this *data)ToInt()int{
	if this.obj_t == STRING{
		return parseInt(this.obj.(string))
	}else{
		return this.obj.(int)
	}
}
func (this *data)ToFloat64()float64{
	if this.obj_t == STRING{
		return parseFloat64(this.obj.(string))
	}else{
		return this.obj.(float64)
	}
}
func parseFloat64(s string)float64{
	dot := strings.Index(s,".")
	I := float64(parseInt(s[0:dot]))
	F := parseFloat(s[dot+1:])
	if I > 0{
		I += F
	}else{
		I -= F
	}
	return I
}
func parseFloat(s string)float64{
	var r,d float64
	r = 0.0;
	start := false;
	d = 1.0;
	for i := range s{
		char := s[i];
		switch char {
		case '1':
			if(!start){
				start = true
				r += 1;
			}else{
				r = r*10+1
			}
			d *= 10
			break;
		case '2':
			if(!start){
				start = true
				r += 2;
			}else{
				r = r*10+2
			}
			d *= 10
			break;
		case '3':
			if(!start){
				start = true
				r += 3;
			}else{
				r = r*10+3
			}
			d *= 10
			break
		case '4':
			if(!start){
				start = true
				r += 4;
			}else{
				r = r*10+4
			}
			d *= 10
			break
		case '5':
			if(!start){
				start = true
				r += 5;
			}else{
				r = r*10+5
			}
			d *= 10
			break
		case '6':
			if(!start){
				start = true
				r += 6;
			}else{
				r = r*10+6
			}
			d *= 10
			break
		case '7':
			if(!start){
				start = true
				r += 7;
			}else{
				r = r*10+7
			}
			d *= 10
			break
		case '8':
			if(!start){
				start = true
				r += 8;
			}else{
				r = r*10+8
			}
			d *= 10
			break
		case '9':
			if(!start){
				start = true
				r += 9;
			}else{
				r = r*10+9
			}
			d *= 10
			break
		case '0':
			if start{
				start = true
				r += 9;
			}else{
				r = r*10
			}
			d *= 10
			break
		default:
			if start {
				return (r/d)
			}
			break
		}
	}
	return r/d;
}
func parseInt(s string)int{
	r := 0;
	start := false;
	direct := 1;
	for i := range s{
		char := s[i];
		switch char {
		case '-':
			if(!start){
				start = true
				direct = -1
			}
			break;
		case '1':
			if(!start){
				start = true
				r += 1;
			}else{
				r = r*10+1
			}
			break;
		case '2':
			if(!start){
				start = true
				r += 2;
			}else{
				r = r*10+2
			}
			break;
		case '3':
			if(!start){
				start = true
				r += 3;
			}else{
				r = r*10+3
			}
			break
		case '4':
			if(!start){
				start = true
				r += 4;
			}else{
				r = r*10+4
			}
			break
		case '5':
			if(!start){
				start = true
				r += 5;
			}else{
				r = r*10+5
			}
			break
		case '6':
			if(!start){
				start = true
				r += 6;
			}else{
				r = r*10+6
			}
			break
		case '7':
			if(!start){
				start = true
				r += 7;
			}else{
				r = r*10+7
			}
			break
		case '8':
			if(!start){
				start = true
				r += 8;
			}else{
				r = r*10+8
			}
			break
		case '9':
			if(!start){
				start = true
				r += 9;
			}else{
				r = r*10+9
			}
			break
		case '0':
			if start{
				r = r*10
			}
			break
		default:
			if start {
				return r
			}
			break
		}
	}
	return r*direct;
}
const(
	ERR = -1
	INT = 0
	INT_ARR = 1
	UINT = 2
	INT8 = 3
	INT16 = 4
	INT32 = 5
	INT64 = 6
	UINT8 = 7
	UINT16 = 8
	UINT32 = 9
	UINT64 = 10
	STRING = 11
	STRING_ARR = 12
	FLOAT32 = 13
	FLOAT64 = 14
	COMPLEX64 = 15
	COMPLEX128 = 16
	MAPLIST = 17
	OBJECT = 18
	BYTE = 7
	BYTE_ARR = 20
	BOOL = 21
	DATA = 22
	DATA_PT = 23
)

type MapNode struct {
	key string
	value data
	next *MapNode
}
type MapList struct {
	first * MapNode
	current * MapNode
	Len int
}

func Map()*MapList{
	mn := &MapNode{"",data{nil,ERR},nil}
	ml := &MapList{mn,mn,0}
	return ml
}

func (this *MapList)Add(key string,value interface{}){
	if(strings.Trim(key," ")==""){
		return;
	}
	t := ERR
	switch value.(type) {
	case int:
		t = INT
		break
	case []int:
		t = INT_ARR
		break
	case int8:
		t = INT8
		break
	case int16:
		t = INT16
		break
	case int32:
		t = INT32
		break
	case int64:
		t = INT64
		break
	case uint:
		t = UINT
		break
	case uint8:
		t = UINT8
		break
	case uint16:
		t = UINT16
		break
	case uint32:
		t = UINT32
		break
	case uint64:
		t = UINT64
		break
	case string:
		t = STRING
		break
	case []string:
		t = STRING_ARR
		break
	case float32:
		t = FLOAT32
		break
	case float64:
		t = FLOAT64
		break
	case complex64:
		t = COMPLEX64
		break
	case complex128:
		t = COMPLEX128
		break
	case *MapList:
		t = MAPLIST
		break
	case []byte:
		t = BYTE_ARR
		break
	case bool:
		t = BOOL
		break
	case *data:
		t = DATA_PT
		break
	case data:
		t = DATA
		break
	default:
		t = OBJECT
		break
	}
	this.current.key = key
	this.current.value = data{value,t}
	this.current.next = &MapNode{"",data{nil,ERR},nil}
	this.current = this.current.next
	this.Len += 1

}
func (this * data)String()string{
	switch this.obj_t {
	case MAPLIST:
		return this.obj.(*MapList).String()
		break
	case STRING:
		return this.obj.(string)
		break
	case INT:
		return fmt.Sprintf("%d",this.obj.(int))
		break
	case DATA:
		return (this.obj.(* data)).String()
		break
	}
	return "[Object]"
}

func (this *MapList)Del(key string)bool{

	root := this.first
	pre := root
	for{
		if(root==nil||root.key==""){
			break;
		}
		if(key==root.key){
			if(root==this.first){
				this.first = root.next
			}else{
				pre.next = root.next
			}
			this.Len -= 1
			return true
		}else{
			pre = root
			root = root.next
		}
	}

	return false

}

func (this *MapList)Get(key string,limit int8)*MapList{
	r := Map()
	var i int8
	i = 0;
	root := this.first
	for{
		if(root==nil||root.key==""){
			break;
		}
		if(key==root.key){
			r.Add(fmt.Sprintf("%d",i),root.value.obj)
			i++
			if limit != -1&& limit < i {
				break
			}
		}
		root = root.next
	}
	return r

}

func (this *MapList)Index(index int)*data{
	var i int
	i = 0;
	root := this.first
	for{
		if(root==nil||root.key==""){
			break;
		}
		if(index == i){
			return &root.value
		}
		i++
		root = root.next
	}
	return &data{Map(),ERR}
}

func (this *MapList)String()string{
	root := this.first
	r := "{"
	for{
		if(root==nil||root.key==""){
			break;
		}

		switch root.value.obj_t{

		case MAPLIST:
			r += "\""+root.key+"\"=>"+ (root.value.obj.(*MapList)).String()
			break
		case STRING:
			r += "\""+root.key+"\"=>\""+(root.value.obj.(string))+"\""
			break
		case INT:
			r += "\""+root.key+"\"=>\""+fmt.Sprintf("%d",(root.value.obj.(int)))+"\""
			break
		default:
			r += "\""+root.key+"\"=>[Object]"
			break
		}
		if(root.next.key!=""){
			r += ","
		}

		root = root.next

	}
	r += "}"
	return r
}
