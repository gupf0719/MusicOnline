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

		<form class="navbar-form navbar-left" role="search" method="post" action="/music">
       		<div class="form-group">
          		<input type="text" class="form-control" placeholder="输入歌名" name="musicname" >
        	</div>
        	<button type="submit" class="btn btn-default">Search</button>
    </form>

</div>
{{end}}
