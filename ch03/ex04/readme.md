# 1.7節のリサージュの例での方法にしたがって、面を計算してSVGデータをクライントに書き出すウェブサーバーを作成する.  
サーバーは次のようにcontent-typeを指定する必要がある  
```
w.Header().Set("Conten-Type"."image/svg+xml")
```