<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Peachな掲示板 v2</title>
    <link rel="stylesheet" href="/static/css/styles.css" />
    <script src="/static/js/reload.min.js"></script>
  </head>
  <body>
    <header>
      <nav class="nav">
        <div class="nav__container">
          <a class="nav__brand" href="/">Go Home</a>
        </div>
      </nav>
    </header>
    <main class="main">
      <h1 class="error">{{ .Message }}</h1>
    </main>
  </body>
</html>