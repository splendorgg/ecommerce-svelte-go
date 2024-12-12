
<script>
  import Router from "svelte-spa-router";
  import Home from "./routes/Home.svelte";
  import Login from "./routes/Login.svelte";
  import Signup from "./routes/Signup.svelte";
  import AdminPanel from "./routes/AdminPanel.svelte";
</script>

<Router
  routes={{
    "/": Home,
    "/login": Login,
    "/signup": Signup,
    "/admin": AdminPanel,
  }}
/>
