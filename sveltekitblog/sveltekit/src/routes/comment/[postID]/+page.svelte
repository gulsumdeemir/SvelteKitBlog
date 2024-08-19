<script lang="ts">
    import { onMount } from 'svelte';
    import { page } from '$app/stores';
    import { getPostById, getCommentsByPostId, createComment } from '$lib/api';

    let postID: number;
    let postTitle: string = '';
    let postDescription: string = '';
    let comments: { id: number, commentt: string }[] = [];
    let newComment: string = '';

    $: currentParams = $page.params;

    onMount(async () => {
        const queryString: string = window.location.search;
        const urlParams: URLSearchParams = new URLSearchParams(queryString);
        const paramValue: string | null = urlParams.get('id');
        postID = Number(currentParams.postID);

        const token = localStorage.getItem('token');
        if (token && postID) {
            try {
                const post = await getPostById(postID, token);
                postTitle = post.postTitle;
                postDescription = post.postDescription;
                comments = await getCommentsByPostId(postID);
            } catch (error) {
                console.error('Veriler getirilemedi:', error);
            }
        }
    });

    async function handleCommentSubmit() {
        if (newComment.trim() !== '') {
            try {
                const token = localStorage.getItem('token');
                if (!token) {
                    throw new Error('Token bulunamadı.');
                }

                await createComment(postID, newComment, token);
                comments = await getCommentsByPostId(postID);
                newComment = '';
            } catch (error) {
                console.error('Yorum eklenemedi:', error);
            }
        }
    }
</script>

<main>
    <article class="post">
        <h1>{postTitle}</h1>
        <p>{postDescription}</p>
    </article>

    <section class="comments-section">
        <h2>Yorumlar</h2>
        {#if comments.length > 0}
            <ul class="comments-list">
                {#each comments as comment}
                    <li class="comment-item">
                        <p>{comment.commentt}</p>
                    </li>
                {/each}
            </ul>
        {:else}
            <p>Henüz yorum yapılmadı.</p>
        {/if}
    </section>

    <form on:submit|preventDefault={handleCommentSubmit} class="comment-form">
        <label for="newComment">Yeni Yorum</label>
        <textarea id="newComment" bind:value={newComment} placeholder="Yorumunuzu buraya yazın..." rows="4"></textarea>
        <button type="submit" class="submit-button">Gönder</button>
    </form>
</main>

<style>
    main {
        max-width: 800px;
        margin: 0 auto;
        padding: 1rem;
        background: #f9f9f9;
        border-radius: 8px;
        box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    }

    .post {
        margin-bottom: 2rem;
    }

    .post h1 {
        font-size: 2rem;
        color: #333;
        margin-bottom: 0.5rem;
    }

    .post p {
        font-size: 1.1rem;
        color: #555;
        line-height: 1.6;
    }

    .comments-section {
        margin-top: 2rem;
    }

    .comments-section h2 {
        font-size: 1.5rem;
        color: #333;
        margin-bottom: 1rem;
    }

    .comments-list {
        list-style: none;
        padding: 0;
    }

    .comment-item {
        background: #fff;
        padding: 0.75rem;
        border-radius: 4px;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
        margin-bottom: 0.5rem;
    }

    .comment-item p {
        margin: 0;
        font-size: 1rem;
        color: #333;
    }

    .comment-form {
        margin-top: 2rem;
        display: flex;
        flex-direction: column;
        gap: 1rem;
    }

    .comment-form label {
        font-size: 1.1rem;
        color: #333;
    }

    .comment-form textarea {
        width: 100%;
        padding: 0.75rem;
        border-radius: 4px;
        border: 1px solid #ccc;
        font-size: 1rem;
        resize: none;
    }

    .submit-button {
        padding: 0.75rem 1.5rem;
        border: none;
        border-radius: 4px;
        background: #FF8000;
        color: #fff;
        font-size: 1rem;
        cursor: pointer;
        transition: background 0.3s;
    }

    .submit-button:hover {
        background: #e67300;
    }
</style>
