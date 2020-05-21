<!DOCTYPE html>
<html lang="ja">
<head>
  <meta charset="utf-8">
  <title>Peachな掲示板 v2</title>
  <meta name="description" content="簡単なマルチスレッド掲示板">
  <link rel="stylesheet" href="/static/css/style.css" type="text/css" />
</head>

<body>
  <header>
    <nav>
      <a href="#!" class="text-link">Peachな掲示板</a>
      <div class="nav-right-part">
        <a href="login" class="button">Log in</a>
        <a href="signup" class="button">Sign up</a>
      </div>
    </nav>
  </header>
  <main>
    {{.LayoutContent}}
  </main>
  <footer>
    Github: <a href="https://github.com/PhysPeach/bbs">github.com/PhysPeach/bbs</a>
  </footer>
  <div class="backdrop"></div>
  <script src="/static/js/reload.min.js"></script>
</body>
</html>