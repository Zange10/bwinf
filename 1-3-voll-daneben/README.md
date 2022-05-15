# Kontext
In Las Vegas veranstaltet Casinobesitzer Al Capone Junior ein neues Gewinnspiel.Eine Runde
des Spiels läuft so ab: Jeder Teilnehmer zahlt einen Einsatz von 25 Dollar. Den Betrag setzt er
auf eine „Glückszahl“, eine ganze Zahl im Bereich von 1 bis 1000. Nachdem alle Teilnehmer 
gesetzt haben, wählt Al zehn Zahlen, ebenfalls im Bereich von 1 bis 1000. 
Anschließend werden die Gewinne ausgezahlt: Für jeden Teilnehmer wird diejenige von Als 
Zahlen bestimmt, die am nächsten an der Glückszahl des Teilnehmers liegt. Der Abstand 
dieser Zahl zu seiner Glückszahl ist der Gewinn des Teilnehmers.
Ein Beispiel: Bei einer Runde des Gewinnspiels waren unter anderem fünf alte Bekannte von 
Al dabei. Sie setzten auf diese Glückszahlen:
Bugsy: 1, Bonnie: 15, Clyde: 100, Mickey: 200, Lucky: 300.
Al wählte danach diese Zahlen:
1, 35, 117, 321, 448, 500, 678, 780, 802, 999 
Damit erhielt Bugsy keinen Gewinn, Bonnie erhielt 14, Clyde 17, Mickey 83 und Lucky 
21 Dollar.
Da Al seinen Reichtum mehren möchte, will er seine Zahlen so wählen, dass er möglichst 
wenig Gewinn auszahlen muss

# Aufgabe
Schreibe ein Programm, das die Glückszahlen aller Teilnehmer einer Runde einliest. Es soll 
dann für Al zehn Zahlen zwischen 1 und 1000 auswählen, die zu einer geringen Auszahlung 
führen. 
Wende dein Programm auf die Beispiele an, die du auf den 
BWINF-Webseiten findest. Für diese Beispiele soll Al weniger auszahlen müssen, als die Teilnehmer eingesetzt haben.

# Problem
- Interpretation von mir
Finden die Nummer mit dem kleinsten Abstand zu allen Zahlen im Array


- Idee
Daten in 10 gleich große Arrays aufteilen
Median von Teilen bestimmen,
Volia ! 10 Lucky Numbers