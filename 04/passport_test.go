package main

import (
	// "fmt"
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

func TestCountActuallyValidForValidPassports(t *testing.T) {
	input := `pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`
	passports := ParseInput(input)
	got := CountValid(passports, Passport.IsActuallyValid)
	want := 4

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCountActuallyValidForInvalidPassports(t *testing.T) {
	input := `eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007`
	passports := ParseInput(input)
	got := CountValid(passports, Passport.IsActuallyValid)
	want := 0

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

func TestIsValidByr(t *testing.T) {
	passport := Passport{
		PassportFields{"byr": "2002"},
	}

	if !passport.isValidByr() {
		t.Errorf("expected byr to be valid: %v", passport.Fields["byr"])
	}
}

func TestIsInValidByr(t *testing.T) {
	passport := Passport{
		PassportFields{"byr": "2003"},
	}

	if passport.isValidByr() {
		t.Errorf("expected byr to be invalid: %v", passport.Fields["byr"])
	}
}

func TestIsValidHgt1(t *testing.T) {
	passport := Passport{
		PassportFields{"hgt": "60in"},
	}

	if !passport.isValidHgt() {
		t.Errorf("expected hgt to be valid: %v", passport.Fields["hgt"])
	}
}

func TestIsValidHgt2(t *testing.T) {
	passport := Passport{
		PassportFields{"hgt": "190cm"},
	}

	if !passport.isValidHgt() {
		t.Errorf("expected hgt to be valid: %v", passport.Fields["hgt"])
	}
}

func TestIsInValidHgt1(t *testing.T) {
	passport := Passport{
		PassportFields{"hgt": "190in"},
	}

	if passport.isValidHgt() {
		t.Errorf("expected hgt to be invalid: %v", passport.Fields["hgt"])
	}
}

func TestIsInValidHgt2(t *testing.T) {
	passport := Passport{
		PassportFields{"hgt": "190"},
	}

	if passport.isValidHgt() {
		t.Errorf("expected hgt to be invalid: %v", passport.Fields["hgt"])
	}
}

func TestIsValidHcl(t *testing.T) {
	passport := Passport{
		PassportFields{"hcl": "#123abc"},
	}

	if !passport.isValidHcl() {
		t.Errorf("expected hcl to be valid: %v", passport.Fields["hcl"])
	}
}

func TestIsInValidHcl1(t *testing.T) {
	passport := Passport{
		PassportFields{"hcl": "#123abz"},
	}

	if passport.isValidHcl() {
		t.Errorf("expected hcl to be invalid: %v", passport.Fields["hcl"])
	}
}

func TestIsInValidHcl2(t *testing.T) {
	passport := Passport{
		PassportFields{"hcl": "123abc"},
	}

	if passport.isValidHcl() {
		t.Errorf("expected hcl to be invalid: %v", passport.Fields["hcl"])
	}
}

func TestIsValidEcl(t *testing.T) {
	passport := Passport{
		PassportFields{"ecl": "brn"},
	}

	if !passport.isValidEcl() {
		t.Errorf("expected ecl to be valid: %v", passport.Fields["ecl"])
	}
}

func TestIsInValidEcl(t *testing.T) {
	passport := Passport{
		PassportFields{"ecl": "wat"},
	}

	if passport.isValidEcl() {
		t.Errorf("expected ecl to be invalid: %v", passport.Fields["ecl"])
	}
}

func TestIsValidPid(t *testing.T) {
	passport := Passport{
		PassportFields{"pid": "000000001"},
	}

	if !passport.isValidPid() {
		t.Errorf("expected pid to be valid: %v", passport.Fields["pid"])
	}
}

func TestIsInValidPid(t *testing.T) {
	passport := Passport{
		PassportFields{"pid": "0123456789"},
	}

	if passport.isValidPid() {
		t.Errorf("expected pid to be invalid: %v", passport.Fields["pid"])
	}
}
