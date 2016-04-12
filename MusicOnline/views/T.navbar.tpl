{{define "navbar"}}
<a class="navbar-brand" href="/"> MusicOnline </a>
	<div>
		<ul class="nav navbar-nav">
			<li {{if .IsHome}}class="active"{{end}}><a href="/">首页</a></li>
			<li {{if .IsNews}}class="active"{{end}}><a href="/news">资讯</a></li>
			<li {{if .IsStage}}class="active"{{end}}><a href="/stage">音乐Stage</a></li>
			<li {{if .IsMessage}}class="active"{{end}}><a href="/message">留言板</a></li>
		</ul>
	</div>

	<div class="pull-right">
		<ul class="nav navbar-nav">
			{{if .IsLogin}}
				<li><a href="/login?exit=true">logout</a></li>
			{{else}}
				<li><a href="/login">login</a></li>
			{{end}}
		</ul>
	</div>
{{end}}
