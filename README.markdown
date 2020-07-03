go-ecdict
===

基于[ECDICT](https://github.com/skywind3000/ECDICT)库的数据实现单次的查询功能。

目前实现的查询方式有：

+ `SimpleDict` 使用完全常驻在内存中的有序列表作为索引，用二分法查找。索引比较大，需要 `70M ~ 80M`
  的内存

## Benchmark

| 名称 | 次数 | 平均耗时 | 平均内存占用 | 平均内存分配 |
|--- |--- |--- |--- |--- |
|BenchmarkSimpleDict_Match-4 |116396 |8676 ns/op |5154 B/op |14 allocs/o |
