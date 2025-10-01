# Go - Das Praxisbuch

Basierend auf dem Buch - _Go – Das Praxisbuch: Einstieg in Go und das Go-Ökosystem_.

## Übersicht

- [x] Kapitel 01: Einleitung
- [ ] Kapitel 02: Vorstellung der Syntax
- [ ] Kapitel 03: Command Line Interface (Projekt)
- [ ] Kapitel 04: Go Tooling
- [ ] Kapitel 05: Ein einfacher Webloader (Projekt)
- [ ] Kapitel 06: Eigene Pakete und Module
- [ ] Kapitel 07: Code generieren (Projekt)
- [ ] Kapitel 08: Concurrency-Grundlagen
- [ ] Kapitel 09: Concurrency Patterns
- [ ] Kapitel 10: Go Concurrency (Projekt)
- [ ] Kapitel 11: Testen und Benchmarks
- [ ] Kapitel 12: Image Resizer (Projekt)
- [ ] Kapitel 13: Interfaces
- [ ] Kapitel 14: Kopieren mit Reflection (Projekt)
- [ ] Kapitel 15: Fehlerbehandlung
- [ ] Kapitel 16: Ein einfacher Webserver (Projekt)

## Kapitel 01
Go ist eine junge Programmiersprache, deren Syntax überschaubar klein und gut dokumentiert ist. Das Ökosystem ist sehr umfangreich und beinhaltet u.a. Dokumentation, Test Framework, statische Codeanalyse, Race Detector, etc.

Go wurde 2007 bei Google entworfen und primär als Sprache für Cloud-Server entwickelt. U.a. sollten Multi-Core-Prozessoren besser unterstützt werden. Seit 2009 ist die Sprache und die Tools Open Source. 2012 erfolgte die Veröffentlichung der Version 1, womit Breaking Changes ausgeschlossen wurden.

Eine Go-Installation beinhaltet einen Compiler, die Standard-Bibliothek und Go Tooling:
* gopls: Language Server für Go
* goimports
* gofmt
* etc.

Das Kapitel beinhaltet Beispiele zur Verwendung von Funktionen aus dem "fmt"-Paket (fmt.Printf(), fmt.Sprintf(), fmt.Fprintf(), etc.).

## Kapitel 02
Go wurde als C-ähnliche Sprache entworfen und ist oft in den Lösungsmöglichkeiten beschränkter (erlaubt nur einen Lösungsweg), als andere Programmiersprachen.

Neben Variable und Konstanten, stellen insbesondere Pointer eine Besonderheit dar. Pointer (aka Zeiger) sind Variablen, die auf den Arbeitsspeicher zeigen und beinhalten selbst keinen Wert.

Lokale Variablen werden idR nach dem Funktionsaufruf innerhalb der Arbeitsspeichers wieder gelöscht. 

Eine `defer`Anweisung wird immer nach der Funktion ausegeführt, egal ob die Funktion durch ihr Ende, ein Return, oder eine Panik beendet wird.

Obwohl es in Go keine echten Klassen und Objekte gibt, ist objektorientierte Programmierung teilweise möglich. In Go gibt es programmieren gegen Schnittstellen, sowie die Objektkomposition (jedoch keine Vererbung). Dabei unterscheiden sich Methoden von Funktionen nur minimal, indem der Parameter des Types einfach vor dem Methodennamen gezogen wird.

```
func (r rectangle) area() int {
    return r.length * r.width
}
```

bzw:

```
func (r *rectangle) setLength(l int) {
    r.length * l
}
```

Ein Array hat immer eine fixe Größe. Im Gegensatz dazu gibt es ein Slice, das eine flexible Größe aufweist. Die Basis eines Slice ist jedoch ein Array, da ein Slice einen Pointer auf ein Array besitzt. Die Größe des Slices ist die `length` und erfasst die Anzahl der Elemente im Slice. Die Größe des Arrays wird über die `capacity` ausgelesen.

Neben der `if` Anweisung steht auch ein `switch` bereit, das sehr flexibel ist, wie in `Kapitel 2.16` gezeigt. Es muss kein `break` in eine `case` geben und innerhalb einer `case` Anweisung können wir auch mehrere Fälle auswerten. Die `default` Anweisung darf in beliebiger Position stehen. Sollte ein `break` jedoch definiert sein, wird die Anweisung an dieser Stelle beendet. Normalerweise wird die am Ende eines `case` Blocks die Anweisung beendet. Wollen wir das Verhalten überschreiben, können wir eine `fallthrough` Anweisung definieren, um weitere `case` Blöcke zu prüfen.

In Go gibt es nur die `for` Schleife, die ebenfalls sehr flexibel ist. Sie kann auch ohne Anweisung definiert werden. Mit `break` wird die Schleife abgebrochen. Mit `continue` wird sie zur nächsten Iterator springen. Oftmals wird eine Schleife auch in Verbindung mit `range` verwendet:

```
list := []string{"one", "two", "three"}
for i, v := range list {
    ...
}
```

Die `goto <label>` Funktion dient zum Springen zu einem Label. Dadurch wird der Code aber schwerer zu lesen. Ähnlich kann auch mit einem `continue <label>` gesprungen werden.

In Go wird jeder Text als UTF-8 codiert. Dadurch gibt es keine Probleme mit Umlauten oder asiatischen Zeichen. Ein String besteht aus mehreren Unicode-Zeichen. Pro Zeichen werden dafür bis zu 4 Bytes benötigt. D.h. Unicode-Zeichen sind bezüglich des Speicherbedarfs flexibel. Deshalb ist ein `string` ein Alias für `[]byte`, der jedoch in UTF-8 codiert ist. Ein Iterieren über einen String zeigt als Index das erste Byte des Zeichens an und der Value beinhaltet das Zeichen als `rune` (=`uint32`) Typ. 

## Referenzen
* Go – Das Praxisbuch: Einstieg in Go und das Go-Ökosystem (1. Auflage 2020) [Book](https://subscription.packtpub.com/book/programming/9781803243054/) [Github](https://github.com/gobuch/code)
* Effective Go [Link](https://go.dev/doc/effective_go)