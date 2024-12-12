<script>
    import { Input, Label, Button, Checkbox, A, Card } from "flowbite-svelte";
    let description;
    let image;
    let rating;
    let price;
    import { adminService } from "../../instance";
    import { onMount } from "svelte";
    onMount(() => {
        const token = localStorage.getItem("token");
        if (!token) {
            window.location.href = "/#/login";
        }
    });
    function addProduct() {
        console.log(description, image, rating, price);
        adminService
            .post("addproduct", {
                Description: description,
                Image: image,
                Rating: parseInt(rating),
                Price: parseInt(price),
            })
            .then((response) => {
                console.log(response);
                window.location.href = "/";
            })
            .catch((error) => {
                console.log(error);
            });
    }
</script>

<div class="adminaddproductdiv">
    <Card class="adminaddproduct">
        <form class="adminaddform" on:submit|preventDefault={addProduct}>
            <div class="grid gap-6 mb-6 md:grid-cols-2 adminadddiv">
                <div>
                    <Label for="last_name" class="mb-2">Product Name</Label>
                    <Input
                        type="text"
                        id="last_name"
                        placeholder="Enter product name"
                        bind:value={description}
                        required
                    />
                </div>

                <div>
                    <Label for="last_name" class="mb-2">Picture Link</Label>
                    <Input
                        type="text"
                        id="last_name"
                        placeholder="Enter picture link"
                        bind:value={image}
                        required
                    />
                </div>
                <div>
                    <Label for="last_name" class="mb-2">Rating</Label>
                    <Input
                        type="number"
                        id="last_name"
                        placeholder="Enter rating"
                        bind:value={rating}
                        required
                    />
                </div>
                <div>
                    <Label for="last_name" class="mb-2">Price</Label>
                    <Input
                        type="number"
                        id="last_name"
                        placeholder="Enter price"
                        bind:value={price}
                        required
                    />
                </div>

                <Button type="submit">Submit</Button>
            </div>
        </form>
    </Card>
</div>
