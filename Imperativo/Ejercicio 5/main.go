package main

import (
	"fmt"
	"sort"
	"strings"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	pos := l.buscarProducto(nombre)

	if pos != -1 {
		(*l)[pos].cantidad += cantidad
		(*l)[pos].precio = precio
	} else {
		*l = append(lProductos, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	}

	//*l = append(lProductos, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	// modificar el código para que cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad
	// de elementos del producto y eventualmente el precio si es que es diferente
}

func (l *listaProductos) buscarProducto(nombre string) int { //el retorno es el índice del producto encontrado y -1 si no existe
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
		}
	}
	return result
}

func (l *listaProductos) venderProducto(nombre string, cantidad int) {
	var prod = l.buscarProducto(nombre)
	if prod != -1 && cantidad < (*l)[prod].cantidad {
		(*l)[prod].cantidad = (*l)[prod].cantidad - cantidad
		//modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista" y notificar
	} else if cantidad == (*l)[prod].cantidad {
		*l = append((*l)[:prod], (*l)[prod+1:]...)
	} else {
		fmt.Println("La cantidad solicitada sobrepasa la disponible.")
	}
}
func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)
	lProductos.agregarProducto("huevos", 8, 2500)
	lProductos.agregarProducto("pasta", 17, 500)
	lProductos.agregarProducto("queso", 3, 2500)
	lProductos.agregarProducto("aceite", 11, 1200)
}

func (l *listaProductos) listarProductosMínimos() listaProductos {
	// debe retornar una nueva lista con productos con existencia mínima
	var l2 []producto
	for _, product := range *l {
		if product.cantidad < existenciaMinima {
			l2 = append(l2, product)
		}
	}
	return l2
}

func (l *listaProductos) aumentarInventarioDeMinimos(listaMinimos []producto) {
	for _, product := range listaMinimos {
		for i := 0; i < len(*l); i++ {
			if (*l)[i].nombre == product.nombre {
				(*l)[i].cantidad = existenciaMinima
			}
		}
	}
}

func (l *listaProductos) ordenarPorLlave(key string) {
	if strings.ToLower(key) == "nombre" {
		sort.Slice(lProductos, func(i, j int) bool {
			return strings.ToLower(lProductos[i].nombre) < strings.ToLower(lProductos[j].nombre)
		})
	} else if strings.ToLower(key) == "precio" {
		sort.Slice(lProductos, func(i, j int) bool {
			return lProductos[i].precio < lProductos[j].precio
		})
	} else if strings.ToLower(key) == "cantidad" {
		sort.Slice(lProductos, func(i, j int) bool {
			return lProductos[i].cantidad < lProductos[j].cantidad
		})
	} else {
		fmt.Println("Atributo invalido.")
	}
}

func main() {
	llenarDatos()
	lProductos.agregarProducto("atun", 7, 1500)
	lProductos.venderProducto("leche", 2)
	lProductos.venderProducto("arroz", 15)
	lProductos2 := lProductos.listarProductosMínimos()
	fmt.Println(lProductos)
	fmt.Println(lProductos2)

	lProductos.ordenarPorLlave("cantidad")
	fmt.Println("Ordenada por cantidad:", lProductos)

	lProductos.aumentarInventarioDeMinimos(lProductos2)
	fmt.Println("Aumentar inventario de minimos:", lProductos)

	lProductos.ordenarPorLlave("precio")
	fmt.Println("Ordenada por precio:", lProductos)

	lProductos.ordenarPorLlave("nombre")
	fmt.Println("Ordenada por nombre:", lProductos)

	lProductos.ordenarPorLlave("cantidad")
	fmt.Println("Ordenada por cantidad:", lProductos)

}
