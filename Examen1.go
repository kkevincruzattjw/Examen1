package main

import "fmt"

var productosVendidos []string
var subtotalesVentas []float64

func RegistrarVenta(nombre string, precio float64, cantidad int) {
	subtotal := precio * float64(cantidad)
	productosVendidos = append(productosVendidos, nombre)
	subtotalesVentas = append(subtotalesVentas, subtotal)
	fmt.Printf("¡Venta registrada con éxito! Producto: %s | Subtotal: $%.2f\n", nombre, subtotal)
}

func MostrarEstadisticas() {
	if len(productosVendidos) == 0 {
		fmt.Println("No hay ventas registradas.")
		return
	}
	
	total := 0.0
	for i := 0; i < len(productosVendidos); i++ {
		fmt.Printf("- %s: $%.2f\n", productosVendidos[i], subtotalesVentas[i])
		total += subtotalesVentas[i]
	}
	fmt.Printf("Total recaudado: $%.2f\n", total)
}

func main() {
	var opcion int

	for {
		fmt.Println("\n--- MENÚ DE VENTAS ---")
		fmt.Println("1. Registrar una nueva venta")
		fmt.Println("2. Mostrar estadísticas")
		fmt.Println("3. Salir")
		fmt.Print("Seleccione una opción: ")
		fmt.Scan(&opcion)

		if opcion == 1 {
			fmt.Println("\n--- LISTA DE PRODUCTOS ---")
			fmt.Println("1. Arroz - $1.25")
			fmt.Println("2. Leche - $0.95")
			fmt.Println("3. Pan - $0.50")
			fmt.Print("Seleccione el número del producto: ")
			
			var prod int
			fmt.Scan(&prod)

			fmt.Print("Ingrese la cantidad vendida: ")
			var cant int
			fmt.Scan(&cant)

			switch prod {
			case 1:
				RegistrarVenta("Arroz", 1.25, cant)
			case 2:
				RegistrarVenta("Leche", 0.95, cant)
			case 3:
				RegistrarVenta("Pan", 0.50, cant)
			default:
				fmt.Println("Producto no válido.")
			}
		} else if opcion == 2 {
			MostrarEstadisticas()
		} else if opcion == 3 {
			fmt.Println("Saliendo del programa...")
			break
		} else {
			fmt.Println("Opción inválida.")
		}
	}
}






