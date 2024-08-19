<script>
    import { onMount } from 'svelte';
    import { goto } from '$app/navigation';
</script>

<aside>
    <h2>Menu</h2>
    <nav>
        <ul>
            <li><a href="#" on:click={() => goto('/blog')}>Blog</a></li>
            <li><a href="#" on:click={() => goto('/write')}>Yazı Yaz</a></li>
            <li><a href="#" on:click={() => goto('/yazılarım')}>Yazılarım</a></li>
            <li><a href="#" on:click={() => goto('/kaydedilenler')}>Kaydedilenler</a></li>
            <li><a href="#" on:click={() => goto('/login')}>Çıkış</a></li>
        </ul>
    </nav>
</aside>

<main class="content">

</main>

<style>
    h2{
        color:#fff;
    }
    aside {
        width: 260px;
        padding: 1rem;
        background: #191d69; 
        border-radius: 10px;
        box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
        position: fixed;
        top: 0;
        left: 0;
        height: 100%;
        overflow-y: auto;
        z-index: 1000; 
    }
    nav ul {
        list-style-type: none;
        padding: 0;
    }
    nav li {
        margin: 0.5rem 0;
    }
    nav a {
        display: block;
        padding: 0.75rem 1.25rem;
        text-decoration: none;
        color: #ffffff;
        border-radius: 6px;
        transition: background-color 0.3s, transform 0.2s;
        position: relative;
    }
    nav a::before {
        content: '';
        position: absolute;
        left: 0;
        width: 4px;
        height: 100%;
        background: #FF8000; 
        transition: background-color 0.3s;
        border-radius: 4px;
    }
    nav a:hover::before {
        background: #ffffff; 
    }
    nav a:hover {
        background: #4a4a7c; 
        transform: translateX(5px);
    }
    nav a.active {
        background: #4a4a7c; 
        font-weight: bold;
    }

    main.content {
        margin-left: 260px; 
        padding: 1rem;
        background: #f0f0f0; 
        min-height: 100vh; 
    }
</style>
