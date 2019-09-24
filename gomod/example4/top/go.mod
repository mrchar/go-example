module github.com/mrchar/go-example/gomod/example4/top

go 1.13

require sub v0.0.0-00010101000000-000000000000

replace sub => github.com/mrchar/go-example/gomod/example4/sub v0.0.0-20190924104040-cda95a8c6540 // indirect
