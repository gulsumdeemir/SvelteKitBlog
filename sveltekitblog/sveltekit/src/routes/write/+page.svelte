<script lang="ts">
    import { onMount } from 'svelte';
    import { createPost, getCategories, getTags } from './api'; 

    let title = '';
    let description = '';
    let selectedCategoryID: number = -1; 
    let selectedTagIDs: number[] = [];
    let categories: { categoryID: number; categoryName: string }[] = [];
    let tags: { tagID: number; tagName: string }[] = [];

  
    onMount(async () => {
        try {
            categories = await getCategories();
            tags = await getTags();
        } catch (error) {
            console.error('Kategoriler veya etiketler alınamadı:', error);
            alert('Kategoriler veya etiketler yüklenemedi. Lütfen sayfayı yenileyin.');
        }
    });

   
    async function handleSubmit() {
        if (selectedCategoryID === -1) { 
            alert('Lütfen bir kategori seçin.');
            return;
        }

        const newPost = {
            postTitle: title,
            postDescription: description,
            categoryID: selectedCategoryID,
            tags: selectedTagIDs,
        };

        try {
            const response = await createPost(newPost);
            if (response) {
                alert('Yazı başarıyla oluşturuldu!');
             
              
                title = '';
                description = '';
                selectedCategoryID = -1;
                selectedTagIDs = [];
                window.location.href = '/blog'; 
            }
        } catch (error) {
            console.error('Yazı oluşturulamadı:', error);
            alert('Yazı oluşturulamadı. Lütfen tekrar deneyin.');
        }
    }
</script>

<main>
    <h1>Yeni Yazı Oluştur</h1>

    <form on:submit|preventDefault={handleSubmit}>
        <div>
            <label for="title">Yazı Başlığı:</label>
            <input type="text" id="title" bind:value={title} required />
        </div>

        <div>
            <label for="description">Açıklama:</label>
            <textarea id="description" bind:value={description} required></textarea>
        </div>

        <div>
            <label for="category">Kategori Seç:</label>
            <select id="category" bind:value={selectedCategoryID} required>
                <option value={-1} disabled>Seçin</option> 
                {#each categories as { categoryID, categoryName }}
                    <option value={categoryID}>{categoryName}</option>
                {/each}
            </select>
        </div>

        <div>
            <label>Etiket Seç:</label>
            {#each tags as { tagID, tagName }}
                <label style="margin-right: 1rem;">
                    <input type="checkbox" bind:group={selectedTagIDs} value={tagID} />
                    {tagName}
                </label>
            {/each}
        </div>

        <button type="submit">Yeni Yazı Oluştur</button>
    </form>
</main>

<style>
    main {
        max-width: 800px;
        margin: 0 auto;
        padding: 2rem;
        background: #f4f4f9;
        border-radius: 8px;
        box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    }

    h1 {
        text-align: center;
        color: #191d69;
        margin-bottom: 1.5rem;
    }

    form {
        display: flex;
        flex-direction: column;
        gap: 1rem;
    }

    label {
        font-weight: bold;
        color: #191d69;
    }

    input[type="text"], textarea, select {
        width: 100%;
        padding: 0.75rem;
        border: 1px solid #ccc;
        border-radius: 4px;
        box-sizing: border-box;
    }

    textarea {
        height: 150px;
        resize: vertical;
    }

    select {
        background: #fff;
    }

    .checkbox-group {
        display: flex;
        flex-wrap: wrap;
        gap: 1rem;
    }

    .checkbox-group label {
        display: flex;
        align-items: center;
        gap: 0.5rem;
        cursor: pointer;
    }

    button {
        padding: 0.75rem;
        background-color: #FF8000;
        color: #fff;
        border: none;
        border-radius: 5px;
        cursor: pointer;
        font-size: 1rem;
        transition: background-color 0.3s, transform 0.2s;
        align-self: center;
    }

    button:hover {
        background-color: #e67300;
        transform: scale(1.05);
    }

    .error {
        color: red;
        margin-top: 1rem;
    }
</style>

