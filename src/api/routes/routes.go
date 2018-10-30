package routes

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)

func GetPhrases(context *gin.Context) {

	phrases := []string{
		"La amistad es un alma que habita en dos cuerpos; un corazón que habita en dos almas.",
		"El ignorante afirma, el sabio duda y reflexiona.",
		"El sabio no dice todo lo que piensa, pero siempre piensa todo lo que dice.",
		"Considero más valiente al que conquista sus deseos que al que conquista a sus enemigos, ya que la victoria más dura es la victoria sobre uno mismo.",
		"Cualquiera puede enfadarse, eso es algo muy sencillo. Pero enfadarse con la persona adecuada, en el grado exacto, en el momento oportuno, con el propósito justo y del modo correcto. Eso, ciertamente, no resulta tan sencillo.",
		"La inteligencia consiste no sólo en el conocimiento, sino también en la destreza de aplicar los conocimientos en la práctica.",
		"Algunos creen que para ser amigos basta con querer, como si para estar sano bastara con desear la salud.",
		"No basta decir solamente la verdad, más conviene mostrar la causa de la falsedad.",
		"La esperanza es el sueño del hombre despierto.",
		"Sólo hay felicidad donde hay virtud y esfuerzo serio, pues la vida no es un juego.",
		"The roots of education are bitter, but the fruit is sweet.",
		"The best friend is the man who in wishing me well wishes it for my sake.",
		"You will never do anything in the world without courage.",
		"Hope is a waking dream.",
		"Quality is not an act, it is a habit.",

		"Avaro es quien no gasta en lo que debe ni cuando debe",
		"Demasiado poco valor es cobardía y demasiado valor es temeridad",
		"El fundamento del orden democrático es la libertad",
		"Es de importancia para quien desee alcanzar una certeza en su investigación el saber dudar a tiempoGarystóteles",
		"La dignidad no consiste en poseer honores, sino en merecerlos",
		"La duda es el principio de la sabiduría",
		"La educación es un ornamento en la prosperidad y un refugio en la adversidad",
		"La esperanza es el sueño del hombre despierto",
		"La felicidad consiste en unir sabiamente la virtud, la contemplación y los bienes exteriores",
		"La felicidad no consiste en tener mucho, sino en necesitar poco",
		"La felicidad no necesita de ningún bien exterior, sino que se basta a sí misma",
		"La finalidad del arte es dar cuerpo a la esencia secreta de las cosas, no el copiar su apariencia",
		"La verdadera felicidad consiste en hacer el bien",
		"Más se estima lo que con más trabajo se gana",
		"No se puede ser y no ser algo al mismo tiempo y bajo el mismo aspecto.",
		"Nuestro carácter es el resultado de nuestra conducta",
		"Si tanto me alaban será por alabarse a sí mismos, pues al alabarme dan a entender que me comprenden",
		"Todos los hombres desean naturalmente saber",
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(phrases)-1)

	context.JSON(http.StatusOK,gin.H{
		"phrase":phrases[n],
	})
}

func Ping(context *gin.Context) {

	context.JSON(http.StatusOK,gin.H{
		"status":http.StatusOK,
		"ping":"pong",
		"err":"",
	})

}

func handleErrors(context *gin.Context, err error) {
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

