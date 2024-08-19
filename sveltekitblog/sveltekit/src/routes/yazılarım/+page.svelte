<script lang="ts">
  import { onMount } from 'svelte';
  import { getUserPosts, deletePost, updatePost } from './api';
  import { goto } from '$app/navigation';
  
  let posts: Array<{ postID: number; postTitle: string; postDescription: string; datee: string }> = [];

  onMount(async () => {
      const token = localStorage.getItem('token');
      if (token) {
          try {
              posts = await getUserPosts(token);
          } catch (error) {
              console.error('Yazılar alınamadı:', error);
          }
      } else {
          console.error('Token bulunamadı');
      }
  });

  async function handleDelete(postID: number) {
      const token = localStorage.getItem('token');
      if (token) {
          try {
              await deletePost(postID, token);
              posts = posts.filter(post => post.postID !== postID);
          } catch (error) {
              console.error('Yazı silinemedi:', error);
          }
      }
  }
  async function handleUpdate(postID: number) {
    try {
        window.location.href = `/yaziguncelle?id=${postID}`;
    } catch (error) {
        console.error('Sayfaya yönlendirilemedi:', error);
    }
}

</script>
<h1>Yazılarım</h1>

<style>
   h1 {
        text-align: center;
        color: #191d69;
        margin-bottom: 2rem;
    }
    ul {
      list-style-type: none;
      padding: 0;
      margin: 0;
    }
  
    li {
      background: #f9f9f9;
      border: 1px solid #e0e0e0;
      border-radius: 8px;
      padding: 1rem;
      margin-bottom: 1rem;
      box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
      position: relative;
    }
  
    h3 {
      margin: 0 0 0.5rem;
      color: #333;
    }
  
    p {
      margin: 0 0 0.5rem;
      color: #666;
    }
  
    small {
      display: block;
      margin-bottom: 0.5rem;
      color: #999;
    }
  
    .button-container {
      display: flex;
      gap: 0.5rem;
      justify-content: flex-start;
      margin-top: 1rem;
    }
  
    button {
      padding: 0.5rem 1rem;
      border: none;
      border-radius: 5px;
      cursor: pointer;
      font-size: 0.875rem;
      transition: background-color 0.3s, transform 0.2s;
    }
  
    button:first-of-type {
      background-color: #FF8000;
      color: #fff;
    }
  
    button:first-of-type:hover {
      background-color: #e67300;
      transform: scale(1.05);
    }
  
    button:last-of-type {
      background-color: #dc3545;
      color: #fff;
    }
  
    button:last-of-type:hover {
      background-color: #c82333;
      transform: scale(1.05);
    }
  
    p.empty-message {
      text-align: center;
      color: #666;
      font-size: 1.25rem;
      margin-top: 2rem;
    }
  </style>
  
  {#if posts.length > 0}
    <ul>
      {#each posts as post}
        <li>
          <h3>{post.postTitle || 'Başlık Yok'}</h3>
          <p>{post.postDescription || 'Açıklama Yok'}</p>
          <small>{new Date(post.datee).toLocaleDateString()}</small>
          <div class="button-container">
            <button on:click={() => handleUpdate(post.postID)}>Güncelle</button>
            <button on:click={() => handleDelete(post.postID)}>Sil</button>
          </div>
        </li>
      {/each}
    </ul>
  {:else}
    <p class="empty-message">Henüz yazınız yok.</p>
  {/if}
  
  