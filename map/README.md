# map

该包用于测试map的特征  

## map可以直接在函数间通过参数传递  

当参数是一个字典时， 被调函数可以直接修改实参。  

## 不能直接对字典中的值的属性赋值  

```
m := make(map[string]struct{
    A string
})

m["hello"].A = "hello"
```  
上面的程序是错误的

## 使用字典中不存在的键执行赋值或删除不会报错
