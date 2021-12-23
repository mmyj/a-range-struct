# 一个用于统计区间和的 range 结构体

## 题目
设计一个`range`结构体，它包含2个方法还有1个构造函数：
- `NewRange(nums []int)`：返回`range`结构体
- `ACC(left, right int)`：添加一个区间`[left, right]`
- `Total() int`：返回所有`ACC`方法添加的区间内元素之和的和

例如
```cgo
r := NewRange([1,2,3,4])
r.ACC(0,2) // 区间和是6
r.ACC(1,3) // 区间和是9
r.ACC(1,3) // 区间和是9
r.Total() // 3个区间和之和是6+9+9=24
```

## 题解
使用一个长度为`len(nums)`的数组`rangeDesc`记录调用`ACC`方法时的区间信息，`rangeDesc`的元素是一个结构体
```cgo
rangeDesc {
  l int // 左区间是该下标的区间数量
  r int // 右区间是该下标的区间数量
}
```
例如，例子中调用3次`ACC`方法之后可以得到的`rangeDesc`数组如下
```cgo
rangeDesc = [{l:1,r:0},{l:2,r:0},{l:0,r:1},{l:0,r:2}]
```
调用`Total`方法时，遍历`nums`数组，遍历期间维护一个统计变量`rangeCount`，`rangeCount`表示当前元素在几个区间内，`rangeCount`的计算规则是：
```cgo
rangeCount += r[i].l-r[i-1].r
```
为了方便计算，`rangeDesc`会从下标1开始，下标0的元素为零值

`Total`方法的计算规则是：
```cgo
total += nums[i]*rangeCount
```
每次调用`ACC`方法需要随机访问2次`rangeDesc`数组，调用`Total`方法需要遍历一次`nums`数组，所以算法复杂度是`O(m+n)`

## 测试
分别测试了长度为`50、100、200、500、1000、10000、100000`的`nums`数组，基准算法`RoughRange`使用的是每次调用`ACC`方法都对区间进行遍历求和累加，测试结果如下
```cgo
BenchmarkRangeRandComparison/50_RoughRange-8         	 5153847	       205.5 ns/op
BenchmarkRangeRandComparison/50_Range-8              	 6104293	       193.1 ns/op
BenchmarkRangeRandComparison/50_RangeWithRangeDescs-8         	 6041862	       192.0 ns/op
BenchmarkRangeRandComparison/100_RoughRange-8                 	 3100610	       385.5 ns/op
BenchmarkRangeRandComparison/100_Range-8                      	 3143158	       378.3 ns/op
BenchmarkRangeRandComparison/100_RangeWithRangeDescs-8        	 3184474	       372.1 ns/op
BenchmarkRangeRandComparison/200_RoughRange-8                 	 1420014	       834.2 ns/op
BenchmarkRangeRandComparison/200_Range-8                      	 1597068	       753.8 ns/op
BenchmarkRangeRandComparison/200_RangeWithRangeDescs-8        	 1643364	       720.8 ns/op
BenchmarkRangeRandComparison/500_RoughRange-8                 	  531212	      2252 ns/op
BenchmarkRangeRandComparison/500_Range-8                      	  624054	      1851 ns/op
BenchmarkRangeRandComparison/500_RangeWithRangeDescs-8        	  660788	      1767 ns/op
BenchmarkRangeRandComparison/1000_RoughRange-8                	  259494	      4476 ns/op
BenchmarkRangeRandComparison/1000_Range-8                     	  311188	      3720 ns/op
BenchmarkRangeRandComparison/1000_RangeWithRangeDescs-8       	  321549	      3520 ns/op
BenchmarkRangeRandComparison/10000_RoughRange-8                      8181      125834 ns/op
BenchmarkRangeRandComparison/10000_Range-8                          29268       38963 ns/op
BenchmarkRangeRandComparison/10000_RangeWithRangeDescs-8            33374       36028 ns/op
BenchmarkRangeRandComparison/100000_RoughRange-8                      810     1362657 ns/op
BenchmarkRangeRandComparison/100000_Range-8                          2450      422910 ns/op
BenchmarkRangeRandComparison/100000_RangeWithRangeDescs-8            2792      532971 ns/op
```
`RangeWithRangeDescs`测试是考虑到每次调用`NewRange`构造时都会`alloc`一个新的`rangeDesc`数组，这会影响测试结果，于是实现了一个复用`rangeDesc`数组的构造方法

测试结论是数据规模在`200`以下时，`ACC`时直接遍历求和的算法性能更好；大于`200`时，通过记录区间信息，最后求和的算法性能会更好，随着数据规模增大，性能差距大约是4倍