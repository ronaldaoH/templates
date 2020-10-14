// Code generated by qtc from "basepage.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// This is a base page template. All the other template pages implement this interface.
//

//line templates/basepage.qtpl:3
package templates

//line templates/basepage.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/basepage.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/basepage.qtpl:4
type Page interface {
//line templates/basepage.qtpl:4
	Title() string
//line templates/basepage.qtpl:4
	StreamTitle(qw422016 *qt422016.Writer)
//line templates/basepage.qtpl:4
	WriteTitle(qq422016 qtio422016.Writer)
//line templates/basepage.qtpl:4
	Body() string
//line templates/basepage.qtpl:4
	StreamBody(qw422016 *qt422016.Writer)
//line templates/basepage.qtpl:4
	WriteBody(qq422016 qtio422016.Writer)
//line templates/basepage.qtpl:4
	UserInfo() string
//line templates/basepage.qtpl:4
	StreamUserInfo(qw422016 *qt422016.Writer)
//line templates/basepage.qtpl:4
	WriteUserInfo(qq422016 qtio422016.Writer)
//line templates/basepage.qtpl:4
}

// Page prints a page implementing Page interface.

//line templates/basepage.qtpl:15
func StreamPageTemplate(qw422016 *qt422016.Writer, p Page) {
//line templates/basepage.qtpl:15
	qw422016.N().S(`
<html>
	<head>
		<title>`)
//line templates/basepage.qtpl:18
	p.StreamTitle(qw422016)
//line templates/basepage.qtpl:18
	qw422016.N().S(`</title>
		<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.1/css/bulma.min.css">
        <!-- You should properly set the path from the main file. -->
        <style>

        </style>
	</head>
	<body>

	    <nav class="navbar" role="navigation" aria-label="main navigation">
          <div class="navbar-brand">
            <a class="navbar-item" href="/">
            <img src="https://ronaldao.com/wp-content/uploads/2020/10/Compara-YA.png" alt="Compara los precios de tus productos favoritos en tiempo real." width="112" height="28">
            </a>

            <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false">
              <span aria-hidden="true"></span>
              <span aria-hidden="true"></span>
              <span aria-hidden="true"></span>
            </a>
          </div>
        </nav>
        <section class="hero is-success is-medium">
          <div class="hero-body">
            <div class="container">
              <h1 class="title">
                ¡Compara YA!
              </h1>


              <h2 class="subtitle">
                Compara el precio de tus productos favoritos en tiempo real. <br>
                Saga Fallabela, Ripley, Sodimac y Tottus.
              </h2>

            </div>
          </div>
        </section>
        <div class="container">
            `)
//line templates/basepage.qtpl:57
	p.StreamBody(qw422016)
//line templates/basepage.qtpl:57
	qw422016.N().S(`
            <section>
               `)
//line templates/basepage.qtpl:59
	p.StreamUserInfo(qw422016)
//line templates/basepage.qtpl:59
	qw422016.N().S(`
            </section>
        </div>
        <footer class="footer">
          <div class="content has-text-centered">
            <p>
              <strong>¡Compara YA!</strong> by <a href="https://github.com/ronaldaoH"> rorolas_</a>.
            </p>
          </div>
        </footer>
	</body>
</html>
`)
//line templates/basepage.qtpl:71
}

//line templates/basepage.qtpl:71
func WritePageTemplate(qq422016 qtio422016.Writer, p Page) {
//line templates/basepage.qtpl:71
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/basepage.qtpl:71
	StreamPageTemplate(qw422016, p)
//line templates/basepage.qtpl:71
	qt422016.ReleaseWriter(qw422016)
//line templates/basepage.qtpl:71
}

//line templates/basepage.qtpl:71
func PageTemplate(p Page) string {
//line templates/basepage.qtpl:71
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/basepage.qtpl:71
	WritePageTemplate(qb422016, p)
//line templates/basepage.qtpl:71
	qs422016 := string(qb422016.B)
//line templates/basepage.qtpl:71
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/basepage.qtpl:71
	return qs422016
//line templates/basepage.qtpl:71
}

//line templates/basepage.qtpl:73
type BasePage struct{}

//line templates/basepage.qtpl:75
func (p *BasePage) StreamTitle(qw422016 *qt422016.Writer) {
//line templates/basepage.qtpl:75
	qw422016.N().S(` titulo por si acaso `)
//line templates/basepage.qtpl:75
}

//line templates/basepage.qtpl:75
func (p *BasePage) WriteTitle(qq422016 qtio422016.Writer) {
//line templates/basepage.qtpl:75
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/basepage.qtpl:75
	p.StreamTitle(qw422016)
//line templates/basepage.qtpl:75
	qt422016.ReleaseWriter(qw422016)
//line templates/basepage.qtpl:75
}

//line templates/basepage.qtpl:75
func (p *BasePage) Title() string {
//line templates/basepage.qtpl:75
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/basepage.qtpl:75
	p.WriteTitle(qb422016)
//line templates/basepage.qtpl:75
	qs422016 := string(qb422016.B)
//line templates/basepage.qtpl:75
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/basepage.qtpl:75
	return qs422016
//line templates/basepage.qtpl:75
}

//line templates/basepage.qtpl:76
func (p *BasePage) StreamBody(qw422016 *qt422016.Writer) {
//line templates/basepage.qtpl:76
	qw422016.N().S(` cuerpo por si acaso`)
//line templates/basepage.qtpl:76
}

//line templates/basepage.qtpl:76
func (p *BasePage) WriteBody(qq422016 qtio422016.Writer) {
//line templates/basepage.qtpl:76
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/basepage.qtpl:76
	p.StreamBody(qw422016)
//line templates/basepage.qtpl:76
	qt422016.ReleaseWriter(qw422016)
//line templates/basepage.qtpl:76
}

//line templates/basepage.qtpl:76
func (p *BasePage) Body() string {
//line templates/basepage.qtpl:76
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/basepage.qtpl:76
	p.WriteBody(qb422016)
//line templates/basepage.qtpl:76
	qs422016 := string(qb422016.B)
//line templates/basepage.qtpl:76
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/basepage.qtpl:76
	return qs422016
//line templates/basepage.qtpl:76
}

//line templates/basepage.qtpl:77
func (p *BasePage) StreamUserInfo(qw422016 *qt422016.Writer) {
//line templates/basepage.qtpl:77
	qw422016.N().S(` No te estamos vigilando :P `)
//line templates/basepage.qtpl:77
}

//line templates/basepage.qtpl:77
func (p *BasePage) WriteUserInfo(qq422016 qtio422016.Writer) {
//line templates/basepage.qtpl:77
	qw422016 := qt422016.AcquireWriter(qq422016)
//line templates/basepage.qtpl:77
	p.StreamUserInfo(qw422016)
//line templates/basepage.qtpl:77
	qt422016.ReleaseWriter(qw422016)
//line templates/basepage.qtpl:77
}

//line templates/basepage.qtpl:77
func (p *BasePage) UserInfo() string {
//line templates/basepage.qtpl:77
	qb422016 := qt422016.AcquireByteBuffer()
//line templates/basepage.qtpl:77
	p.WriteUserInfo(qb422016)
//line templates/basepage.qtpl:77
	qs422016 := string(qb422016.B)
//line templates/basepage.qtpl:77
	qt422016.ReleaseByteBuffer(qb422016)
//line templates/basepage.qtpl:77
	return qs422016
//line templates/basepage.qtpl:77
}
