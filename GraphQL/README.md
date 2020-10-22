## query の書き方

参考(

https://qiita.com/mitubaEX328/items/77ccc4f6ac0ad2e76996
http://snowtooth.moonhighway.com/
https://github.com/mitubaEX/graphQL_sample

)



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
