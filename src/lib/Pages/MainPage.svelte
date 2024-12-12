<script>
    import { onMount } from "svelte";
    import Navbar from "./Components/Navbar.svelte";
    import Products from "./Components/Products.svelte";
    import Sidebar from "./Components/SideBar.svelte";

    import axios from "axios";
    $: products = [];
    onMount(() => {
        axios.get("http://localhost:8080/api/products").then((response) => {
            products = response.data;
        });
    });
</script>

<header>
    <Navbar></Navbar>
</header>

<main>
    <div class="container">
        <div class="sidebardiv">
            <Sidebar></Sidebar>
        </div>
        <div class="products">
            {#each products as product}
                <Products>
                    image={product.Image}
                    description={product.Description}
                    rating={product.Rating}
                    price={product.Price}
                </Products>
            {/each}
        </div>
    </div>
</main>
