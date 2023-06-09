# GO的异常机制

> panic触发宕机
> 
> defer延时执行
>
> recover宕机时恢复执行
> 
> 值类型、引用类型与深浅拷贝
>
> 类型别名与自定义
> 
> new 和 make 的区别
>

## 异常处理关键字：defer、panic 和 recover

panic：通过内置函数 panic() 表示，设有参数v（空接口类型），参数v代表异常信息；

defer：使用了会在最后执行，美欧返回值则在 代码执行完成后 和 函数返回值 之间执行；

recover：使用了 panic 抛出异常会停止执行。只有和 defer 结合实现异常捕捉和处理

## 值类型、引用类型与深浅拷贝
变量拷贝概念：将一个变量的数据复制并存到另一个变量中。主要是保存数据处理前的变量而重新定义新的变量。
```text
变量拷贝分为： 深拷贝 和 浅拷贝，这两者的区别在于 变量之间是否共用一个内存地址。
深拷贝：不共用；
浅拷贝共用：一个改变，另一个的值也改变；
```

```text
值类型变量：整型、布尔型、浮点型、字符串、数据、结构体，是直接存储数据，内存通常在栈中分配。
值类型变量的 赋值 到另一个变量是 深拷贝；

引用类型变量是变量储存 内存地址，这个内存地址再存数据，内存在堆上，通过GC回收。
引用类型：指针、切片、集合、通道和接口等。
引用类型的数据赋值到另一个变量都是浅拷贝；
```

## 类型别名与自定义
```text
类型别名与类型自定义是由关键字type实现的；
1.类型别名是对已有数据赋予新命名，主要用来解决代码升级、迁移中存在的类型兼容性问题。
2.类型定义是自定义数据类型，但是这种定义是在已有数据类型基础上进行定义的，最常见的包括结构体和接口。
```

## new 和 make 的区别 
```text
内置函数 new() 和 make() 用于内存分配：
new()只分配内存，
make() 为切片、集合以及通道的数据类型分配内存和初始化。
```

## 泛型的概念与应用
```text
泛型全称为泛型程序设计（Generic Programming），它是一种分割或范式。它允许我们在
强类型编程语言中实例化对某个对象的时候才指明参数的数据类型。

一般来说定义 函数方法 的时候，必须对参数和返回值设置数据类型，
使用过程中必须严格按照定义来使用，不然会报错。

若想要函数的参数和返回值不受数据类型限制，可以将参数和返回值设置成空接口，
空接口能给使用者传递任意数据类型的数据，但如果 函数只允许传递部分数据类型的数据，则需要泛型实现。
传

有了空接口为什么还要用泛型？
因为空接口不受数据类型的限制，如果调用过程中，函数传入参数是无法处理的数据类型，则容易引起异常。
使用泛型可以保证参数（返回值）类型的多样性，也能保证调用过程中不会传入非法参数。
```
