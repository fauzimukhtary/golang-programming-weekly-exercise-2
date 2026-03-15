package main

import "fmt"

const NMAX int = 999

type character struct {
	name         string
	hp, nAttacks int
	attacks      [NMAX]int
	totalAttacks int
}

func main() {
	var link, ganon int
	fmt.Scan(&link, &ganon)
	gaming(link, ganon)
}

func attack(damage int, attack, hp, nAttack *int) {
	*hp = *hp - damage
	*attack = *attack + damage
	*nAttack = *nAttack + 1
}

func defense(damage int, hp *int) {
	*hp = *hp + damage
}

func parry(damage int, selectorhp, victimhp *int) {
	*victimhp = *victimhp - damage
	*selectorhp = *selectorhp + damage
}

func winner(p1, p2 character) {
	if p2.hp < 0 {
		p2.hp = 0
	}

	if p1.hp < 0 {
		p1.hp = 0
	}

	fmt.Println()
	if p2.hp == 0 {
		fmt.Println("Link menang! putri Zelda berhasil diselamatkan.")
	} else {
		fmt.Println("Ganon menang! putri Zelda dalam bahaya.")
	}

	fmt.Printf("%v: health %d, attacks %d, damage dealt %d\n", p1.name, p1.hp, p1.nAttacks, p1.totalAttacks)
	fmt.Printf("%v: health %d, attacks %d, damage dealt %d\n", p2.name, p2.hp, p2.nAttacks, p2.totalAttacks)
	fmt.Println()
}

func gaming(linkHP, ganonHP int) {
	var lin, gan character
	var player, action string
	var selector, victim *character
	var i int

	lin.name, gan.name = "Link", "Ganon"
	lin.hp, gan.hp = linkHP, ganonHP
	lin.nAttacks, gan.nAttacks = 0, 0
	lin.totalAttacks, gan.totalAttacks = 0, 0
	i = 0

	for lin.hp > 0 && gan.hp > 0 {
		fmt.Scan(&player, &action)
		switch player {
		case "Link":
			selector = &lin
			victim = &gan
		case "Ganon":
			selector = &gan
			victim = &lin
		}

		switch action {
		case "ATTACK":
			fmt.Scan(&selector.attacks[i])
			attack(selector.attacks[i], &selector.totalAttacks, &victim.hp, &selector.nAttacks)
		case "DEFENSE":
			if i > 0 {
				defense(victim.attacks[i-1], &selector.hp)
			}
		case "PARRY":
			if i > 0 {
				parry(victim.attacks[i-1], &selector.hp, &victim.hp)
			}
		}
		i = i + 1
	}
	winner(lin, gan)
}
