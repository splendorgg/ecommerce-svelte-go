<script>
    import { Card, Button, Label, Input, Checkbox } from "flowbite-svelte";
    import axios from "axios";

    let email;
    let password;
    let errorMessage;
    function signup() {
        axios
            .post(
                "http://localhost:8080/signup",
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
                localStorage.setItem("token", response.data.token);
                window.location.href = "/";
                errorMessage = "";
            })
            .catch((error) => {
                errorMessage = error.response
                    ? error.response.data
                    : "Signup failed. Please try again.";
                console.error("Signup error:", errorMessage);
            });
    }
</script>

<div class="signcarddiv">
    <Card class="signcard">
        <form
            class="flex flex-col space-y-6"
            action="/"
            on:submit|preventDefault={signup}
        >
            <h3 class="text-xl font-medium text-gray-900 dark:text-white">
                Sign Up to our platform
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

            <Button type="submit" class="w-full">Sign Up</Button>
            {#if errorMessage}
                <p style="color: red;">{errorMessage}</p>
            {/if}
        </form>
    </Card>
</div>
