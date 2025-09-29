# 笔记复制

```
"Polymorphism is the ability to write code that can take on different behavior through the
 implementation of types. Once a type implements an interface, an entire world of
 functionality can be opened up to values of that type."
 - Bill Kennedy

"Interfaces are types that just declare behavior. This behavior is never implemented by the
 interface type directly, but instead by user-defined types via methods. When a
 user-defined type implements the set of methods declared by an interface type, values of
 the user-defined type can be assigned to values of the interface type. This assignment
 stores the value of the user-defined type into the interface value.

 If a method call is made against an interface value, the equivalent method for the
 stored user-defined value is executed. Since any user-defined type can implement any
 interface, method calls against an interface value are polymorphic in nature. The
 user-defined type in this relationship is often called a concrete type, since interface values
 have no concrete behavior without the implementation of the stored user-defined value."
  - Bill Kennedy

Receivers       Values
-----------------------------------------------
(t T)           T and *T
(t *T)          *T

Values          Receivers
-----------------------------------------------
T               (t T)
*T              (t T) and (t *T)


SOURCE:
Go In Action
William Kennedy
/////////////////////////////////////////////////////////////////////////

Interface types express generalizations or abstractions about the behaviors of other types.
By generalizing, interfaces let us write functions that are more flexible and adaptable
because they are not tied to the details of one particular implementation.

Many object-oriented lagnuages have some notion of interfaces, but what makes Go's interfaces
so distinctive is that they are SATISIFIED IMPLICITLY. In other words, there's no need to declare
all the interfaces that a given CONCRETE TYPE satisifies; simply possessing the necessary methods
is enough. This design lets you create new interfaces that are satisifed by existing CONCRETE TYPES
without changing the existing types, which is particularly useful for types defined in packages that
you don't control.

All the types we've looked at so far have been CONCRETE TYPES. A CONCRETE TYPE specifies the exact
representation of its values and exposes the intrinsic operations of that representation, such as
arithmetic for numbers, or indexing, append, and range for slices. A CONCRETE TYPE may also provide
additional behaviors through its methods. When you have a value of a CONCRETE TYPE, you know exactly
what is IS and what you can DO with it.

There is another kind of type in Go called an INTERFACE TYPE. An interface is an ABSTRACT TYPE. It doesn't
expose the representation or internal structure of its values, or the set of basic operations they support;
it reveals only some of their methods. When you have a value of an interface type, you know nothing about
what it IS; you know only what it can DO, or more precisely, what BEHAVIORS ARE PROVIDED BY ITS METHODS.

-------------------

type ReadWriter interface {
    Reader
    Writer
}

This is called EMBEDDING an interface.


-------------------

A type SATISFIES an interface if it possesses all the methods the interface requires.

-------------------

Conceptually, a value of an interface type, or INTERFACE VALUE, has two components,
    a CONCRETE TYPE and a
    VALUE OF THAT TYPE.
These are called the interface's
    DYNAMIC TYPE and
    DYNAMIC VALUE.

For a statically typed language like Go, types are a compile-time concept, so a type is not a value.
In our conceptual model, a set of values called TYPE DESCRIPTORS provide information about each type,
such as its name and methods. In an interface value, the type component is represented by the appropriate
type descriptor.


var w io.Writer
w = os.Stdout
w = new(bytes.Buffer)
w = nil


var w io.Writer
w
type: nil
value: nil

w = os.Stdout
w
type: *os.File
value: the address where a value of type os.File is stored

w = new(bytes.Buffer)
w
type: *bytes.Buffer
value: the address where a value of type bytes.Buffer is stored

w = nil
w
type: nil
value: nil

-------------------
The Go Programming Language
Donovan and Kernighan

Caplitalization and identation mine.
```

# AI 翻译

---

## Bill Kennedy 的解释

> **多态 (Polymorphism)** 是指：通过类型的实现，编写的代码能够表现出不同的行为。
> 一旦某个类型实现了一个接口，那么该类型的值就可以获得一整套新的功能。

---

> **接口 (Interface)** 是只声明行为的类型。
> 这些行为从不直接由接口本身实现，而是通过用户自定义类型的方法来实现。
> 当一个用户自定义类型实现了接口声明的所有方法时，这个自定义类型的值就可以赋给该接口类型的变量。
> 这种赋值操作会把该自定义类型的值存储到接口值中。

> 当你对接口值调用方法时，实际执行的就是存储在其中的具体类型值对应的方法。
> 因为任何自定义类型都可以实现任何接口，所以接口值上的方法调用天然就是**多态**的。
> 在这种关系中，自定义类型通常被称为**具体类型 (concrete type)**，因为接口值本身没有任何具体行为，必须依赖存储的具体类型。

---

### 接收者 (Receivers) 与可调用的值 (Values)

| 方法接收者    | 可以调用的值类型   |
|----------|------------|
| `(t T)`  | `T` 和 `*T` |
| `(t *T)` | `*T`       |

---

### 值 (Values) 与它们能调用的方法接收者 (Receivers)

| 值类型  | 能调用的方法接收者          |
|------|--------------------|
| `T`  | `(t T)`            |
| `*T` | `(t T)` 和 `(t *T)` |

---

📖 **来源**：《Go In Action》 William Kennedy

---

## Donovan & Kernighan 的解释

**接口类型 (Interface type)** 表达了其他类型的**一般化 (generalization)** 或**抽象 (abstraction)** 的行为。
通过一般化，接口让我们编写更灵活、更适应性的函数，而不必绑定在某一个特定实现的细节上。

许多面向对象语言都有接口的概念，但 **Go 的接口独特之处在于它们是隐式满足的 (satisfied implicitly)**。
换句话说，一个具体类型不需要显式声明它实现了哪些接口；只要它拥有接口要求的方法集合，就算实现了该接口。
这种设计让你可以基于现有的具体类型定义新的接口，而不需要修改这些已有类型——特别适用于你无法修改的外部包类型。

---

到目前为止我们看到的都是**具体类型 (concrete types)**。
具体类型指定了值的确切表现形式，并暴露了与之相关的基本操作（比如数字的算术运算，切片的索引、append、range）。
具体类型还可以通过方法提供额外的行为。
当你持有一个具体类型的值时，你清楚地知道它是什么以及能做什么。

Go 还有一种叫做 **接口类型 (interface type)** 的类型。
接口是一种**抽象类型 (abstract type)**。它既不暴露值的表现形式，也不暴露支持的基本操作；它只揭示了一些方法。
当你持有一个接口类型的值时，你不知道它是什么，只知道它能做什么——更确切地说，你知道它的方法所提供的行为。

---

### 接口嵌入 (Embedding)

```go
type ReadWriter interface {
Reader
Writer
}
```

这是 **接口嵌入**，可以通过组合已有接口定义新的接口。

---

### 满足接口 (Satisfying an interface)

只要某个类型实现了接口要求的所有方法，就**满足**该接口。

---

### 接口值的组成 (Interface Value)

一个接口值由两部分组成：

* 一个**具体类型 (concrete type)**
* 该类型的一个**值 (value)**

在接口的语境里，这两部分被称为：

* **动态类型 (dynamic type)**
* **动态值 (dynamic value)**

---

### 概念模型

在 Go 这样的静态语言中，类型是编译期概念，本身不是值。
可以把“类型描述符 (type descriptor)”看作是值的集合，用来表示类型信息，比如类型的名字和它的方法集合。
在接口值中，类型部分就是相应的类型描述符。

---

### 例子

```go
var w io.Writer

w = os.Stdout         // type: *os.File, value: 指向 os.File 的地址
w = new(bytes.Buffer) // type: *bytes.Buffer, value: 指向 Buffer 的地址
w = nil // type: nil, value: nil
```

* 当 `w = nil` → 接口值的 **类型 = nil, 值 = nil**
* 当 `w = os.Stdout` → **类型 = *os.File, 值 = os.File 的地址**
* 当 `w = new(bytes.Buffer)` → **类型 = *bytes.Buffer, 值 = Buffer 的地址**

---

📖 **来源**：《The Go Programming Language》 Alan Donovan & Brian Kernighan

---

✅ **总结一句**：

* **接口是行为的抽象**，本身不实现方法。
* **具体类型通过实现方法来满足接口**。
* **接口值包含动态类型和动态值**，方法调用是多态的。
* **Go 的接口是隐式实现**，更灵活。

---