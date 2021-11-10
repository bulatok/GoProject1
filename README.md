# ScriptOnGolang

---------------------------

При установке пути для переменной "ParentPath", выводит дерево каталогов это пути (используется рекусивный алгоритм)<p></p>
(в примере используется каталог ./testdata)
```bash 
$ go run main.go . -f
``` 
```
├───go.mod (26b)
├───main.go (1966b)
├───main_test.go (1376b)
└───testdata
    ├───project
        ├───file.txt (19b)
        └───gopher.png (70372b)
    ├───static
        ├───a_lorem
            ├───dolor.txt (empty)
            ├───gopher.png (70372b)
            └───ipsum
                └───gopher.png (70372b)
        ├───css
            └───body.css (28b)
        ├───empty.txt (empty)
        ├───html
            └───index.html (57b)
        ├───js
            └───site.js (10b)
        └───z_lorem
            ├───dolor.txt (empty)
            ├───gopher.png (70372b)
            └───ipsum
                └───gopher.png (70372b)
    ├───zline
        ├───empty.txt (empty)
        └───lorem
            ├───dolor.txt (empty)
            ├───gopher.png (70372b)
            └───ipsum
                └───gopher.png (70372b)
    └───zzfile.txt (empty)
   ```
или
```bash 
$ go run main.go . 
``` 
```
├───project
├───static
    ├───a_lorem
        └───ipsum
    ├───css
    ├───html
    ├───js
    └───z_lorem
        └───ipsum
├───zline
    └───lorem
        └───ipsum
```