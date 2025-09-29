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

## Referenzen
* Go – Das Praxisbuch: Einstieg in Go und das Go-Ökosystem (1. Auflage 2020) [Book](https://subscription.packtpub.com/book/programming/9781803243054/) [Github](https://github.com/gobuch/code)
* Effective Go [Link](https://go.dev/doc/effective_go)