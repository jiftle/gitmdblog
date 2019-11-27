<!DOCTYPE html>
<html>
<head>
  <title>{{.siteName}}</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="/static/bootstrap/4.3.1/css/bootstrap.min.css">
  <script src="/static/jquery/3.2.1/jquery.min.js"></script>
  <script src="/static/popper.js/1.15.0/umd/popper.min.js"></script>
  <script src="/static/bootstrap/4.3.1/js/bootstrap.min.js"></script>
</head>
<body>


<div class="jumbotron text-center">
  <h2>你我相逢，皆是有缘分</h2>
</div>

<div class="container">
  <div class="row">
    <div class="col-sm-6">
     <h3>左侧</h3>
        {{range .topics_l}}
        <ul class="list-group">{{range .Topics}}
        <li class="list-group-item list-group-item-action">[{{.Time.Format "06-01-02"}}] <a href="{{$.domain}}/{{.TopicID}}.html">{{.Title}}</a></li>{{end}}
        </ul>
        {{end}}
    </div>
    <div class="col-sm-6">
      <h3>右侧</h3>
        {{range .topics_r}}
        <ul  class="list-group">{{range .Topics}}
            <li class="list-group-item list-group-item-action">[{{.Time.Format "06-01-02"}}] <a href="{{$.domain}}/{{.TopicID}}.html">{{.Title}}</a></li>{{end}}
        </ul>
        {{end}}
    </div>
  </div>
</div>

</body>
</html>