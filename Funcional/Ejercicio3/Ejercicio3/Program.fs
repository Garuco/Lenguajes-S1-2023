// For more information see https://aka.ms/fsharp-console-apps
type Producto = {
    Nombre: string
    Descripcion: string
    MontoVenta: int
}

let rangoPagoImpuestos = 20000
let porcentajeImpuesto = 0.13

type Productos = Producto list

let calcularImpuestoFactura (f: Productos) =
    let lista = f |> List.filter (fun p -> p.MontoVenta > rangoPagoImpuestos)
    let lista2 = lista |> List.map (fun p -> int(float p.MontoVenta * porcentajeImpuesto))
    lista2 |> List.reduce (+)

let calcularMontoFactura (f: Productos) =
    let lista = f |> List.map (fun p -> p.MontoVenta)
    lista |> List.reduce (+)


let factura = [
    { Nombre = "tarjeta madre"; Descripcion = "Asus"; MontoVenta = 54200 }
    { Nombre = "mouse"; Descripcion = "alámbrico"; MontoVenta = 15000 }
    { Nombre = "teclado"; Descripcion = "gamer con luces"; MontoVenta = 30000 }
    { Nombre = "memoria ssd"; Descripcion = "2 gb"; MontoVenta = 65000 }
    { Nombre = "cable video"; Descripcion = "display port 1m"; MontoVenta = 18000 }
]

printfn "Monto total: %d" (calcularMontoFactura factura)
printfn "Impuesto total: %d" (calcularImpuestoFactura factura)