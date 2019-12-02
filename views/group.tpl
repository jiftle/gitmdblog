<!DOCTYPE html>
<html>
<head>
  <title>分组 - {{.siteName}}</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="/static/bootstrap/4.3.1/css/bootstrap.min.css">
  <script src="/static/jquery/3.2.1/jquery.min.js"></script>
  <script src="/static/popper.js/1.15.0/umd/popper.min.js"></script>
  <script src="/static/bootstrap/4.3.1/js/bootstrap.min.js"></script>
</head>
<body>
<!-- 导航栏 -->
{{template "header"}}

<div class="jumbotron text-center">
  <h2>你我相逢，皆是有缘分</h2>
</div>
<div class="container">
    <div class="row">
    <div class="col-sm-6">
     <h3>标签</h3>
<ul>{{range .tags}}
    <li>[{{.TagName}}] <a href="/tag/{{.TagID}}">{{.TagName}}</a></li>
    {{end}}
</ul>
    </div>
    <div class="col-sm-6">
      <h3>月份</h3>
        <ul>{{range .months}}
    <li>[{{.Month}}] <a href="/{{.Month}}">{{.Month}}</a></li>
    {{end}}
</ul>
    </div>
  </div>

  </div>

</div>

<!-- Footer 页脚 -->
{{template "footer"}}

</body>
</html>
