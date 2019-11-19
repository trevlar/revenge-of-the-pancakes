package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestFlipPancake(t *testing.T) {

	// Case #1: 1
	results := FlipPancake(0, "-")
	if results != "+" {
		t.Errorf("Pancake flipped on wrong side, got %s, want: %s.", results, "+")
	}
	// Case #2: 1
	results = FlipPancake(0, "-+")
	if results != "++" {
		t.Errorf("Pancake flipped on wrong side, got %s, want: %s.", results, "++")
	}
	// Case #3: 2
	results = FlipPancake(0, "+-")
	if results != "--" {
		t.Errorf("Pancake flipped on wrong side, got %s, want: %s.", results, "--")
	}
	// Case #4: 0
	results = FlipPancake(0, "+++")
	if results != "-++" {
		t.Errorf("Pancake flipped on wrong side, got %s, want: %s.", results, "-++")
	}
	// Case #5: 3
	results = FlipPancake(1, "--+-")
	if results != "+++-" {
		t.Errorf("Pancake flipped on wrong side, got %s, want: %s.", results, "+++-")
	}
}

func TestFindFromTop(t *testing.T) {

	// Case #1: 1
	results := FindFromTop("-")
	if results != 0 {
		t.Errorf("Found wrong pancake from stack to flip, got %d, want: %d.", results, 0)
	}
	// Case #2: 1
	results = FindFromTop("-+")
	if results != 0 {
		t.Errorf("Found wrong pancake from stack to flip, got %d, want: %d.", results, 0)
	}
	// Case #3: 2
	results = FindFromTop("+-")
	if results != 0 {
		t.Errorf("Found wrong pancake from stack to flip, got %d, want: %d.", results, 0)
	}
	// Case #4: 0
	results = FindFromTop("+++")
	if results != -1 {
		t.Errorf("Found wrong pancake from stack to flip, got %d, want: %d.", results, 0)
	}
	// Case #5: 3
	results = FindFromTop("--+-")
	if results != 1 {
		t.Errorf("Found wrong pancake from stack to flip, got %d, want: %d.", results, 0)
	}
}

func TestFindAndTop(t *testing.T) {

	// Case #1: 1
	resultStack, numberOfFlips := FindAndFlip("-")
	if resultStack != "-" && numberOfFlips != 1 {
		t.Errorf("Wrong results for flipped pancake stack, got %s %d flips, want: %s %d flips.", resultStack, numberOfFlips, "+", 1)
	}
	// Case #2: 1
	resultStack, numberOfFlips = FindAndFlip("+-")
	if resultStack != "++" && numberOfFlips != 1 {
		t.Errorf("Wrong results for flipped pancake stack, got %s %d flips, want: %s %d flips.", resultStack, numberOfFlips, "++", 1)
	}
	// Case #3: 2
	resultStack, numberOfFlips = FindAndFlip("-+")
	if resultStack != "++" && numberOfFlips != 1 {
		t.Errorf("Wrong results for flipped pancake stack, got %s %d flips, want: %s %d flips.", resultStack, numberOfFlips, "++", 2)
	}
	// Case #4: 0
	resultStack, numberOfFlips = FindAndFlip("+++")
	if resultStack != "+++" && numberOfFlips != 0 {
		t.Errorf("Wrong results for flipped pancake stack, got %s %d flips, want: %s %d flips.", resultStack, numberOfFlips, "+++", 0)
	}
	// Case #5: 3
	resultStack, numberOfFlips = FindAndFlip("--+-")
	if resultStack != "++++" && numberOfFlips != 1 {
		t.Errorf("Wrong results for flipped pancake stack, got %s %d flips, want: %s %d flips.", resultStack, numberOfFlips, "++++", 3)
	}
}