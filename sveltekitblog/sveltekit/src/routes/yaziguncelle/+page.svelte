<script lang="ts">
    import { onMount } from 'svelte';
    import { updatePost, getPostById } from './api';
    import { page } from '$app/stores';
    import { goto } from '$app/navigation';

    let postID: number;
    let postTitle: string = '';
    let postDescription: string = '';
    let tags: number[] = [];

    $: currentParams = $page.params;

    onMount(async () => {
        const queryString: string = window.location.search;
        const urlParams: URLSearchParams = new URLSearchParams(queryString);
        const paramValue: string | null = urlParams.get('id');
        postID = Number(paramValue);  

        const token = localStorage.getItem('token');
        if (token && postID) {
            try {
                const post = await getPostById(postID, token);
                postTitle = post.postTitle;
                postDescription = post.postDescription;
               
            } catch (error) {
                console.error('Yazı getirilemedi:', error);
            }
        }
    });

    async function handleSubmit() {
        const token = localStorage.getItem('token');
        if (token) {
            try {
                await updatePost(postID, token, { postTitle, postDescription, tags });
                goto('/yazılarım');  
            } catch (error) {
                console.error('Yazı güncellenemedi:', error);
            }
        }
    }
</script>

<form on:submit|preventDefault={handleSubmit}>
    <div>
        <label for="postTitle">Başlık</label>
        <input id="postTitle" type="text" bind:value={postTitle} required />
    </div>
    <div>
        <label for="postDescription">Açıklama</label>
        <textarea id="postDescription" bind:value={postDescription} required></textarea>
    </div>
 
    <button type="submit">Güncelle</button>
</form>
<style>
    form {
        max-width: 600px;
        margin: 2rem auto;
        padding: 1.5rem;
        border: 1px solid #e0e0e0;
        border-radius: 8px;
        background: #ffffff;
        box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    }

    div {
        margin-bottom: 1rem;
    }

    label {
        display: block;
        font-size: 1rem;
        font-weight: bold;
        margin-bottom: 0.5rem;
        color: #333;
    }

    input, textarea {
        width: 100%;
        padding: 0.75rem;
        border: 1px solid #dcdcdc;
        border-radius: 5px;
        font-size: 1rem;
        color: #333;
        box-sizing: border-box;
    }

    textarea {
        resize: vertical;
        min-height: 150px;
    }

    button {
        display: inline-block;
        padding: 0.75rem 1.25rem;
        background-color: #FF8000;
        color: #fff;
        border: none;
        border-radius: 5px;
        font-size: 1rem;
        cursor: pointer;
        transition: background-color 0.3s, transform 0.2s;
    }

    button:hover {
        background-color: #e67300;
        transform: scale(1.05);
    }
</style>

