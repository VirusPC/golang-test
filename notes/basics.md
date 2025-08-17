# Basics

0. 初始化仓库

    ```bash
    go mod init github.com/VirusPC/golang-test
    ```

1. package name。按照惯例，包名与导入路径的最后一个元素相同。例如，"math/rand" 包中的文件都以 package rand 语句开头。

    - Java 的包名严格映射到文件路径。例如，包 com.example.utils 对应的文件路径是 src/com/example/utils。
    - Java 的文件导出内容并不按包名最后一个元素导出，具体导出内容根据文件内容而定。一个文件只能有一个 public 类，该类的类名必须与文件名相同。
  
2. 导入。下面两个代码等价

    ```go
    import "fmt"
    import "math"

    import (
        "fmt"
        "math"
    )
    ```

3. 导出。在Go语言中，一个名称如果以大写字母开头则表示它是被导出的，例如Pizza和math包中的Pi

    ```go
    var (
      ToBe   bool       = false
      MaxInt uint64     = 1<<64 - 1
      Z      complex128 = cmplx.Sqrt(-5 + 12i)
    )
    ```

4. 函数定义

    ```golang
    /* 定义 swap 函数，交换两个字符串
     * 1. 连续两个参数类型相同时，可以省略之前的类型
     * 2. 函数s可以有多个返回值
     */
    func swap(x, y string) (string, string) {
      return y, x
    }
    ```

    ```java
    // 必须放到类中
    public class Main {
        // 定义 swap 方法
        public static string swap(string x, string y) {
            return [y, x];
        }
    }
    ```

5. 变量定义。可以是 function 或 package level

    ```golang
    // 定义变量
    var c, python, java bool // 连续定义多个相同类型的变量
    var i, j int = 1, 2 // 类似 js 中的解构赋值

    func main() {
      k := 3 // 函数内，可以使用 := 来省略变量类型，implicit type. 自动类型推导。
      i := 42           // int 
      f := 3.142        // float64
      g := 0.867 + 0.5i // complex128
    }
    ```

    基础类型如下：

    ```md
    bool

    string

    int  int8  int16  int32  int64
    uint uint8 uint16 uint32 uint64 uintptr

    byte // alias for uint8

    rune // alias for int32
        // represents a Unicode code point

    float32 float64

    complex64 complex128
    ```

    没有初始化的变量会被设置为零值。

    - 数值类型为 0
    - 字符串为空字符串
    - 布尔类型为 false
    - 指针、函数、接口、切片、通道和映射为 nil

6. 类型转换 (Type conversions)

    ```go
    var i int = 42
    var f float64 = float64(i)
    var u uint = uint(f)

    // 或更简单地（函数中）
    i := 42
    f := float64(i)
    u := uint(f)
    ```

7. 常量

    ```go
    const World = "世界" // 不可以用 := 定义常量
    ```
