package tags

import (
	"fmt"
	"testing"
)

func TestTagCalculation(t *testing.T) {
	txt := `Atlético-Trainer Simeone provoziert wieder Ärger

Madrid - Diego Simeone hat mal wieder Ärger. Der Trainer von Atlético Madrid soll seine Kontaktsperre mit dem Team umgangen haben. Der Dreikampf um die spanische Meisterschaft bleibt derweil packend


Auf die Tribüne musste der heißblütige Südamerikaner, weil im Ligaspiel zuvor gegen den FC Málaga ein Konter des Gegners mit einem Ballwurf von der Seitenlinie unterbunden worden war. Im Rückspiel der Champions League bei den Bayern darf Simeone jedoch wieder die Coaching-Zone bis auf den letzten Zentimeter ausreizen und sein Team mit seiner emotionalen Art anpeitschen. Allerdings mussten die Königlichen in San Sebastián ohne ihre Stars Cristiano Ronaldo, Karim Benzema und Toni Kroos lange um den Erfolg bangen, ehe der walisische 100-Millionen-Mann Gareth Bale zehn Minuten vor dem Abpfiff den Sieg sicherstellte.`
	result := Calculate(txt, "de")
	fmt.Println(result.Words)
}
