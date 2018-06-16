package main

import (
	"fmt"
)

type krieger struct {
	kraft      int
	gesundheit int
}

func (k *krieger) GetKraft() int { return k.kraft }

func (k *krieger) GetGesundheit() int { return k.gesundheit }

func (k *krieger) Schaden(i int) {
	k.gesundheit = k.gesundheit - i
}

func (k *krieger) Angriff() {
	k.kraft--
}

func (k *krieger) Schlaf() {
	k.gesundheit = k.gesundheit + 5
	k.kraft = k.kraft + 5
}

func (k *krieger) String() string {
	return fmt.Sprintf("Kraft: %d Gesundheit: %d", k.kraft, k.gesundheit)
}

func main() {
	k := &krieger{kraft: 100, gesundheit: 100}
	fmt.Println(k.GetKraft(), k.GetGesundheit())
	k.Schaden(10)
	k.Angriff()
	fmt.Println(k)
	k.Schlaf()
	fmt.Println(k)
}
