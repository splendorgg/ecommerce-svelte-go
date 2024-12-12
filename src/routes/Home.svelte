<script>
  import { onMount } from "svelte";
  import Navbar from "../Components/Navbar.svelte";
  import SideBar from "../Components/SideBar.svelte";
  import Products from "../Components/Products.svelte";
  import { publicService } from "../../instance";
  export let productmap = [];
  onMount(() => {
    publicService.get("api/products").then((response) => {
      productmap = response.data;
    });
  });
  let searchTerm = "";
  let filteredProducts = [];

  let selectedBrands = [];

  $: if (selectedBrands) getProductsbyBrands();
  const getProductsbyBrands = () => {
    searchTerm = "";
    if (selectedBrands.length === 0) {
      return (filteredProducts = []);
    }
    return (filteredProducts = productmap.filter((product) =>
      // product.Description.toLowerCase().includes(selectedBrands),
      selectedBrands.some((brand) =>
        product.Description.toLowerCase().includes(brand),
      ),
    ));
  };

  //SEARCH BAR
  const searchProducts = () => {
    return (filteredProducts = productmap.filter((product) =>
      product.Description.toLowerCase().includes(searchTerm.toLowerCase()),
    ));
  };
</script>

<head>
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
</head>
<header>
  <Navbar bind:searchTerm on:input={searchProducts} />
</header>

<main>
  <div class="container">
    <div class="sidebardiv">
      <SideBar bind:selectedBrands></SideBar>
    </div>
    <div class="products">
      {#if searchTerm && filteredProducts.length === 0}
        <p>No products found</p>
      {:else if filteredProducts.length > 0}
        {#each filteredProducts as product}
          <Products
            description={product.Description}
            image={product.Image}
            price={product.Price}
            rating={product.Rating}
          />
        {/each}
      {:else}
        {#each productmap as product}
          <Products
            description={product.Description}
            image={product.Image}
            price={product.Price}
            rating={product.Rating}
            onsale={product.IsOnSale}
          />
        {/each}
      {/if}
    </div>
  </div>
</main>
