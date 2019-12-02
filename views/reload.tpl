<!DOCTYPE html>
<html lang="en">
<head>
  <title>刷新缓存 - {{.siteName}}</title>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="/static/bootstrap/4.3.1/css/bootstrap.min.css">
  <link rel="stylesheet" href="/static/markdown/github-markdown-css/3.0.1/github-markdown.min.css">
  <script src="/static/jquery/3.2.1/jquery.min.js"></script>
  <script src="/static/popper.js/1.15.0/umd/popper.min.js"></script>
  <script src="/static/bootstrap/4.3.1/js/bootstrap.min.js"></script>
  <script type="text/javascript" language="javascript">
var i = 5;
intervalid = setInterval("fun()", 300);
function fun(){
	if(i==0){
		history.go(-1);
		clearInterval(intervalid);
	}
	i--;
}
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
缓存已刷新
</article>

<!-- Footer 页脚 -->
{{template "footer"}}

</body>
</html>
