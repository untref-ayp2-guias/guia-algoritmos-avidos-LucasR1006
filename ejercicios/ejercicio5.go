package ejercicios

var billetes = []int{10000, 2000, 1000, 500, 200, 100, 50, 20, 10, 5, 2, 1}

func Cambiar(cantidad int) map[int]int {

	cambio := make(map[int]int)
	if cantidad <= 0 {
		return nil
	}

	for _, value := range billetes {
		if cantidad >= value {
			cambio[value] = cantidad / value
			cantidad = cantidad % value
		}
	}
	Aux(cantidad, cambio, 0)

	return cambio
}

func Aux(cantidad int, m map[int]int, index int) {
	if cantidad == 0 {
		return
	}
	if cantidad >= billetes[index] {
		m[billetes[index]] = cantidad / billetes[index]
		cantidad = cantidad % billetes[index]
	}
	Aux(cantidad, m, index+1)
}
