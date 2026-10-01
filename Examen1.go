package main

import "fmt"

var productosVendidos []string
var subtotalesVentas []float64

func RegistrarVenta(nombre string, precio float64, cantidad int) {
	subtotal := precio * float64(cantidad)
	productosVendidos = append(productosVendidos, nombre)
	subtotalesVentas = append(subtotalesVentas, subtotal)
}

func MostrarEstadisticas() {
	if len(productosVendidos) == 0 {
		fmt.Println("No hay ventas registradas.")
		return
	}
	
	total := 0.0
	for i := 0; i < len(productosVendidos); i++ {
		total += subtotalesVentas[i]
	}
	fmt.Printf("Total recaudado: $%.2f\n", total)
}

func main() {
	RegistrarVenta("Arroz", 1.25, 2)
	MostrarEstadisticas()
}












