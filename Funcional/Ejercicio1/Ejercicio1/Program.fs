// For more information see https://aka.ms/fsharp-console-apps

open System.Text.RegularExpressions

(*Se utiliza el patrón de expresiones regulares "\\b" + palabra + "\\b",
de esta manera se buscaría la palabra como independiente, es decir que no tenga ningún otro caracter
alfanúmerico en sus extremos.*)
let sub_cadenas (palabra:string) (lista:string list) =
    let pattern = "\\b" + palabra + "\\b"
    let regex = new Regex(pattern)
    lista
    |> List.filter (fun frase -> regex.IsMatch(frase))
    
let frases = ["helado de chocolate"; "dedo"; "es de él"; "destino"; "dioxido de carbono"]
printfn "%A" (sub_cadenas "de" frases)
