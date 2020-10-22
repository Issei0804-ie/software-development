## query の書き方

```
query {
  User{
    name
    status
  }
}
```

名前の指定の掛け方

```
query {
  User(name:"hoge"){
    status
  }
}
```

フィルタの掛け方

```
query {
  User(status:OFFLINE){
    name
  }
}
```
エイリアスの付け方

```
query {
  fuga: User(status:OFFLINE){
   hoge:  name
  }
}
```
