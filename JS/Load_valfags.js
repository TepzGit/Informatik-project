const Alle_Valgfag = ["Fest", "Film", "Lan", "Musik", "Science", "Sport"]

const valgfags = document.getElementsByClassName("Valgfags");

for (let i=0;i<Alle_Valgfag.length;i++) {
	const valfag = document.createElement("div")
	valfag.className = "valgfag"
	valgfags[0].appendChild(valfag)

	const valgfag_text = document.createElement("p")
	valgfag_text.style = "color: #ffffff; font-size: 50px"
	valgfag_text.textContent = Alle_Valgfag[i]

	valfag.appendChild(valgfag_text)
}
