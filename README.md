# GO-JSON
GO Language JSON Parser

一个Go语言的JSON解析器

## 用法:

### JSON.parse([JSON String]:string)

    ml := JSON.Parse("{\"name\":\"tansoz\",\"num\":[0,3,{\"name\":[\"abc\",\"abd\"]}]}")

    fmt.Println(ml)
    
    --------------------------------------------

    Result: {"name"=>"tansoz","num"=>{"0"=>"0","1"=>"3","2"=>{"name"=>{"0"=>"abc","1"=>"abd"}}}}

### Get([name]:string,[limit]:int)

    ml := JSON.Parse("{\"name\":\"tansoz\",\"num\":[0,3,{\"name\":[\"abc\",\"abd\"]}]}")

    fmt.Println(ml.Get("name",1))
    
    --------------------------------------------

    Result: {"0"=>"tansoz"}

### Index([index]:int)

    ml := JSON.Parse("{\"name\":\"tansoz\",\"num\":[0,3,{\"name\":[\"abc\",\"abd\"]}]}")
    
    fmt.Println(ml.Index(1))
    
    --------------------------------------------

    Result: {"0"=>"0","1"=>"3","2"=>{"name"=>{"0"=>"abc","1"=>"abd"}}}
    
### ToInt / ToMapList / ToString / ToFloat64 / ToBool

就以ToInt作为例子使用演示

    ml := JSON.Parse("{\"name\":\"tansoz\",\"num\":[0,3,{\"name\":[\"abc\",\"abd\"]}]}")
    
    fmt.Println(ml.Index(1).ToMapList().Index(1).ToInt()+10)
    
    --------------------------------------------

    Result: 13

## Usage:
  
### JSON.parse([JSON String]:string)

    ml := JSON.Parse("{\"name\":\"tansoz\",\"num\":[0,3,{\"name\":[\"abc\",\"abd\"]}]}")

    fmt.Println(ml)
    
    --------------------------------------------

    Result: {"name"=>"tansoz","num"=>{"0"=>"0","1"=>"3","2"=>{"name"=>{"0"=>"abc","1"=>"abd"}}}}

### Get([name]:string,[limit]:int)

    ml := JSON.Parse("{\"name\":\"tansoz\",\"num\":[0,3,{\"name\":[\"abc\",\"abd\"]}]}")

    fmt.Println(ml.Get("name",1))
    
    --------------------------------------------

    Result: {"0"=>"tansoz"}

### Index([index]:int)

    ml := JSON.Parse("{\"name\":\"tansoz\",\"num\":[0,3,{\"name\":[\"abc\",\"abd\"]}]}")
    
    fmt.Println(ml.Index(1))
    
    --------------------------------------------

    Result: {"0"=>"0","1"=>"3","2"=>{"name"=>{"0"=>"abc","1"=>"abd"}}}
    
### ToInt / ToMapList / ToString / ToFloat64 / ToBool

Let me use the ToInt function as a example to show how to use these function

    ml := JSON.Parse("{\"name\":\"tansoz\",\"num\":[0,3,{\"name\":[\"abc\",\"abd\"]}]}")
    
    fmt.Println(ml.Index(1).ToMapList().Index(1).ToInt()+10)
    
    --------------------------------------------

    Result: 13
  
