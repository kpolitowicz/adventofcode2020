package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	got := ParseInput(`ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`)

	want := []Passport{
		{PassportFields{"ecl": "gry", "pid": "860033327", "eyr": "2020", "hcl": "#fffffd", "byr": "1937", "iyr": "2017", "cid": "147", "hgt": "183cm"}},
		{PassportFields{"iyr": "2013", "ecl": "amb", "cid": "350", "eyr": "2023", "pid": "028048884", "hcl": "#cfa07d", "byr": "1929"}},
		{PassportFields{"hcl": "#ae17e1", "iyr": "2013", "eyr": "2024", "ecl": "brn", "pid": "760753108", "byr": "1931", "hgt": "179cm"}},
		{PassportFields{"hcl": "#cfa07d", "eyr": "2025", "pid": "166559648", "iyr": "2011", "ecl": "brn", "hgt": "59in"}},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCountValid(t *testing.T) {
	got := CountValid([]Passport{
		{PassportFields{"ecl": "gry", "pid": "860033327", "eyr": "2020", "hcl": "#fffffd", "byr": "1937", "iyr": "2017", "cid": "147", "hgt": "183cm"}},
		{PassportFields{"iyr": "2013", "ecl": "amb", "cid": "350", "eyr": "2023", "pid": "028048884", "hcl": "#cfa07d", "byr": "1929"}},
		{PassportFields{"hcl": "#ae17e1", "iyr": "2013", "eyr": "2024", "ecl": "brn", "pid": "760753108", "byr": "1931", "hgt": "179cm"}},
		{PassportFields{"hcl": "#cfa07d", "eyr": "2025", "pid": "166559648", "iyr": "2011", "ecl": "brn", "hgt": "59in"}},
	}, Passport.IsValid)
	want := 2

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestIsValid(t *testing.T) {
	passport := Passport{
		PassportFields{"ecl": "gry", "pid": "860033327", "eyr": "2020", "hcl": "#fffffd", "byr": "1937", "iyr": "2017", "cid": "147", "hgt": "183cm"},
	}

	if !passport.IsValid() {
		t.Errorf("expected passport to be valid")
	}
}

func TestIsValidWithoutCid(t *testing.T) {
	passport := Passport{
		PassportFields{"hcl": "#ae17e1", "iyr": "2013", "eyr": "2024", "ecl": "brn", "pid": "760753108", "byr": "1931", "hgt": "179cm"},
	}

	if !passport.IsValid() {
		t.Errorf("expected passport to be valid")
	}
}

func TestIsInvalidNoHgt(t *testing.T) {
	passport := Passport{
		PassportFields{"iyr": "2013", "ecl": "amb", "cid": "350", "eyr": "2023", "pid": "028048884", "hcl": "#cfa07d", "byr": "1929"},
	}

	if passport.IsValid() {
		t.Errorf("expected passport to be invalid")
	}
}

func TestIsInvalidNoByrCid(t *testing.T) {
	passport := Passport{
		PassportFields{"hcl": "#cfa07d", "eyr": "2025", "pid": "166559648", "iyr": "2011", "ecl": "brn", "hgt": "59in"},
	}

	if passport.IsValid() {
		t.Errorf("expected passport to be invalid")
	}
}
