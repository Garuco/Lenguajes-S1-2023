// For more information see https://aka.ms/fsharp-console-apps
(*Generar una lista con una secuencia de números desde un punto A hasta un punto B,
Caso base: cuando A sobrepase a B, devolvemos una lista vacía, sino
agregamos A a la lista con el operador de concatenación,
y llamamos recursivamente la función pero incrementando A en 1 para formar la secuencia*)
let rec generar_sec A B =
    if A > B
    then []
    else A :: generar_sec (A+1) B
    
let resultado = generar_sec 0 7
printfn "%A" resultado
