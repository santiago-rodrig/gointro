package main

func main() {

}

func findGreatest(ns ...float64) float64 {
	greatest := ns[0]
	for _, n := range ns[1:] {
		if n > greatest {
			greatest = n
		}
	}
	return greatest
}
