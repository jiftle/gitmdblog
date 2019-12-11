<!DOCTYPE html>
<html lang="en">
<head>
<title>{{.topic.Title}} - {{.topic.SiteName}}</title>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="/static/bootstrap/4.3.1/css/bootstrap.min.css">
<link rel="stylesheet" href="/static/markdown/github-markdown-css/3.0.1/github-markdown.min.css">
<script src="/static/jquery/3.2.1/jquery.min.js"></script>
<script src="/static/popper.js/1.15.0/umd/popper.min.js"></script>
<script src="/static/bootstrap/4.3.1/js/bootstrap.min.js"></script>
<!-- 右下角live2d效果 -->
<script type="text/javascript" charset="utf-8" async="" src="https://cdn.jsdelivr.net/npm/live2d-widget@3.1.4/lib/L2Dwidget.min.js"></script>
<!-- 右下角live2d效果 -->
<script type="text/javascript">
setTimeout(() => {
        L2Dwidget.init({
                "model": {
                "scale": 1
                },
                "display": {
                "position": "right",
                "width": 180,
                "height": 260,
                "hOffset": 0,
                "vOffset": -20
                },
                "mobile": {
                "show": false
                },
                "react": {
                "opacityDefault": 0.7,
                "opacityOnHover": 0.2
                }
                });
        }, 100)
</script>
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
<!-- 导航栏 -->
{{template "header"}}

<article class="markdown-body">
{{.topic.Content}}
</article>

<!-- Footer 页脚 -->
{{template "footer"}}

</body>
</html>
