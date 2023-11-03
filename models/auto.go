package models

import (
	"fmt"
	"math/rand"
//	"sync"
	"time"
)

type Auto struct {
	ID int
}

func (a Auto) Estacionar(e *Estacionamiento, indice int) {
	fmt.Printf("Auto %d se está estacionando en el cajón %d.\n", a.ID, indice)
	time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second) // Simular tiempo de estacionamiento
	fmt.Printf("Auto %d ha salido del cajón %d.\n", a.ID, indice)
	e.LiberarCajon(indice)
}