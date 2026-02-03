package main

import (
	"fmt"
	"math/rand"
	"time"
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

func main() {
	fmt.Printf(" Bienvenue sur DEXX-Quest ! \n ------------------------ \n")
	fmt.Printf("\n Entrez le nom de votre personnage : ")
	var input string
	fmt.Scanln(&input)
	fmt.Printf("\n Bonjour %s voici vos statistiques : \n\n", input)

	p1 := Player{input, 100, 50}
	m1 := Enemy_nul{"Code-Crawler", 50, 10}
	b1 := Boss{"Kernel-Panic", 100, 30}

	statistiques(p1)

	var jeu string
	fmt.Printf("Maintenant que tout est prêt voulez vous entrez dans le jeu ? (oui / non)\n\n")
	fmt.Scanln(&jeu)
	if jeu == "oui" {
		clearscreen()
		for {
			var choix string
			fmt.Printf(`Le jeu commence ...\n`)
			fmt.Printf(`Vos yeux s'ouvre... vous êtes dans une fôret que faire ? 
			1. Explorer la fôret 
			2. S'allonger (quitter le jeu)
			3. Voir mes statistiques\n\n
			`)
			fmt.Scanln(&choix)
			time.Sleep(2000 * time.Millisecond)
			clearscreen()
			switch choix {
			case "1":
				clearscreen()
				fmt.Printf(`
					L'air vibre d'une électricité statique inhabituelle. Les arbres de cette forêt ne sont
					pas faits de bois, mais de câbles tressés et de fibres optiques lumineuses. Soudain,
					le sol tremble... Une distorsion visuelle apparaît devant toi : un Code-Crawler, une créature
					faite de pixels instables et de segments de mémoire corrompue, bloque le chemin. Il semble vouloir 
					absorber tes données vitales
					`)
				fmt.Printf("----------------------------------------------------\n\n")
				fmt.Printf(`
					! COMBAT ! 
					`)
				fmt.Printf(`
					1. ATTAQUER 
					2. FUIR
					3. VOIR STAT
					`)
				for p1.hp > 0 && m1.hp > 0 {
					var combat1 string
					fmt.Scanln(&combat1)
					if combat1 == "1" {
						clearscreen()
						random_attack := rand.Intn(50)
						time.Sleep(2000 * time.Millisecond)
						fmt.Printf("\n\ncoup de spas12 dans sa mère : -%d hp pour le Code-Crawler\n\n", random_attack)
						m1.hp = m1.hp - random_attack
						time.Sleep(1000 * time.Millisecond)
						if m1.hp <= 0 {
							fmt.Printf("\n\nle Code-Crawler est mort BRAVO !\n\n")
							break
						} else {
							fmt.Printf("\n\nATTENTION, le Code-Crawler attaque !\n\n")
							time.Sleep(2000 * time.Millisecond)
							random_attack_m1 := rand.Intn(30)
							fmt.Printf("\n\nAïe ! - %d hp\n\n", random_attack_m1)
							fmt.Printf(`
								1. ATTAQUER 
								2. FUIR
								3. VOIR STAT
							`)

							p1.hp = p1.hp - random_attack_m1

							if p1.hp <= 0 {
								fmt.Printf(`Vous êtes mort...`)
								return
							}
						}
					} else if combat1 == "2" {
						fmt.Printf("Tu à opté pour la fuite, sale lâche")
						time.Sleep(2000 * time.Millisecond)
						the_end()
						return
					} else if combat1 == "3" {
						statistiques(p1)
						fmt.Printf(`
								1. ATTAQUER 
								2. FUIR
								3. VOIR STAT
							`)
						continue
					}
				}
				fmt.Printf(`
						Après avoir réduit le Code-Crawler en poussière binaire, %s s'enfonce plus profondément dans la forêt. Soudain, les arbres s'effacent pour laisser
						place à une forteresse monolithique dont les tours semblent toucher le ciel de données.

						En explorant les couloirs froids du château, tu tombes sur une porte cyclopéenne noyée dans un brouillard épais et glacial. Animé par un courage 
						hors norme, tu franchis le seuil. Là, trônant au milieu d'une salle dévastée, se dresse l'entité suprême : le Kernel-Panic. L'air se fige,
						 la réalité scintille... Que vas-tu faire ?
						`)
				fmt.Printf(`
						1. Lever la tête et affronter le Kernel-Panic (COMBAT)
						2. Baissez la tête et s'agenouiller devant lui(FUITE)
						3. VOIR STAT
						`)

				for {
					var choix_final string
					fmt.Scanln(&choix_final)

					switch choix_final {
					case "1":
						clearscreen()
						fmt.Printf("Bizarrement, le Kernel-Panic semble être un humain mais corrompu par un virus puissant... Mais rien à faire le combat va avoir lieu !\n")
						time.Sleep(2000 * time.Millisecond)
						attack_random := rand.Intn(50)
						fmt.Printf("\n\ncoup de spas12 dans sa mère : -%d hp pour le Kernel-Boss\n\n", attack_random)
						b1.hp = b1.hp - attack_random
						if b1.hp <= 0 {
							fmt.Printf(`
										Félicitations ! le boss est vaincu, vous avancez et récuprez son épée, vous deven....
										système corrompu.........
										
										quoi ? quelque chose est étrange, regardez vos statistiques...
									`)
							time.Sleep(2000 * time.Millisecond)
							fmt.Printf(`
										REGARDEZ VOS STATISTIQUES TIQUES
									`)
							time.Sleep(2000 * time.Millisecond)
							fmt.Printf(".")
							time.Sleep(1000 * time.Millisecond)
							fmt.Printf(".")
							time.Sleep(1000 * time.Millisecond)
							fmt.Printf(".\n")
							time.Sleep(1000 * time.Millisecond)
							p1.nom = "Kernel-Boss"
							statistiques(p1)
							time.Sleep(2000 * time.Millisecond)
							fmt.Printf(`
										Votre personnage est devenu le Kernel-Boss, vous avez perdu le jeu !
									`)
							time.Sleep(2000 * time.Millisecond)
							clearscreen()
							the_end()
							time.Sleep(4000 * time.Millisecond)
							return
						} else {
							fmt.Printf("\n\nLe Kernel-Boss a encore %d hp\n\n", b1.hp)
						}
						fmt.Printf("\n\nATTENTION, le Kernel-Boss attaque !\n\n")
						time.Sleep(3000 * time.Millisecond)
						attack_random_b1 := rand.Intn(30)
						p1.hp = p1.hp - attack_random_b1
						fmt.Printf("Kernel-Boss damage : %d", attack_random_b1)
						if p1.hp <= 0 {
							fmt.Printf("VOUS ÊTES MORT...")
							return
						}
						fmt.Printf(`
									1. ATTAQUER 
									2. FUIR
									3. VOIR STAT
								`)

					case "2":
						fmt.Printf("Tu à opté pour la fuite, sale lâche")
						time.Sleep(2000 * time.Millisecond)
						clearscreen()
						the_end()
						return

					case "3":
						statistiques(p1)
						continue

					default:
						fmt.Printf("Veuillez entrez un choix valide !")
						continue
					}

				}

			case "2":
				clearscreen()
				fmt.Printf(`Vous vous laissez tombez en arrière et fermé les yeux...`)
				time.Sleep(2000 * time.Millisecond)
				clearscreen()
				fmt.Printf(`...`)
				the_end()
				return

			case "3":
				clearscreen()
				fmt.Printf(`Voici vos statistiques : 
				`)
				statistiques(p1)
			default:
				fmt.Printf("Veuillez entrez un choix valide !")
				continue
			}
		}
	} else {
		fmt.Printf("FIN DU JEU !")

	}

}
