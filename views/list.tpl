<!DOCTYPE html>
<html>
<head>
  <title>列表 - {{.siteName}}</title>
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

<div class="container">
  <div class="row">
    <div class="col-sm-12">
    <span>
<h1 class="title">{{.title}}存档</h1>
<ul>
    {{range .topics}}
    <li>[{{.Time.Format "06-01-02"}}] <a href="{{$.domain}}/{{.TopicID}}.html">{{.Title}}</a></li>
    {{end}}
</ul>
    </div>
  </div>
</div>

<!-- Footer 页脚 -->
{{template "footer"}}

</body>
</html>
