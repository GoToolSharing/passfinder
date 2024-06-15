package username

import "fmt"

func PatternFirst(firstName, lastName string) []string {
	return []string{firstName}
}

func PatternFirstLastNoSpace(firstName, lastName string) []string {
	return []string{firstName + lastName}
}

func PatternFLastNoDot(firstName, lastName string) []string {
	return []string{fmt.Sprintf("%c%s", firstName[0], lastName)}
}

func PatternSFirst(firstName, lastName string) []string {
	return []string{fmt.Sprintf("%c%s", lastName[0], firstName)}
}

func PatternLastFirstInit(firstName, lastName string) []string {
	return []string{fmt.Sprintf("%s%c", lastName, firstName[0])}
}

func PatternLast(firstName, lastName string) []string {
	return []string{lastName}
}

func PatternFirstInitLastInit(firstName, lastName string) []string {
	return []string{fmt.Sprintf("%c%c", firstName[0], lastName[0])}
}
