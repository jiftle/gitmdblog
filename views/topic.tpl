<!DOCTYPE html>
<html lang="en">
<head>
  <title>{{.topic.Title}} - {{.topic.SiteName}}</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="/static/bootstrap/4.3.1/css/bootstrap.min.css">
  <script src="/static/jquery/3.2.1/jquery.min.js"></script>
  <script src="/static/popper.js/1.15.0/umd/popper.min.js"></script>
  <script src="/static/bootstrap/4.3.1/js/bootstrap.min.js"></script>
</head>
<style>
.markdown-body {
    box-sizing: border-box;
    min-width: 200px;
    max-width: 980px;
    margin: 0 auto;
    padding: 45px;
}
 
@media (max-width: 767px) {
    .markdown-body {
        padding: 15px;
    }
}
</style>
<body>
<div class="container">
  <div class="row">
    <div class="col-sm-12">
     <h3>{{.topic.Title}}</h3>
     <article class="markdown-body">
        {{.topic.Content}}
    </article>
    </div>
  </div>
</div>

</body>
</html>