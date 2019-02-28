package main

import "container/ring"

func main() {

}

func circuloDosDesempregados(qtPessoas int) {
	// Create a new ring of size 10
	r := ring.New(qtPessoas)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// K POSIÇÕES QUE O OFICIAL 1 VAI ANDAR E RETIRAR

	// M POSIÇÕES QUE O OFICIAL 2 VAI ANDAR E RETIRAR

}
