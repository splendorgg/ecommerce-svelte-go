<script>
  import { Card, Button, Label, Input, Checkbox } from "flowbite-svelte";
  let email;
  let password;
  let errorMessage;
  import { publicService } from "../../instance";

  function login() {
    publicService
      .post(
        "login",
        {
          Email: email,
          Password: password,
        },
        {
          headers: {
            "Content-Type": "application/json",
          },
        },
      )
      .then((response) => {
        console.log(response);
        errorMessage = "";
        localStorage.setItem("token", response.data.token);
        if (email == "admin@hotmail.com") {
          window.location.href = "/#/admin";
        } else {
          window.location.href = "/";
        }
      })
      .catch((error) => {
        errorMessage = error.response
          ? error.response.data
          : "Login failed. Please try again.";
        console.error("Login error:", errorMessage);
      });
  }
</script>

<div class="logincarddiv">
  <Card class="logincard">
    <form
      class="flex flex-col space-y-6"
      action="/"
      on:submit|preventDefault={login}
    >
      <h3 class="text-xl font-medium text-gray-900 dark:text-white">
        Sign in to our platform
      </h3>
      <Label class="space-y-2">
        <span>Email</span>
        <Input
          type="email"
          name="email"
          placeholder="name@company.com"
          bind:value={email}
          required
        />
      </Label>
      <Label class="space-y-2">
        <span>Your password</span>
        <Input
          type="password"
          name="password"
          placeholder="•••••"
          bind:value={password}
          required
        />
      </Label>
      <div class="flex items-start">
        <Checkbox>Remember me</Checkbox>
        <a
          href="/"
          class="ms-auto text-sm text-primary-700 hover:underline dark:text-primary-500"
        >
          Lost password?
        </a>
      </div>
      <Button type="submit" class="w-full">Login to your account</Button>
      {#if errorMessage}
        <p style="color: red;">{errorMessage}</p>
      {/if}
      <div class="text-sm font-medium text-gray-500 dark:text-gray-300">
        Not registered? <a
          href="/#/signup"
          class="text-primary-700 hover:underline dark:text-primary-500"
        >
          Create account
        </a>
      </div>
    </form>
  </Card>
</div>
