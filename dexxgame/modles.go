package main

import (
	"fmt"
)

type Player struct {
	nom    string
	hp     int
	attack int
}

type Enemy_nul struct {
	nom    string
	hp     int
	attack int
}

type Boss struct {
	nom    string
	hp     int
	attack int
}

func statistiques(a Player) {
	fmt.Printf("[ NAME : %s | HP : %d | ATTACK DMG : %d]\n\n", a.nom, a.hp, a.attack)
}

func clearscreen() {
	fmt.Print("\033[H\033[2J")
}

func the_end() {
	fmt.Printf(`
	  _______ _    _ ______   ______ _   _ _____  
 |__   __| |  | |  ____| |  ____| \ | |  __ \ 
    | |  | |__| | |__    | |__  |  \| | |  | |
    | |  |  __  |  __|   |  __| | . | |  | |
    | |  | |  | | |____  | |____| |\  | |__| |
    |_|  |_|  |_|______| |______|_| \_|_____/ 
                                              `)
}
