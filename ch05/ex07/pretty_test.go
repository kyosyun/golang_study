package pretty

import (
	"bytes"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestForEachNode(t *testing.T) {
	for _, test := range []struct {
		query, expected string
	}{
		//{"", ""},
		{`
<!DOCTYPE html>
<html lang="ja">
<head>
<meta charset="UTF-8" />
<title>サイト名</title>
<link rel="stylesheet" type="text/css" media="screen" href="style.css" />
<!--[if lt IE 9]>
<script src="http://html5shiv.googlecode.com/svn/trunk/html5.js" type="text/javascript"></script>
<!--[endif]-->
</head>
</html>
`,
			`<!DOCTYPE html>
<html lang="ja">
  <head>
    <meta charset="UTF-8"/>
    <title>
      サイト名
    </title>
    <link rel="stylesheet" type="text/css" media="screen" href="style.css"/>
    <!-- [if lt IE 9]>
<script src="http://html5shiv.googlecode.com/svn/trunk/html5.js" type="text/javascript"></script>
<!--[endif] -->
  </head>
  <body>
  </body>
</html>
`},
	} {
		writer = new(bytes.Buffer)
		node, err := html.Parse(strings.NewReader(test.query))
		if err != nil {
			t.Errorf("parseErr \n %v", err)
		}
		forEachNode(node, startElement, endElement)
		result := writer.(*bytes.Buffer).String()

		if result != test.expected {
			t.Errorf("result =\n%v, but was =\n%v", result, test.expected)
		}
	}
}

// <body>
// <div id="wrap">
// <header id="head"><h1>サイト名</h1></header>
// <nav id="navi"><ul><li><a href="./">Home</a></li></ul></nav>
// <article class="post">
// <h2>記事タイトル</h2>
// <p>html5の構造化タグを使った練習ページ。</p>
// </article>
// <footer id="foot"><a href="./">サイト名</a></footer>
// </div>
// </body>

//  <body>
//    <div id="wrap">
//      <header id="head">
//        <h1>
//          サイト名
//        </h1>
//      </header>
//      <nav id="navi">
//        <ul>
//          <li>
//            <a href="./">
//              Home
//            </a>
//          </li>
//        </ul>
//      </nav>
//      <article class="post">
//        <h2>
//          記事タイトル
//        </h2>
//        <p>
//          html5の構造化タグを使った練習ページ。
//        </p>
//      </article>
//      <footer id="foot">
//        <a href="./">
//          サイト名
//        </a>
//      </footer>
//    </div>
//  </body>
