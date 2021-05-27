package frenchspace

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestStream(t *testing.T) {
	r := strings.NewReader(`<!DOCTYPE html><html lang=fr><head>
	<meta charset="utf-8">
	<title>Yolo</title>
</head>
<body>
	<p>
		Lorem! ipsum? dolor: sit <i>amet?</i>.
	</p>
	<pre>// code: très important à ne pas formater!
	</pre>
</body>
</html>`)
	w := strings.Builder{}
	assert.NoError(t, Stream(&w, r))

	// expected := "<!DOCTYPE html><html lang=\"fr\"><head><meta charset=\"utf-8\"/><title>Yolo</title></head><body><p>Lorem\u202F! ipsum\u202F? dolor\u00A0: sit <i>amet\u202F?</i>.</p><pre>// code: très important à ne pas formater!\n\t</pre></body></html>"

	expected := `<!DOCTYPE html><html lang="fr"><head>
	<meta charset="utf-8"/>
	<title>Yolo</title>
</head>
<body>
	<p>
		` + "Lorem\u202F! ipsum\u202F? dolor\u00A0: sit <i>amet\u202F?</i>." + `
	</p>
	<pre>// code: très important à ne pas formater!
	</pre>

</body></html>`

	assert.Equal(t, expected, w.String())
}

func TestGetSpace(t *testing.T) {
	begin, body, end := getSpace("  Hello World! \n")
	assert.Equal(t, "  ", begin)
	assert.Equal(t, "Hello World!", body)
	assert.Equal(t, " \n", end)
}
