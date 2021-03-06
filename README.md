Go言語のポインタ（値参照）について
=====

このコードはGo言語における引数の扱い方を表しています。

次のような関数が定義されています。

inc1 ・・・ 値を受け取り、インクリメントします。  
inc2 ・・・ 値をポインタの値渡しで受け取り、インクリメントします。  

次のようにnum変数を定義します。  
`num := 10`

結果は次のようになることが確認されています。  
>inc1: i =  11  
>i =  10  
>inc2: i =  11  
>i =　11  

この結果から、go言語は値渡しであることが分かります。  
参照先の値を変える場合は、呼び出し元では「&」を使って変数のアドレスを渡し、  
呼び出し先では「*」を使ってポインタの参照している先の値を使うことができます。  

しかし、これは参照渡しではなく、  
ポインタをコピーして渡しているので「ポインタの値渡し」と言えます。  
