<script lang="ts">
    import { onMount } from "svelte";
    import { getSavedPosts } from "./api";
  
    let saves: Array<{ postID: number; postTitle: string; postDescription: string; datee: string }> = [];
  
    onMount(async () => {
      const token = localStorage.getItem('token');
      if (token) {
        try {
          saves = await getSavedPosts(token);
          console.log('Kaydedilen Yazılar:', saves);
        } catch (error) {
          console.error('Kaydedilen yazılar alınamadı:', error);
        }
      } else {
        console.error('Token bulunamadı');
      }
    });
  </script>
  
  <h1>Kaydedilen Yazılar</h1>
  
  {#if saves.length > 0}
    <ul>
      {#each saves as save}
        <li>
          <h3>{save.postTitle || 'Başlık Yok'}</h3>
          <p>{save.postDescription || 'Açıklama Yok'}</p>
          <small>{new Date(save.datee).toLocaleDateString()}</small>
        </li>
      {/each}
    </ul>
  {:else}
    <p>Henüz kaydedilmiş yazınız yok.</p>
  {/if}
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
  
    li:hover {
      background: #f1f1f1;
      box-shadow: 0 6px 12px rgba(0, 0, 0, 0.2);
    }
  
    h3 {
      margin: 0;
      color: #333;
    }
  
    p {
      margin: 0.5rem 0;
      color: #666;
    }
  
    small {
      display: block;
      margin-top: 0.5rem;
      color: #aaa;
    }
  
    p {
      font-size: 1rem;
    }
  </style>
  