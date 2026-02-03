package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("DEXX-Quest")

	titre := widget.NewLabel("BIENVENUE DANS DEXX-Quest")

	saisieNom := widget.NewEntry()
	saisieNom.SetPlaceHolder("Veuillez entrez votre nom...")

	boutonValider := widget.NewButton("...", func() {

		nomJoueur := saisieNom.Text
		p1 := Player{nomJoueur, 100, 50}

		titre.SetText("Bonjour " + p1.nom + "... connexion établie")
	})

	window.SetContent(container.NewVBox(
		titre,
		saisieNom,
		boutonValider,
	))

	window.ShowAndRun()

}
