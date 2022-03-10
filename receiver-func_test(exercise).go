//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main

import "testing"

func newPlayer() Player {
	return Player {
		name: 		"Wizard",
		health: 	100,
		maxHealth: 	100,
		energy: 	250,
		maxEnergy: 	250,
	}
}

func TestHealth(test *testing.T) {
	player := newPlayer()
	player.addHealth(500)
	if player.health > player.maxHealth {
		test.Fatalf("Health exceeded limit: %v, want: %v", player.health, player.maxHealth)  //fatalf statment completely ignores everything after the fatality of the code
	}
	player.applyDamage(player.maxHealth + 2) //maximum health plus 2 damage points
	if player.health < 0 {
		test.Fatalf("Health: %v. Minimum: 0", player.health)
	}
	if player.health > player.maxHealth {
		test.Fatalf("Health: %v. Maximum: %v", player.health, player.maxHealth)
	}
}

func TestEnergy(test *testing.T) {
	player := newPlayer()
	player.addEnergy(350)
	if player.energy > player.maxEnergy {
		test.Fatalf("Energy exceeded limit: %v, want: %v", player.energy, player.maxEnergy)  //fatalf statment completely ignores everything after the fatality of the code
	}
	player.consumeEnergy(player.maxEnergy + 2) //maximum energy plus 2 energy consuming points
	if player.energy < 0 {
		test.Fatalf("Health: %v. Minimum: 0", player.energy)
	}
	if player.energy > player.maxEnergy {
		test.Fatalf("Energy: %v. Maximum: %v", player.energy, player.maxEnergy)
	}
}