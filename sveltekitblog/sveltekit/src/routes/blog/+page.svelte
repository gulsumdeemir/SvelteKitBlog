<script lang="ts">
  import { onMount } from 'svelte';
  import { getCategories, getPostsByCategory, getAllPosts, savePost } from './api'; 
  import { goto } from '$app/navigation'; 

  let categories: { categoryID: number; categoryName: string }[] = [];
  let posts: { postID: number; postTitle: string; postDescription: string }[] = [];
  let selectedCategoryID: number | null = null;

  onMount(async () => {
    categories = await getCategories();
    posts = await getAllPosts(); 
  });

  async function fetchPostsByCategory(categoryID: number) {
    selectedCategoryID = categoryID;
    posts = await getPostsByCategory(categoryID);
  }

  function goToPost(postID: number) {
    goto(`/comment/${postID}`); 
  }

  async function save(postID: number) {
  const token = localStorage.getItem('token');
   
  if (token) {
    try {
      const response = await savePost(postID, token); 
      console.log('Yazı başarıyla kaydedildi:', response);
      goto('/user/saved'); 
    } catch (error) {
      console.error('Yazı kaydedilemedi:', error);
    }
  } else {
    console.error('Token bulunamadı', { token });
  }
}

</script>

<main>
  <h1>Kategoriler</h1>

  <div class="categories">
    {#each categories as { categoryID, categoryName }}
      <button on:click={() => fetchPostsByCategory(categoryID)}>
        {categoryName}
      </button>
    {/each}
  </div>

  <div class="posts">
    {#each posts as { postID, postTitle, postDescription }}
      <div class="post-card">
        <h2>{postTitle}</h2>
        <p>{postDescription}</p>
        <button on:click={() => goToPost(postID)}>Devamını Oku</button>
        <div class="save-icon" on:click={() => save(postID)}>
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" class="icon">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 3h14a2 2 0 012 2v16l-7-4-7 4V5a2 2 0 012-2z"/>
          </svg>
        </div>
      </div>
    {/each}
  </div>
</main>

<style>
  main {
    padding: 2rem;
    color: #333;
    font-family: Arial, sans-serif;
  }

  h1 {
    text-align: center;
    color: #191d69;
    margin-bottom: 2rem;
  }

  .categories {
    display: flex;
    gap: 1rem;
    justify-content: center;
    margin-bottom: 2rem;
  }

  .categories button {
    padding: 0.75rem 1.5rem;
    border: none;
    border-radius: 8px;
    background: #FF8000;
    color: #fff;
    font-size: 1rem;
    cursor: pointer;
    transition: background 0.3s, transform 0.2s;
  }

  .categories button:hover {
    background: #e67300;
    transform: scale(1.05);
  }

  .posts {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
    max-width: 800px;
    margin: 0 auto;
    position: relative;
  }

  .post-card {
    background: #ffffff;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    padding: 1rem;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    position: relative; 
  }

  .post-card h2 {
    color: #191d69;
    margin-bottom: 0.5rem;
    font-size: 1.25rem;
  }

  .post-card p {
    color: #555;
    margin-bottom: 1rem;
    font-size: 1rem;
  }

  .post-card button {
    padding: 0.25rem 0.75rem;
    font-size: 0.875rem;
    border: none;
    border-radius: 6px;
    background: #FF8000;
    color: #fff;
    cursor: pointer;
    transition: background 0.3s, transform 0.2s;
  }

  .post-card button:hover {
    background: #e67300;
    transform: scale(1.05);
  }

  .save-icon {
    position: absolute;
    bottom: 1rem;
    right: 1rem;
    cursor: pointer;
  }

  .save-icon .icon {
    width: 24px;
    height: 24px;
    stroke: #FF8000;
    stroke-width: 2;
    fill: none;
    transition: stroke 0.3s;
  }

  .save-icon:hover .icon {
    stroke: #e67300;
  }
</style>
