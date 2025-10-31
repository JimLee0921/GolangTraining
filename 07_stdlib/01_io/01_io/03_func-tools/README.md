# io 包常用工具函数

## 数据复制相关

### **`io.Copy(dst Writer, src Reader)`**

从 `src` 不断读，写入 `dst`，直到 `EOF`。

```
io.Copy(os.Stdout, strings.NewReader("hello"))
```

用途：

* 文件复制
* 网络转发
* 标准输入输出转管道

---

### **`io.CopyN(dst, src, n)`**

复制 `n` 字节。

```
io.CopyN(os.Stdout, os.Stdin, 5)
```

---

### **`io.ReadAll(r)`**

一次读完所有内容，返回 `[]byte`。

> 用于内存内容（如小文件、HTTP body），不用来读大文件。

```
b, _ := io.ReadAll(r)
```

---

### **`io.ReadFull(r, buf)`**

把 `buf` 填满，否则报错。
多用于协议读取。

```
buf := make([]byte, 10)
io.ReadFull(r, buf)
```

---

## 多路输入/输出组合

### **`io.MultiReader(r1, r2, r3...)`**

多个 `Reader` 串成一个。

```
r := io.MultiReader(strings.NewReader("A"), strings.NewReader("B"))
io.Copy(os.Stdout, r) // 输出 AB
```

用途：

* 拼接文件头+文件体
* 拼接多块数据成为一个流

---

### **`io.MultiWriter(w1, w2, w3...)`**

写一次 -> 多个地方同时写。

```
w := io.MultiWriter(os.Stdout, logfile)
fmt.Fprintln(w, "Hello")
```

用途：

* 同时写终端 + 日志文件（最常用！）

---

## 流拦截与旁路操作

### **`io.TeeReader(r, w)`**

读 `r` 时，会顺便写一份到 `w`，但输出流仍然来自 `r`。

```
r := io.TeeReader(resp.Body, logfile)
io.ReadAll(r) // 一边读，一边写到日志
```

用途：

* 抓包 / 调试网络响应
* 打印 HTTP Body

---

### **`io.LimitReader(r, n)`**

只允许读前 `n` 字节。

```
limited := io.LimitReader(r, 100)
```

用途：

* 避免恶意超大输入攻击（协议安全用途）
