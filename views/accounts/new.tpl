<section class="sign-up">
  <h1>Sign Up</h1>
  {{if .signupErrors}}
    <div class="errors">
      <ul>
        {{range $key, $signupError := .signupErrors}}
        <li>{{$signupError}}</li>
        {{end}}
      <ul>
    </div>
  {{end}}
  <div class="black-block">
    <form action="/signup" method="POST">
      {{.xsrf}}
      <input type="text" class="form-underbar-input" placeholder="Name" name="Name">
      <br>
      <input type="password" class="form-underbar-input" placeholder="Password" name="Password">
      <br>
      <input type="password" class="form-underbar-input" placeholder="Password Again" name="PasswordConfirmation">
      <br>
      <input type="submit" class="form-submit" value="Create!">
    </form>
  </div>
</section>