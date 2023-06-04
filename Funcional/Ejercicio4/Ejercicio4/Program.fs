// Ejercicios de rutas en Grafos con búsqueda en profundidad

//        a ---- c ---- x
//      /   \  /
//     /     \/
//   i       X
//     \    / \
//      \  /   \
//        b ---- d ---- f
let grafo = [
    [('i', 0);('a', 1);('b', 2)];
    [('a', 1);('i', 0);('c', 3);('d', 2)];
    [('b', 2);('i', 0);('c', 1);('d', 4)];
    [('c', 3);('a', 1);('b', 1);('x', 2)];
    [('d', 2);('a', 4);('b', 1);('f', 3)];
    [('f', 3);('d', 2)];
    [('x', 2);('c', 3)]
]

let miembro e (lista: ('a * 'b) list) =
    lista
    |> List.map (fun (x,_) -> x = e)
    |> List.reduce (fun x y -> x || y)

let rec vecinos nodo grafo =
    match grafo with
    | [] -> []
    | ((head, _)::adyacencias)::_ when nodo = head -> adyacencias
    | _::tail -> vecinos nodo tail 
(*let rec vecinos nodo grafo =
    match grafo with
    | [] -> []
    | ((head::_)::(tail::_))::_ when nodo = head -> tail
    | _::tail -> vecinos nodo tail*) 

let extender (ruta: _ list) grafo =
    (vecinos (fst ruta.Head) grafo)
    |> List.map (fun (x,peso) -> if (miembro x ruta) then [] else (x,peso)::ruta )
    |> List.filter (fun x -> x <> [])

let rec prof_aux fin (rutas: _ list list) grafo =
    if rutas = [] then
        []
    elif fin = (fst rutas.Head.Head) then
        //List.rev rutas.Head
        List.append
            [List.rev rutas.Head]
            (prof_aux fin rutas.Tail grafo)
    else
        prof_aux fin (List.append (extender rutas.Head grafo) rutas.Tail) grafo
        
let prof ini fin grafo =
    prof_aux fin [[(ini,0)]] grafo
    

printf "%A" (prof 'i' 'f' grafo)

//Ejercicio 4

let rec rutaCorta_aux fin (rutas: (char * int) list list) grafo =
    if rutas = [] then
        []
    elif fin = (fst rutas.Head.Head) then
        let rutaActual = List.rev rutas.Head
        let pesoRutaActual = rutaActual |> List.map snd |> List.sum
        let rutasRestantes = rutaCorta_aux fin rutas.Tail grafo
        let rutasConPesoTotal = (pesoRutaActual, rutaActual)::rutasRestantes
        rutasConPesoTotal
    else
        rutaCorta_aux fin (List.append (extender rutas.Head grafo) rutas.Tail) grafo

let rutaCorta ini fin grafo =
    rutaCorta_aux fin [[(ini,0)]] grafo
    |> List.sortBy (fun (peso,_) -> peso)
    |> List.head
    |> snd

printf "\n\nLa ruta mas corta es:\n"
printf "%A" (rutaCorta 'i' 'f' grafo)