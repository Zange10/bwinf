# Aufgabe 4 - Schrebergärten

### Aufgabenstellung

Der Vorstand der neu gegründeten Schrebergartensiedlung von Rechteckingen bietet jedem Einwohner an, sich auf einer Wiese vor dem Ort einen Schrebergarten anzulegen. Ein Schrebergarten ist dabei immer rechteckig und nach den Himmelsrichtungen ausgerichtet. Jeder Schrebergarten hat also eine Nord-Süd-Ausdehnung (Länge) und eine Ost-West- Ausdehnung (Breite).Jeder Interessent soll ein Stück Land der von ihm gewünschten Ausmaße seines Schrebergartens zugeteilt bekommen. Dem Eigentümer der Wiese soll dafür ein rechteckiges Grundstück abgekauft werden, in das alle Schrebergärten hineinpassen und dessen Fläche  möglichst klein ist. Dazu müssen die Schrebergärten auf der Wiese geeignet nebeneinander angeordnet werden. Zufahrtswege werden hierbei nicht berücksichtigt. Die Bilder unten zeigen als Beispiel eine Reihe von gewünschten Schrebergartenflächen und eine gute Möglichkeit, sie flächensparend anzuordnen.

Hilf dem Vorstand und schreibe ein Programm, welches für eine gegebene Menge von Schrebergärten (jeweils bestimmt durch ihre Länge und Breite) eine möglichst günstige Anordnung auf der Wiese berechnet und grafisch ausgibt. Du kannst davon ausgehen, dass alle Längen und Breiten der Schrebergärten natürliche Zahlen sind.

### Lösungsansatz

1. Alle Möglichkeiten von Länge + Breite addieren und kleinste Zahl wählen, damit finden wir die obersten Zwei Gärten. Unten das gleiche machen oder ? einfach einpassen

Problem: 2D Packing, dazu gibt es viele verschiedene Algorithmen
https://en.wikipedia.org/wiki/Packing_problems

## Quick Reminder, 

You can run this Projekt with 
``` 
go run main.go drawing.go array.go calculation.go
```
or use ``` go build ``` to create an binary