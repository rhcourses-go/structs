package dict2

import "strings"

/*
Dict ist ein Struct, das ein einziges Feld enthält, nämlich eine Liste von Einträgen.
*/
type Dict struct {
	Entries []Entry
}

/*
NewDict ist ein *Konstruktor* für ein Dict-Objekt.
Diese Funktion erwartet eine Reihe von Einträgen und liefert ein neues Dict,
das diese Einträge enthält.
*/
func NewDict(entries ...Entry) Dict {
	return Dict{Entries: entries}
}

/*
Append erwartet eine beliebige Anzahl an Entry-Objekten und fügt sie dum dict hinzu.

Append ist eine Methode: Die Entry-Objekte sind normale Parameter,
zusätzlich gibt es aber noch den Receiver.
Das ist ein zusätzlicher Parameter, der vor dem Funktionsnamen steht und angibt,
zu welchem Typ diese Methode gehört.
Der Receiver heißt hier dict und ist in diesem Fall ein Pointer,
weil die Methode den Inhalt von dict manipuliert.

Anmerkung: Diese Methode macht genau das, was wir im Usage-Beispiel in dict1 gemacht haben.
Dadurch wird die Komplexität hierher in die Definition von Dict verlagert und die
Verwendung vereinfacht.
*/
func (dict *Dict) Append(entries ...Entry) {
	dict.Entries = append(dict.Entries, entries...)
}

/*
LookupDe erwartet ein deutsches Wort als String und liefert eine Liste aller
englischen Übersetzungen dieses Wortes, die in dict bekannt sind.

Der Receiver dict muss hier kein Pointer sein, weil diese Methode das Dict nicht verändert.

Anmerkung: Auch diese Methode macht genau das, was wir im Usage-Beispiel in dict1 gemacht haben.
Es wurde i.W. die Schleife aus dem Beispiel hierher verschoben.
*/
func (dict Dict) LookupDe(de string) []string {
	result := []string{}
	for _, entry := range dict.Entries {
		if entry.De == de {
			result = append(result, entry.En)
		}
	}
	return result
}

/*
LookupEn erwartet ein englisches Wort als String und liefert eine Liste aller
deutschen Übersetzungen dieses Wortes, die in dict bekannt sind.
*/
func (dict Dict) LookupEn(en string) []string {
	result := []string{}
	for _, entry := range dict.Entries {
		if entry.En == en {
			result = append(result, entry.De)
		}
	}
	return result
}

/*
String liefert eine menschenlesbare String-Darstellung von Dict.
*/
func (dict Dict) String() string {
	if len(dict.Entries) == 0 {
		return "keine Einträge"
	}
	result_list := []string{}
	for _, entry := range dict.Entries {
		result_list = append(result_list, entry.String())
	}
	return strings.Join(result_list, "\n")
}
